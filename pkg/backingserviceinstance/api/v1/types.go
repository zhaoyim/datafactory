package v1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	kapi "k8s.io/kubernetes/pkg/api/v1"
)

type BackingServiceInstance struct {
	unversioned.TypeMeta `json:",inline"`
	kapi.ObjectMeta      `json:"metadata,omitempty"`

	// Spec defines the behavior of the Namespace.
	Spec BackingServiceInstanceSpec `json:"spec,omitempty" description:"spec defines the behavior of the BackingServiceInstance"`

	// Status describes the current status of a Namespace
	Status BackingServiceInstanceStatus `json:"status,omitempty" description:"status describes the current status of a Project; read-only"`
}

type BackingServiceInstanceList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`

	// Items is a list of routes
	Items []BackingServiceInstance `json:"items" description:"list of BackingServiceInstances"`
}

/*
type BackingServiceInstanceSpec struct {
	Config                 map[string]string `json:"config, omitempty"`
	InstanceID             string            `json:"instance_id, omitempty"`
	DashboardUrl           string            `json:"dashboard_url, omitempty"`
	BackingServiceName     string            `json:"backingservice_name, omitempty"`
	BackingServiceID       string            `json:"backingservice_id, omitempty"`
	BackingServicePlanGuid string            `json:"backingservice_plan_guid, omitempty"`
	Parameters             map[string]string `json:"parameters, omitempty"`
	Binding                bool              `json:"binding, omitempty"`
	BindUuid               string            `json:"bind_uuid, omitempty"`
	BindDeploymentConfig   map[string]string `json:"bind_deploymentconfig, omitempty"`
	Credential             map[string]string `json:"credential, omitempty"`
	Tags                   []string          `json:"tags, omitempty"`
}
*/
type BackingServiceInstanceSpec struct {
	InstanceProvisioning `json:"provisioning, omitempty"`
	Binding              []InstanceBinding `json:"binding, omitempty"`
	Bound                int               `json:"bound, omitempty"`
	InstanceID           string            `json:"instance_id, omitempty"`
	Tags                 []string          `json:"tags, omitempty"`
}

/*
type InstanceProvisioning struct {
	DashboardUrl           string            `json:"dashboard_url, omitempty"`
	BackingService         string            `json:"backingservice, omitempty"`
	BackingServiceName     string            `json:"backingservice_name, omitempty"`
	BackingServiceID       string            `json:"backingservice_id, omitempty"`
	BackingServicePlanGuid string            `json:"backingservice_plan_guid, omitempty"`
	BackingServicePlanName string            `json:"backingservice_plan_name, omitempty"`
	Parameters             map[string]string `json:"parameters, omitempty"`
}
*/

type InstanceProvisioning struct {
	DashboardUrl           string            `json:"dashboard_url, omitempty"`
	BackingServiceName     string            `json:"backingservice_name, omitempty"`
	BackingServiceSpecID   string            `json:"backingservice_spec_id, omitempty"`
	BackingServicePlanGuid string            `json:"backingservice_plan_guid, omitempty"`
	BackingServicePlanName string            `json:"backingservice_plan_name, omitempty"`
	Parameters             map[string]string `json:"parameters, omitempty"`
}

type InstanceBinding struct {
	BoundTime            *unversioned.Time `json:"bound_time,omitempty"`
	BindUuid             string            `json:"bind_uuid, omitempty"`
	BindDeploymentConfig string            `json:"bind_deploymentconfig, omitempty"`
	Credentials          map[string]string `json:"credentials, omitempty"`
}

type BackingServiceInstanceStatus struct {
	Phase  BackingServiceInstancePhase  `json:"phase, omitempty"`
	Action BackingServiceInstanceAction `json:"action, omitempty"`

	LastOperation *LastOperation `json:"last_operation, omitempty"`
}

type LastOperation struct {
	State                    string `json:"state"`
	Description              string `json:"description"`
	AsyncPollIntervalSeconds int    `json:"async_poll_interval_seconds, omitempty"`
}

type BackingServiceInstancePhase string
type BackingServiceInstanceAction string

const (
	BackingServiceInstancePhaseProvisioning BackingServiceInstancePhase = "Provisioning"
	BackingServiceInstancePhaseUnbound      BackingServiceInstancePhase = "Unbound"
	BackingServiceInstancePhaseBound        BackingServiceInstancePhase = "Bound"
	BackingServiceInstancePhaseDeleted      BackingServiceInstancePhase = "Deleted"

	BackingServiceInstanceActionToBind   BackingServiceInstanceAction = "_ToBind"
	BackingServiceInstanceActionToUnbind BackingServiceInstanceAction = "_ToUnbind"
	BackingServiceInstanceActionToDelete BackingServiceInstanceAction = "_ToDelete"

	BindDeploymentConfigBinding   string = "binding"
	BindDeploymentConfigUnbinding string = "unbinding"
	BindDeploymentConfigBound     string = "bound"
)

//=====================================================
//
//=====================================================

const BindKind_DeploymentConfig = "DeploymentConfig"

//type BindingRequest struct {
//	unversioned.TypeMeta
//	kapi.ObjectMeta
//
//	// the dc
//	DeploymentConfigName string `json:"deployment_name, omitempty"`
//}

type BindingRequestOptions struct {
	unversioned.TypeMeta `json:",inline"`
	kapi.ObjectMeta      `json:"metadata,omitempty"`

	BindKind            string `json:"bindKind, omitempty"`
	BindResourceVersion string `json:"bindResourceVersion, omitempty"`
	ResourceName        string `json:"resourceName, omitempty"`
}
