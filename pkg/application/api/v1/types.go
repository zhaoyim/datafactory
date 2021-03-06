package v1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	kapi "k8s.io/kubernetes/pkg/api/v1"
)

const (
	// These are internal finalizer values to Origin
	FinalizerOrigin kapi.FinalizerName = "openshift.io/origin"

	// ApplicationNew is create by administrator.
	ApplicationNew ApplicationPhase = "New"

	// ApplicationRunning indicates that Application service working well.
	ApplicationActive       ApplicationPhase = "Active"
	ApplicationActiveUpdate ApplicationPhase = "ActiveUpdate"

	// ApplicationDeletingLabel indicates that Application is delete label.
	ApplicationTerminatingLabel ApplicationPhase = "Terminating Label"

	// ApplicationFailed indicates that Application is delete all resources.
	ApplicationTerminating ApplicationPhase = "Terminating"

	// ApplicationFailed indicates that Application is checking
	ApplicationChecking ApplicationPhase = "Checking"

	ApplicationItemDelete string = "Resource Not Found"

	ApplicationItemLabelDelete string = "Resource Label Not Found"
)

var ApplicationItemSupportKinds = []string{
	"Build", "BuildConfig", "DeploymentConfig", "ImageStream", "ImageStreamTag", "ImageStreamImage", //openshift kind
	"Event", "Node", "Job", "Pod", "ReplicationController", "Service", "PersistentVolume", "PersistentVolumeClaim", //k8s kind
	"ServiceBroker", "BackingServiceInstance",
}

type ApplicationPhase string

// Application describe an Application
type Application struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	kapi.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of the Namespace.
	Spec ApplicationSpec `json:"spec,omitempty" description:"spec defines the behavior of the Application"`

	// Status describes the current status of a Namespace
	Status ApplicationStatus `json:"status,omitempty" description:"status describes the current status of a Application; read-only"`
}

// ApplicationList describe an ApplicationList
type ApplicationList struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	unversioned.ListMeta `json:"metadata,omitempty"`

	// Items is a list of applications
	Items []Application `json:"items" description:"list of Applications"`
}

// ApplicationSpec describes the attributes on a Application
type ApplicationSpec struct {
	// name defines the name of a Application
	Name string `json:"name" description:"name defines the name of a Application"`
	// items defines the resources to be labeled in a Application
	Items ItemList `json:"items" description:"items defines the resources to be labeled in a Application"`
	//destory defines the resources to be removed in a Application
	Destory bool `json:"destoryOption" description:"destory defines the resources to be removed in a Application"`
	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage
	Finalizers []kapi.FinalizerName `json:"finalizers,omitempty" description:"an opaque list of values that must be empty to permanently remove object from storage"`
}

// ApplicationStatus is information about the current status of a Application
type ApplicationStatus struct {
	// phase is the current lifecycle phase of the Application
	Phase ApplicationPhase `json:"phase,omitempty" description:"phase is the current lifecycle phase of the Application"`
}

type ItemList []Item

// Item  describe an application item
type Item struct {
	// kind defines the item kind of a item in Application
	Kind string `json:"kind" description:"kind defines the item kind of a item in Application"`
	// name defines the item name of a item in Application
	Name string `json:"name" description:"name defines the item name of a item in Application"`
	// status defines a operate to the item label
	Status string `json:"status,omitempty" description:"status defines a operate to the item label"`
}

const (
	ApplicationItemStatusAdd    = "add"
	ApplicationItemStatusDelete = "delete"
	ApplicationItemStatusErr    = "error"
	ApplicationItemStatusOk     = "ok"
)
