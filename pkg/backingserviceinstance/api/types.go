package api

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
)

type BackingServiceInstance struct {
	unversioned.TypeMeta
	kapi.ObjectMeta

	// Spec defines the behavior of the Namespace.
	Spec BackingServiceInstanceSpec

	// Status describes the current status of a Namespace
	Status BackingServiceInstanceStatus
}

type BackingServiceInstanceList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	// Items is a list of routes
	Items []BackingServiceInstance
}

/*
type BackingServiceInstanceSpec struct {
	Config                 map[string]string `json:"config, omitempty"`
	DashboardUrl           string            `json:"dashboard_url, omitempty"`
	BackingServiceGuid     string            `json:"backingservice_guid, omitempty"`
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
	InstanceProvisioning
	Binding    []InstanceBinding
	Bound      int
	Tags       []string
	InstanceID string
	// InstanceID is blank means to delete (when len(Parameters) > 0)
}

/*
type InstanceProvisioning struct {
	DashboardUrl           string
	BackingServiceName     string
	BackingServiceID       string
	BackingServicePlanGuid string
	BackingServicePlanName string
	Parameters             map[string]string
	// len(Parameters) == 0 means not inited
}
*/

type InstanceProvisioning struct {
	DashboardUrl           string
	BackingServiceName     string
	BackingServiceSpecID   string
	BackingServicePlanGuid string
	BackingServicePlanName string
	Parameters             map[string]string
	// len(Parameters) == 0 means not inited
}

type InstanceBinding struct {
	// BindUuid is blank for not bound (Bound=false) or to unbind (Bound=true)
	// BindUuid != "" and Bound=false means to bind
	BindUuid             string
	BoundTime            *unversioned.Time
	BindDeploymentConfig string
	Credentials          map[string]string
}

// ProjectStatus is information about the current status of a Project
type BackingServiceInstanceStatus struct {
	Phase  BackingServiceInstancePhase
	Action BackingServiceInstanceAction

	LastOperation *LastOperation
}

type LastOperation struct {
	State                    string
	Description              string
	AsyncPollIntervalSeconds int
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
	unversioned.TypeMeta
	kapi.ObjectMeta
	BindKind            string
	BindResourceVersion string
	ResourceName        string
}

func NewBindingRequestOptions(kind, version, name string) *BindingRequestOptions {
	return &BindingRequestOptions{
		BindKind:            kind,
		BindResourceVersion: version,
		ResourceName:        name,
	}
}
