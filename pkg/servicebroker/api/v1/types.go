package v1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	kapi "k8s.io/kubernetes/pkg/api/v1"
)

const (
	// These are internal finalizer values to Origin
	FinalizerOrigin kapi.FinalizerName = "openshift.io/origin"

	// ServiceBrokerNew is create by administrator.
	ServiceBrokerNew ServiceBrokerPhase = "New"

	// ServiceBrokerRunning indicates that servicebroker service working well.
	ServiceBrokerActive ServiceBrokerPhase = "Active"

	// ServiceBrokerFailed indicates that servicebroker stopped.
	ServiceBrokerFailed ServiceBrokerPhase = "Failed"

	// ServiceBrokerDeleting indicates that servicebroker is going to be deleted.
	ServiceBrokerDeleting ServiceBrokerPhase = "Deleting"

	// ServiceBrokerLastPingTime indicates that servicebroker last ping time.
	PingTimer string = "ServiceBroker/LastPing"

	// ServiceBrokerNewRetryTimes indicates that new servicebroker retry times.
	ServiceBrokerNewRetryTimes string = "ServiceBroker/NewRetryTimes"

	// ServiceBrokerLastRefreshBSTime indicates that servicebroker last refresh backingservice time.
	RefreshTimer string = "ServiceBroker/LastRefresh"
)

type ServiceBrokerPhase string

// ServiceBroker describe a servicebroker
type ServiceBroker struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	kapi.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the Namespace.
	Spec ServiceBrokerSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" description:"spec defines the behavior of the ServiceBroker"`

	// Status describes the current status of a Namespace
	Status ServiceBrokerStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" description:"status describes the current status of a ServiceBroker; read-only"`
}

// ServiceBrokerList is a list of ServiceBroker objects.
type ServiceBrokerList struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	unversioned.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is a list of routes
	Items []ServiceBroker `json:"items" protobuf:"bytes,2,rep,name=items" description:"list of servicebrokers"`
}

// ServiceBrokerSpec describes the attributes on a ServiceBroker
type ServiceBrokerSpec struct {
	// url defines the address of a ServiceBroker service
	Url string `json:"url" protobuf:"bytes,1,opt,name=url" description:"url defines the address of a ServiceBroker service"`
	// name defines the name of a ServiceBroker service
	Name string `json:"name" protobuf:"bytes,2,opt,name=name" description:"name defines the name of a ServiceBroker service"`
	// username defines the username to access ServiceBroker service
	UserName string `json:"username" protobuf:"bytes,3,opt,name=username" description:"username defines the username to access ServiceBroker service"`
	// password defines the password to access ServiceBroker service
	Password string `json:"password" protobuf:"bytes,4,opt,name=password" description:"password defines the password to access ServiceBroker service"`
	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage
	Finalizers []kapi.FinalizerName `json:"finalizers,omitempty" protobuf:"bytes,5,rep,name=finalizers" description:"an opaque list of values that must be empty to permanently remove object from storage"`
}

// ServiceBrokerStatus is information about the current status of a ServiceBroker
type ServiceBrokerStatus struct {
	// Phase is the current lifecycle phase of the project
	Phase ServiceBrokerPhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase" description:"phase is the current lifecycle phase of the servicebroker"`
}

const (
	ServiceBrokerLabel = "asiainfo.io/servicebroker"
)
