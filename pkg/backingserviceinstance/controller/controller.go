package controller

import (
	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"

	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
	osclient "github.com/openshift/origin/pkg/client"
	"io/ioutil"
	kapi "k8s.io/kubernetes/pkg/api"
	kerrors "k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/record"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/util"
	"net/http"
	"regexp"
	"strings"
)

// NamespaceController is responsible for participating in Kubernetes Namespace termination
// Use the NamespaceControllerFactory to create this controller.
type BackingServiceInstanceController struct {
	// Client is an OpenShift client.
	Client osclient.Interface
	// KubeClient is a Kubernetes client.
	KubeClient kclient.Interface
	recorder   record.EventRecorder
}

type fatalError string

func (e fatalError) Error() string {
	return "fatal error handling BackingServiceInstanceController: " + string(e)
}

// Handle processes a namespace and deletes content in origin if its terminating
func (c *BackingServiceInstanceController) Handle(bsi *backingserviceinstanceapi.BackingServiceInstance) (result error) {
	//glog.Infoln("bsi handler called.", bsi.Name)
	//c.recorder.Eventf(bsi, "Debug", "bsi handler called.%s", bsi.Name)

	changed := false
	bs := &backingserviceapi.BackingService{}
	if bsi.Annotations[backingserviceinstanceapi.UPS] != "true" {
	bsp, err := c.Client.BackingServices("openshift").Get(bsi.Spec.BackingServiceName)
	if err != nil {
		return err
	}else{
		bs = bsp
	}
	}

	switch bsi.Status.Phase {
	default:

		result = fmt.Errorf("unknown phase: %s", bsi.Status.Phase)

	case backingserviceinstanceapi.BackingServiceInstancePhaseDeleted:

		glog.Infoln("bsi delete etcd ", bsi.Name)

		result = c.Client.BackingServiceInstances(bsi.Namespace).Delete(bsi.Name)

	case "":

		bsi.Status.Phase = backingserviceinstanceapi.BackingServiceInstancePhaseProvisioning

		changed = true

		fallthrough

	case backingserviceinstanceapi.BackingServiceInstancePhaseProvisioning:
		if bsi.Status.Action == backingserviceinstanceapi.BackingServiceInstanceActionToDelete {
			return c.Client.BackingServiceInstances(bsi.Namespace).Delete(bsi.Name)
		}

		glog.Infoln("bsi provisioning ", bsi.Name)
		//c.recorder.Eventf(bsi, "Provisioning", "bsi:%s, service:%s", bsi.Name, bsi.Spec.BackingServiceName)

		plan_found := false
		for _, plan := range bs.Spec.Plans {
			if bsi.Spec.BackingServicePlanGuid == plan.Id {
				bsi.Spec.BackingServicePlanName = plan.Name
				plan_found = true
				break
			}
		}

		if !plan_found {
			c.recorder.Eventf(bsi, kapi.EventTypeNormal, "Provisioning", "plan (%s) in bs(%s) for bsi (%s) not found",
				bsi.Spec.BackingServicePlanGuid, bsi.Spec.BackingServiceName, bsi.Name)
			result = fmt.Errorf("plan (%s) in bs(%s) for bsi (%s) not found",
				bsi.Spec.BackingServicePlanGuid, bsi.Spec.BackingServiceName, bsi.Name)
			break
		}

		// ...

		glog.Infoln("bsi provisioning servicebroker_load, ", bsi.Name)
		//c.recorder.Eventf(bsi, "Provisioning", "bsi %s provisioning servicebroker_load", bsi.Name)
		bsInstanceID := string(util.NewUUID())

		servicebroker, err := servicebroker_load(c.Client, bs.GenerateName)
		if err != nil {
			result = err
			break
		}

		serviceinstance := &ServiceInstance{}
		serviceinstance.ServiceId = bs.Spec.Id
		serviceinstance.PlanId = bsi.Spec.BackingServicePlanGuid
		serviceinstance.OrganizationGuid = bsi.Namespace

		glog.Infoln("bsi provisioning servicebroker_create_instance, ", bsi.Name)

		svcinstance, err := servicebroker_create_instance(serviceinstance, bsInstanceID, servicebroker)
		if err != nil {
			result = err
			c.recorder.Eventf(bsi, kapi.EventTypeWarning, "Provisioning", err.Error())
			break
		} else {
			c.recorder.Eventf(bsi, kapi.EventTypeNormal, "Provisioning", "bsi provisioning done, instanceid: %s", bsInstanceID)
			glog.Infoln("bsi provisioning servicebroker_create_instance done, ", bsi.Name)
		}

		bsi.Spec.DashboardUrl = svcinstance.DashboardUrl
		bsi.Spec.InstanceID = bsInstanceID
		bsi.Spec.BackingServiceSpecID = bs.Spec.Id
		if bsi.Spec.Parameters == nil {
			bsi.Spec.Parameters = make(map[string]string)
		}
		bsi.Spec.Parameters["instance_id"] = bsInstanceID

		bsi.Status.Phase = backingserviceinstanceapi.BackingServiceInstancePhaseUnbound

		changed = true

		glog.Infoln("bsi inited. ", bsi.Name)

	case backingserviceinstanceapi.BackingServiceInstancePhaseUnbound:
		switch bsi.Status.Action {
		case backingserviceinstanceapi.BackingServiceInstanceActionToDelete:
			if bsi.Annotations[backingserviceinstanceapi.UPS] == "true" {
				glog.Infoln(backingserviceinstanceapi.UPS, " bsi deleted ", bsi.Name)

				bsi.Status.Phase = backingserviceinstanceapi.BackingServiceInstancePhaseDeleted

				bsi.Status.Action = ""
				changed = true
			} else {
				if result = c.deleteInstance(bs, bsi); result == nil {
					changed = true
				}
			}
			c.recorder.Eventf(bsi, kapi.EventTypeNormal, "Deleting", "instance:%s [%v]", bsi.Name, changed)
		case backingserviceinstanceapi.BackingServiceInstanceActionToBind:
			dcname := c.get_deploymentconfig_name(bsi, backingserviceinstanceapi.BindDeploymentConfigBinding)
			if bsi.Annotations[backingserviceinstanceapi.UPS] == "true" {
				if result = c.bindInstanceUPS(dcname, bsi); result == nil {
					changed = true
				}
			}else {
				if result = c.bindInstance(dcname, bs, bsi); result == nil {
					changed = true
				}
			}
			c.recorder.Eventf(bsi, kapi.EventTypeNormal, "Binding", "instance: %s, dc: %s [%v]", bsi.Name, dcname, changed)
		}
	case backingserviceinstanceapi.BackingServiceInstancePhaseBound:
		switch bsi.Status.Action {
		case backingserviceinstanceapi.BackingServiceInstanceActionToUnbind:

			dcname := c.get_deploymentconfig_name(bsi, backingserviceinstanceapi.BindDeploymentConfigUnbinding)
			if bsi.Annotations[backingserviceinstanceapi.UPS] == "true"{
				if result = c.unbindInstanceUPS(dcname, bsi); result == nil {
					changed = true
				}
			}else {
				if result = c.unbindInstance(dcname, bs, bsi); result == nil {
					changed = true
				}
			}
			c.recorder.Eventf(bsi, kapi.EventTypeNormal, "Unbinding", "instance: %s, dc: %s [%v]", bsi.Name, dcname, changed)
		case backingserviceinstanceapi.BackingServiceInstanceActionToBind:
			dcname := c.get_deploymentconfig_name(bsi, backingserviceinstanceapi.BindDeploymentConfigBinding)
			if bsi.Annotations[backingserviceinstanceapi.UPS] == "true" {
				if result = c.bindInstanceUPS(dcname, bsi); result == nil {
					changed = true
				}
			}else {
			if result = c.bindInstance(dcname, bs, bsi); result == nil {
				changed = true
			}
			}
			c.recorder.Eventf(bsi, kapi.EventTypeNormal, "Binding", "instance: %s, dc: %s [%v]", bsi.Name, dcname, changed)

		default:
			glog.Info("check dc healthy.")
			c.check_dc_healthy(bsi)

		}

	}

	if result != nil {

		err_msg := result.Error()
		/*
			if err_msg != bsi.Status.Error {
				glog.Info("#:", err_msg, "#:", bsi.Status.Error)
				changed = true
				bsi.Status.Error = err_msg
			}
		*/

		glog.Infoln("bsi controller error. ", err_msg)
		c.recorder.Eventf(bsi, kapi.EventTypeWarning, "Error", err_msg)
	}

	if changed {
		glog.Infoln("bsi etc changed and update. ")

		c.Client.BackingServiceInstances(bsi.Namespace).Update(bsi)
	}

	return
}

func has_action_word(text, word backingserviceinstanceapi.BackingServiceInstanceAction) bool {
	return strings.Index(string(text), string(word)) >= 0
}

func remove_action_word(text, word backingserviceinstanceapi.BackingServiceInstanceAction) backingserviceinstanceapi.BackingServiceInstanceAction {
	for {
		index := strings.Index(string(text), string(word))
		if index >= 0 {
			text = text[:index] + text[index+len(word):]
			continue
		}
		break
	}

	return text
}

func servicebroker_load(c osclient.Interface, name string) (*ServiceBroker, error) {
	servicebroker := &ServiceBroker{}
	if sb, err := c.ServiceBrokers().Get(name); err != nil {
		return nil, err
	} else {
		servicebroker.Url = sb.Spec.Url
		servicebroker.UserName = sb.Spec.UserName
		servicebroker.Password = sb.Spec.Password
		return servicebroker, nil
	}
}

func checkIfPlanidExist(client osclient.Interface, planId string) (bool, *backingserviceapi.BackingService, error) {

	items, err := client.BackingServices("openshift").List(kapi.ListOptions{})

	if err != nil {
		return false, nil, err
	}

	for _, bs := range items.Items {
		for _, plans := range bs.Spec.Plans {
			if planId == plans.Id {
				glog.Info("we found plan id at plan", bs.Spec.Name)

				return true, &bs, nil
			}
		}
	}
	return false, nil, fatalError(fmt.Sprintf("Can't find plan id %s", planId))

}

func commToServiceBroker(method, path string, jsonData []byte, header map[string]string) (resp *http.Response, err error) {

	fmt.Println(method, path, string(jsonData))

	req, err := http.NewRequest(strings.ToUpper(method) /*SERVICE_BROKER_API_SERVER+*/, path, bytes.NewBuffer(jsonData))

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}
	return http.DefaultClient.Do(req)
}

type ServiceBroker struct {
	Url      string
	UserName string
	Password string
}

type ServiceInstance struct {
	ServiceId        string `json:"service_id"`
	PlanId           string `json:"plan_id"`
	OrganizationGuid string `json:"organization_guid"`
	SpaceGuid        string `json:"space_guid"`
	//Incomplete       bool        `json:"accepts_incomplete, omitempty"`
	Parameters interface{} `json:"parameters, omitempty"`
}
type LastOperation struct {
	State                    string `json:"state"`
	Description              string `json:"description"`
	AsyncPollIntervalSeconds int    `json:"async_poll_interval_seconds, omitempty"`
}
type CreateServiceInstanceResponse struct {
	DashboardUrl  string         `json:"dashboard_url"`
	LastOperation *LastOperation `json:"last_operation, omitempty"`
}

type ServiceBinding struct {
	ServiceId       string                 `json:"service_id"`
	PlanId          string                 `json:"plan_id"`
	AppGuid         string                 `json:"app_guid,omitempty"`
	BindResource    map[string]string      `json:"bind_resource,omitempty"`
	Parameters      map[string]interface{} `json:"parameters,omitempty"`
	svc_instance_id string
}

type ServiceBindingResponse struct {
	Credentials     Credential `json:"credentials"`
	SyslogDrainUrl  string     `json:"syslog_drain_url"`
	RouteServiceUrl string     `json:"route_service_url"`
}

type Credential struct {
	Uri      string `json:"uri"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Vhost    string `json:"vhost"`
	//Database string `json:"database"`
}

func servicebroker_create_instance(param *ServiceInstance, instance_guid string, sb *ServiceBroker) (*CreateServiceInstanceResponse, error) {
	jsonData, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = basicAuthStr(sb.UserName, sb.Password)

	resp, err := commToServiceBroker("PUT", "http://"+sb.Url+"/v2/service_instances/"+instance_guid, jsonData, header)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	glog.Infof("respcode from /v2/service_instances/%s: %v", instance_guid, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	svcinstance := &CreateServiceInstanceResponse{}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated ||
		resp.StatusCode == http.StatusAccepted {

		if len(body) > 0 {
			err = json.Unmarshal(body, svcinstance)

			if err != nil {
				glog.Error(err)
				return nil, err
			}
		}
	} else {
		return nil, fmt.Errorf("%d returned from broker %s", resp.StatusCode, sb.Url)
	}
	glog.Infof("%v,%+v\n", string(body), svcinstance)
	return svcinstance, nil
}

func servicebroker_binding(param *ServiceBinding, binding_guid string, sb *ServiceBroker) (*ServiceBindingResponse, error) {
	jsonData, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = basicAuthStr(sb.UserName, sb.Password)

	resp, err := commToServiceBroker("PUT", "http://"+sb.Url+"/v2/service_instances/"+param.svc_instance_id+"/service_bindings/"+binding_guid, jsonData, header)
	if err != nil {

		glog.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	glog.Infof("respcode from PUT /v2/service_instances/%s/service_bindings/%s: %v", param.svc_instance_id, binding_guid, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	svcBinding := &ServiceBindingResponse{}

	glog.Infof("%v,%+v\n", string(body), svcBinding)
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		if len(body) > 0 {
			err = json.Unmarshal(body, svcBinding)

			if err != nil {
				glog.Error(err)
				return nil, err
			}
		}
	}

	return svcBinding, nil
}

func servicebroker_unbinding(bindId string, svc *ServiceBinding, sb *ServiceBroker) (interface{}, error) {

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = basicAuthStr(sb.UserName, sb.Password)

	resp, err := commToServiceBroker("DELETE", "http://"+sb.Url+"/v2/service_instances/"+svc.svc_instance_id+"/service_bindings/"+bindId+"?service_id="+svc.ServiceId+"&plan_id="+svc.PlanId, nil, header)
	if err != nil {

		glog.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	glog.Infof("respcode from DELETE /v2/service_instances/%s/service_bindings/%s: %v", svc.svc_instance_id, bindId, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	type UnBindindResp struct {
		Response interface{}
	}
	svcUnbinding := &UnBindindResp{}

	if resp.StatusCode == http.StatusOK {
		if len(body) > 0 {
			err = json.Unmarshal(body, svcUnbinding)

			if err != nil {
				glog.Error(err)
				return nil, err
			}
		}
	}
	glog.Infof("%v,%+v\n", string(body), svcUnbinding)
	return svcUnbinding, nil
}

func servicebroker_deprovisioning(bsi *backingserviceinstanceapi.BackingServiceInstance, sb *ServiceBroker) (interface{}, error) {

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = basicAuthStr(sb.UserName, sb.Password)

	resp, err := commToServiceBroker("DELETE", "http://"+sb.Url+"/v2/service_instances/"+bsi.Spec.InstanceID+"?service_id="+bsi.Spec.BackingServiceSpecID+"&plan_id="+bsi.Spec.BackingServicePlanGuid, nil, header)
	if err != nil {

		glog.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	glog.Infof("respcode from DELETE /v2/service_instances/%s: %v", bsi.Spec.InstanceID, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	type DeprovisioningResp struct {
		Response interface{}
	}
	svcDeprovisioning := &DeprovisioningResp{}

	if resp.StatusCode == http.StatusOK {
		if len(body) > 0 {
			err = json.Unmarshal(body, svcDeprovisioning)

			if err != nil {
				glog.Error(err)
				return nil, err
			}
		}
	}
	glog.Infof("%v,%+v\n", string(body), svcDeprovisioning)
	return svcDeprovisioning, nil
}

func basicAuthStr(username, password string) string {
	auth := username + ":" + password
	authstr := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + authstr
}

var InvalidCharFinder = regexp.MustCompile("[^a-zA-Z0-9]")

func deploymentconfig_env_prefix(bsName, bsiName string) string {
	return strings.ToUpper(fmt.Sprintf("BSI_%s_%s_", InvalidCharFinder.ReplaceAllLiteralString(bsName, ""), InvalidCharFinder.ReplaceAllLiteralString(bsiName, "")))
}

func deploymentconfig_env_name(prefix string, envName string) string {
	return strings.ToUpper(fmt.Sprintf("%s%s", prefix, InvalidCharFinder.ReplaceAllLiteralString(envName, "_")))
}

func (c *BackingServiceInstanceController) deploymentconfig_inject_envs(dc string, bsi *backingserviceinstanceapi.BackingServiceInstance, b *backingserviceinstanceapi.InstanceBinding) error {
	return c.deploymentconfig_modify_envs(dc, bsi, b, true)
}

func (c *BackingServiceInstanceController) deploymentconfig_clear_envs(dc string, bsi *backingserviceinstanceapi.BackingServiceInstance, b *backingserviceinstanceapi.InstanceBinding) error {
	return c.deploymentconfig_modify_envs(dc, bsi, b, false)
}


func (c *BackingServiceInstanceController) deploymentconfig_inject_envs_ups(dc string, bsi *backingserviceinstanceapi.BackingServiceInstance, b *backingserviceinstanceapi.InstanceBinding) error {
	return c.deploymentconfig_modify_envs_ups(dc, bsi, b, true)
}
func (c *BackingServiceInstanceController) deploymentconfig_clear_envs_ups(dc string, bsi *backingserviceinstanceapi.BackingServiceInstance, b *backingserviceinstanceapi.InstanceBinding) error {
	return c.deploymentconfig_modify_envs_ups(dc, bsi, b, false)
}


func (c *BackingServiceInstanceController) check_dc_healthy(bsi *backingserviceinstanceapi.BackingServiceInstance) {
	if bsi.Spec.Bound < 1{
		return
	}
	for _, binding := range bsi.Spec.Binding{
		dc, err := c.Client.DeploymentConfigs(bsi.Namespace).Get(binding.BindDeploymentConfig)
		if err != nil {
			if kerrors.IsNotFound(err) {
				glog.Infof("dc %s is not found.", binding.BindDeploymentConfig)
			}else{
				glog.Error(err.Error())
			}
		}else{
			if dc.Annotations["backingservice.instance/"+bsi.Name] != "bound"{
				glog.Infof("rebind envs in to dc %s",binding.BindDeploymentConfig)
				if bsi.Annotations[backingserviceinstanceapi.UPS] == "true" {
					err = c.deploymentconfig_inject_envs_ups(binding.BindDeploymentConfig, bsi, &binding)
					if err != nil {
						glog.Error(err.Error())
					}
				}else{
					err = c.deploymentconfig_inject_envs(binding.BindDeploymentConfig, bsi, &binding)
					if err != nil {
						glog.Error(err.Error())
					}
				}
			}
		}
	}
	return
}

// return exists or not
func env_get(envs []kapi.EnvVar, envName string) (bool, string) {
	for i := len(envs) - 1; i >= 0; i-- {
		if envs[i].Name == envName {
			return true, envs[i].Value
		}
	}

	return false, ""
}

// return overritten or not
func env_set(envs []kapi.EnvVar, envName, envValue string) (bool, []kapi.EnvVar) {
	if envs == nil {
		envs = []kapi.EnvVar{}
	}

	for i := len(envs) - 1; i >= 0; i-- {
		if envs[i].Name == envName {
			envs[i] = kapi.EnvVar{Name: envName, Value: envValue}
			return true, envs
		}
	}

	envs = append(envs, kapi.EnvVar{Name: envName, Value: envValue})
	return false, envs
}

// return unset or not
func env_unset(envs []kapi.EnvVar, envName string) (bool, []kapi.EnvVar) {
	if envs == nil {
		return false, nil
	}

	n := len(envs)
	index := 0
	for i := 0; i < n; i++ {
		if envs[i].Name != envName {
			if index < i {
				envs[index] = envs[i]
			}
			index++
		}
	}

	return index < n, envs[:index]
}

func (c *BackingServiceInstanceController) deploymentconfig_modify_envs_ups(dcname string, bsi *backingserviceinstanceapi.BackingServiceInstance, binding *backingserviceinstanceapi.InstanceBinding, toInject bool) error {
	dc, err := c.Client.DeploymentConfigs(bsi.Namespace).Get(dcname)
	if err != nil {
		return err
	}

	if dc.Spec.Template == nil {
		return nil
	}

	env_prefix := deploymentconfig_env_prefix(bsi.Spec.BackingServiceName, bsi.Name)
	containers := dc.Spec.Template.Spec.Containers
	if toInject {


		vsp := &VcapServiceParameters{
				Name:        bsi.Name,
				Label:       "",
				Plan:        backingserviceinstanceapi.UPS,
				Credentials: binding.Credentials,
			}


		for i := range containers {
			for k, v := range binding.Credentials {
				_, containers[i].Env = env_set(containers[i].Env, deploymentconfig_env_name(env_prefix, k), v)
			}

			if vsp != nil {
				_, containers[i].Env = modifyVcapServicesEnvNameEnv(containers[i].Env, backingserviceinstanceapi.UPS, vsp, "")
			}
		}
		if dc.Annotations == nil{
			dc.Annotations = make(map[string]string)
		}
		dc.Annotations["backingservice.instance/"+bsi.Name] = "bound"
	} else {
		for i := range containers {
			for k := range binding.Credentials {
				_, containers[i].Env = env_unset(containers[i].Env, deploymentconfig_env_name(env_prefix, k))
			}

			_, containers[i].Env = modifyVcapServicesEnvNameEnv(containers[i].Env, backingserviceinstanceapi.UPS, nil, bsi.Name)
		}
		delete(dc.Annotations, "backingservice.instance/"+bsi.Name)
	}

	if _, err := c.Client.DeploymentConfigs(bsi.Namespace).Update(dc); err != nil {
		return err
	}

	c.deploymentconfig_print_envs(bsi.Namespace, binding)
	return nil
}


func (c *BackingServiceInstanceController) deploymentconfig_modify_envs(dcname string, bsi *backingserviceinstanceapi.BackingServiceInstance, binding *backingserviceinstanceapi.InstanceBinding, toInject bool) error {
	dc, err := c.Client.DeploymentConfigs(bsi.Namespace).Get(dcname)
	if err != nil {
		return err
	}

	if dc.Spec.Template == nil {
		return nil
	}

		bs, err := c.Client.BackingServices("openshift").Get(bsi.Spec.BackingServiceName)
		if err != nil {
			return err
		}

		var plan = (*backingserviceapi.ServicePlan)(nil)
		for k := range bs.Spec.Plans {
			if bsi.Spec.BackingServicePlanGuid == bs.Spec.Plans[k].Id {
				plan = &(bs.Spec.Plans[k])
			}
		}


	env_prefix := deploymentconfig_env_prefix(bsi.Spec.BackingServiceName, bsi.Name)
	containers := dc.Spec.Template.Spec.Containers

	if toInject {
		var vsp *VcapServiceParameters = nil
		if plan != nil {
			vsp = &VcapServiceParameters{
				Name:        bsi.Name,
				Label:       "",
				Plan:        plan.Name,
				Credentials: binding.Credentials,
			}
		}

		for i := range containers {
			for k, v := range binding.Credentials {
				_, containers[i].Env = env_set(containers[i].Env, deploymentconfig_env_name(env_prefix, k), v)
			}

			if vsp != nil {
				_, containers[i].Env = modifyVcapServicesEnvNameEnv(containers[i].Env, bs.Name, vsp, "")
			}
		}
		if dc.Annotations == nil{
			dc.Annotations = make(map[string]string)
		}
		dc.Annotations["backingservice.instance/"+bsi.Name] = "bound"
	} else {
		for i := range containers {
			for k := range binding.Credentials {
				_, containers[i].Env = env_unset(containers[i].Env, deploymentconfig_env_name(env_prefix, k))
			}

			_, containers[i].Env = modifyVcapServicesEnvNameEnv(containers[i].Env, bs.Name, nil, bsi.Name)
		}
		delete(dc.Annotations, "backingservice.instance/"+bsi.Name)
	}

	if _, err := c.Client.DeploymentConfigs(bsi.Namespace).Update(dc); err != nil {
		return err
	}

	c.deploymentconfig_print_envs(bsi.Namespace, binding)

	return nil
}

func modifyVcapServicesEnvNameEnv(env []kapi.EnvVar, bsName string, vsp *VcapServiceParameters, bsiName string) (bool, []kapi.EnvVar) {
	_, json_env := env_get(env, VcapServicesEnvName)

	vs := VcapServices{}
	if len(strings.TrimSpace(json_env)) > 0 {
		err := json.Unmarshal([]byte(json_env), &vs)
		if err != nil {
			glog.Warningln("unmarshalVcapServices error: ", err.Error())
		}
	}

	if vsp != nil {
		vs = addVcapServiceParameters(vs, bsName, vsp)
	}
	if bsiName != "" {
		vs = removeVcapServiceParameters(vs, bsName, bsiName)
	}

	if len(vs) == 0 {
		return env_unset(env, VcapServicesEnvName)
	}
	json_data, err := json.Marshal(vs)
	if err != nil {
		glog.Warningln("marshalVcapServices error: ", err.Error())
		return false, env
	}

	json_env = string(json_data)

	glog.Info("new ", VcapServicesEnvName, " = ", json_env)

	return env_set(env, VcapServicesEnvName, json_env)
}

const VcapServicesEnvName = "VCAP_SERVICES"

type VcapServices map[string][]*VcapServiceParameters

type VcapServiceParameters struct {
	Name        string            `json:"name, omitempty"`
	Label       string            `json:"label, omitempty"`
	Plan        string            `json:"plan, omitempty"`
	Credentials map[string]string `json:"credentials, omitempty"`
}

func addVcapServiceParameters(vs VcapServices, serviceName string, vsParameters *VcapServiceParameters) VcapServices {
	if vs == nil {
		vs = VcapServices{}
	}

	if vsParameters == nil {
		return vs
	}

	removeVcapServiceParameters(vs, serviceName, vsParameters.Name)

	vsp_list := vs[serviceName]
	if vsp_list == nil {
		vsp_list = []*VcapServiceParameters{}
	}
	vsp_list = append(vsp_list, vsParameters)
	vs[serviceName] = vsp_list

	return vs
}

func removeVcapServiceParameters(vs VcapServices, serviceName, instanceName string) VcapServices {
	if vs == nil {
		vs = VcapServices{}
	}

	vsp_list := vs[serviceName]
	if len(vsp_list) == 0 {
		return vs
	}
	num := len(vsp_list)
	vsp_list2 := make([]*VcapServiceParameters, 0, num)
	for i := 0; i < num; i++ {
		vsp := vsp_list[i]
		if vsp != nil && vsp.Name != instanceName {
			vsp_list2 = append(vsp_list2, vsp)
		}
	}
	if len(vsp_list2) == 0 {
		delete(vs, serviceName)
	} else {
		vs[serviceName] = vsp_list2
	}

	return vs
}

func (c *BackingServiceInstanceController) deploymentconfig_print_envs(ns string, binding *backingserviceinstanceapi.InstanceBinding) {
	dc, err := c.Client.DeploymentConfigs(ns).Get(binding.BindDeploymentConfig)
	if err != nil {
		fmt.Println("dc not found: ", binding.BindDeploymentConfig)
		return
	}

	if dc.Spec.Template == nil {
		fmt.Println("dc.Spec.Template is nil")
		return
	}

	containers := dc.Spec.Template.Spec.Containers

	for _, c := range containers {
		fmt.Println("**********  envs in container")

		for _, env := range c.Env {
			fmt.Println("     env[", env.Name, ",] = ", env.Value)
		}
	}
}

func (c *BackingServiceInstanceController) deleteInstance(bs *backingserviceapi.BackingService, bsi *backingserviceinstanceapi.BackingServiceInstance) (result error) {
	glog.Infoln("bsi to delete ", bsi.Name)

	servicebroker, err := servicebroker_load(c.Client, bs.GenerateName)
	if err != nil {
		return err

	}

	glog.Infoln("deleting ", bsi.Name)
	if _, err = servicebroker_deprovisioning(bsi, servicebroker); err != nil {
		return err

	}

	glog.Infoln("bsi deleted ", bsi.Name)

	bsi.Status.Phase = backingserviceinstanceapi.BackingServiceInstancePhaseDeleted

	bsi.Status.Action = remove_action_word(bsi.Status.Action, backingserviceinstanceapi.BackingServiceInstanceActionToDelete)
	return

}


func (c *BackingServiceInstanceController) bindInstanceUPS(dc string,  bsi *backingserviceinstanceapi.BackingServiceInstance) (err error) {
	glog.Infoln(backingserviceinstanceapi.UPS, "bsi to bind ", bsi.Name, " and ", dc)

	instanceBinding := backingserviceinstanceapi.InstanceBinding{}
	now := unversioned.Now()
	instanceBinding.BoundTime = &now //&unversioned.Now()
	instanceBinding.BindUuid = backingserviceinstanceapi.UPS
	instanceBinding.BindDeploymentConfig = dc
	instanceBinding.Credentials = bsi.Spec.Credentials


	/*
	instanceBinding.Credentials = make(map[string]string)
	for k, v := range bsi.Spec.Credentials{
		instanceBinding.Credentials[k] = v
	}
	*/

	glog.Infoln("deploymentconfig_inject_envs")

	err = c.deploymentconfig_inject_envs_ups(dc, bsi, &instanceBinding)
	if err != nil {
		return err
	} else {
		bsi.Spec.Binding = append(bsi.Spec.Binding, instanceBinding)
	}

	glog.Infoln("bsi bound. ", bsi.Name)

	bsi.Spec.Bound += 1

	bsi.Status.Phase = backingserviceinstanceapi.BackingServiceInstancePhaseBound
	bsi.Annotations[dc] = backingserviceinstanceapi.BindDeploymentConfigBound

	bsi.Status.Action = "" //remove_action_word(bsi.Status.Action, backingserviceinstanceapi.BackingServiceInstanceActionToBind)
	return


}


func (c *BackingServiceInstanceController) bindInstance(dc string, bs *backingserviceapi.BackingService, bsi *backingserviceinstanceapi.BackingServiceInstance) (result error) {
	glog.Infoln("bsi to bind ", bsi.Name, " and ", dc)

	servicebroker, err := servicebroker_load(c.Client, bs.GenerateName)
	if err != nil {
		return err
	}

	bind_uuid := string(util.NewUUID())

	servicebinding := &ServiceBinding{
		ServiceId: bs.Spec.Id,
		PlanId:    bsi.Spec.BackingServicePlanGuid,
		AppGuid:   bsi.Namespace,
		//BindResource: ,
		//Parameters: ,
		svc_instance_id: bsi.Spec.InstanceID,
	}

	glog.Infoln("bsi to bind", bsi.Name)

	bindingresponse, err := servicebroker_binding(servicebinding, bind_uuid, servicebroker)
	if err != nil {
		return err
	}

	instanceBinding := backingserviceinstanceapi.InstanceBinding{}
	now := unversioned.Now()
	instanceBinding.BoundTime = &now //&unversioned.Now()
	instanceBinding.BindUuid = bind_uuid
	instanceBinding.BindDeploymentConfig = dc
	instanceBinding.Credentials = make(map[string]string)
	instanceBinding.Credentials["Uri"] = bindingresponse.Credentials.Uri
	instanceBinding.Credentials["Name"] = bindingresponse.Credentials.Name
	instanceBinding.Credentials["Username"] = bindingresponse.Credentials.Username
	instanceBinding.Credentials["Password"] = bindingresponse.Credentials.Password
	instanceBinding.Credentials["Host"] = bindingresponse.Credentials.Host
	instanceBinding.Credentials["Port"] = bindingresponse.Credentials.Port
	instanceBinding.Credentials["Vhost"] = bindingresponse.Credentials.Vhost
	// = bindingresponse.SyslogDrainUrl
	// = bindingresponse.RouteServiceUrl

	glog.Infoln("deploymentconfig_inject_envs")

	err = c.deploymentconfig_inject_envs(dc, bsi, &instanceBinding)
	if err != nil {
		return err
	} else {
		bsi.Spec.Binding = append(bsi.Spec.Binding, instanceBinding)
	}

	glog.Infoln("bsi bound. ", bsi.Name)

	bsi.Spec.Bound += 1

	bsi.Status.Phase = backingserviceinstanceapi.BackingServiceInstancePhaseBound
	bsi.Annotations[dc] = backingserviceinstanceapi.BindDeploymentConfigBound

	bsi.Status.Action = "" //remove_action_word(bsi.Status.Action, backingserviceinstanceapi.BackingServiceInstanceActionToBind)
	return

}


func (c *BackingServiceInstanceController) unbindInstanceUPS(dc string, bsi *backingserviceinstanceapi.BackingServiceInstance) (err error) {

	glog.Infoln(backingserviceinstanceapi.UPS," bsi to unbind ", bsi.Name)


	for idx, b := range bsi.Spec.Binding {
		if b.BindDeploymentConfig == dc {

			glog.Infoln("deploymentconfig_clear_envs")
			err = c.deploymentconfig_clear_envs_ups(dc, bsi, &b)
			if err != nil {
				return err
			} else {
				bsi.Spec.Binding = append(bsi.Spec.Binding[:idx], bsi.Spec.Binding[idx+1:]...)
				delete(bsi.Annotations, dc)
			}
			break
		}
	}
	/*
		err = c.deploymentconfig_clear_envs(bsi.Namespace, dc)
		if err != nil {
			return err
		}
	*/
	glog.Infoln("bsi is unbound ", bsi.Name)
	/*
		delete(bsi.Annotations, dc)
		bsi.Spec.BindDeploymentConfig = ""
		bsi.Spec.Credentials = nil
		bsi.Spec.BoundTime = nil
		bsi.Spec.BindUuid = ""
		bsi.Spec.Bound = false
	*/
	bsi.Spec.Bound -= 1
	if bsi.Spec.Bound == 0 {
		bsi.Status.Phase = backingserviceinstanceapi.BackingServiceInstancePhaseUnbound
	}

	bsi.Status.Action = "" //remove_action_word(bsi.Status.Action, backingserviceinstanceapi.BackingServiceInstanceActionToUnbind)

	return

}

func (c *BackingServiceInstanceController) unbindInstance(dc string, bs *backingserviceapi.BackingService, bsi *backingserviceinstanceapi.BackingServiceInstance) (result error) {

	glog.Infoln("bsi to unbind ", bsi.Name)

	servicebroker, err := servicebroker_load(c.Client, bs.GenerateName)
	if err != nil {
		return err
	}

	glog.Infoln("servicebroker_unbinding")

	svc := &ServiceBinding{}

	svc.PlanId = bsi.Spec.BackingServicePlanGuid
	svc.ServiceId = bsi.Spec.BackingServiceSpecID
	svc.svc_instance_id = bsi.Spec.InstanceID

	var bindId string

	for idx, b := range bsi.Spec.Binding {
		if b.BindDeploymentConfig == dc {
			bindId = b.BindUuid
			_, err = servicebroker_unbinding(bindId, svc, servicebroker)
			if err != nil {
				return err
			}
			glog.Infoln("deploymentconfig_clear_envs")
			err = c.deploymentconfig_clear_envs(dc, bsi, &b)
			if err != nil {
				return err
			} else {
				bsi.Spec.Binding = append(bsi.Spec.Binding[:idx], bsi.Spec.Binding[idx+1:]...)
				delete(bsi.Annotations, dc)
			}
			break
		}
	}
	/*
		err = c.deploymentconfig_clear_envs(bsi.Namespace, dc)
		if err != nil {
			return err
		}
	*/
	glog.Infoln("bsi is unbound ", bsi.Name)
	/*
		delete(bsi.Annotations, dc)
		bsi.Spec.BindDeploymentConfig = ""
		bsi.Spec.Credentials = nil
		bsi.Spec.BoundTime = nil
		bsi.Spec.BindUuid = ""
		bsi.Spec.Bound = false
	*/
	bsi.Spec.Bound -= 1
	if bsi.Spec.Bound == 0 {
		bsi.Status.Phase = backingserviceinstanceapi.BackingServiceInstancePhaseUnbound
	}

	bsi.Status.Action = "" //remove_action_word(bsi.Status.Action, backingserviceinstanceapi.BackingServiceInstanceActionToUnbind)

	return

}

func (c *BackingServiceInstanceController) get_deploymentconfig_name(bsi *backingserviceinstanceapi.BackingServiceInstance, stat string) string {
	for dc, bound := range bsi.Annotations {
		if bound == stat {
			return dc
		}
	}
	return ""
}
