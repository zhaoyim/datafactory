package v1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	kapi "k8s.io/kubernetes/pkg/api/v1"
)

// BackingService describe a BackingService
type BackingService struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	kapi.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the Namespace.
	Spec BackingServiceSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" description:"specification of the desired behavior for a BackingService"`

	// Status describes the current status of a Namespace
	Status BackingServiceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" description:"status describes the current status of a BackingService"`
}

// BackingServiceList describe a list of BackingService
type BackingServiceList struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	unversioned.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is a list of routes
	Items []BackingService `json:"items" protobuf:"bytes,2,rep,name=items" description:"list of backingservice"`
}

// BackingServiceSpec describe the attributes on a Backingservice
type BackingServiceSpec struct {
	// name of backingservice
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" description:"name of backingservice"`
	// id of backingservice
	Id string `json:"id" protobuf:"bytes,2,opt,name=id" description:"id of backingservice"`
	// description of a backingservice
	Description string `json:"description" protobuf:"bytes,3,opt,name=description" description:"description of a backingservice"`
	// is backingservice bindable
	Bindable bool `json:"bindable" protobuf:"varint,4,opt,name=bindable" description:"is backingservice bindable?"`
	// is  backingservice plan updateable
	PlanUpdateable bool `json:"plan_updateable, omitempty" protobuf:"varint,5,opt,name=plan_updateable" description:"is  backingservice plan updateable"`
	// list of backingservice tags of BackingService
	Tags []string `json:"tags, omitempty" protobuf:"bytes,6,rep,name=tags" description:"list of backingservice tags of BackingService"`
	// require condition of backingservice
	Requires []string `json:"requires, omitempty" protobuf:"bytes,7,rep,name=requires" description:"require condition of backingservice"`

	// metadata of backingservice
	Metadata map[string]string `json:"metadata, omitempty" protobuf:"bytes,8,rep,name=metadata" description:"metadata of backingservice"`
	// plans of a backingservice
	Plans []ServicePlan `json:"plans" protobuf:"bytes,9,rep,name=plans" description:"plans of a backingservice"`
	// DashboardClient of backingservic
	DashboardClient map[string]string `json:"dashboard_client" protobuf:"bytes,10,rep,name=dashboard_client" description:"DashboardClient of backingservice"`
}

// ServiceMetadata describe a ServiceMetadata
type ServiceMetadata struct {
	// displayname of a ServiceMetadata
	DisplayName string `json:"displayName, omitempty" protobuf:"bytes,1,opt,name=displayName"`
	// imageurl of a ServiceMetadata
	ImageUrl string `json:"imageUrl, omitempty" protobuf:"bytes,2,opt,name=imageUrl"`
	// long description of a ServiceMetadata
	LongDescription string `json:"longDescription, omitempty" protobuf:"bytes,3,opt,name=longDescription"`
	// providerdisplayname of a ServiceMetadata
	ProviderDisplayName string `json:"providerDisplayName, omitempty" protobuf:"bytes,4,opt,name=providerDisplayName"`
	// documrntation url of a ServiceMetadata
	DocumentationUrl string `json:"documentationUrl, omitempty" protobuf:"bytes,5,opt,name=documentationUrl"`
	// support url of a ServiceMetadata
	SupportUrl string `json:"supportUrl, omitempty" protobuf:"bytes,6,opt,name=supportUrl"`
}

// ServiceDashboardClient describe a ServiceDashboardClient
type ServiceDashboardClient struct {
	// id of a ServiceDashboardClient
	Id string `json:"id, omitempty" protobuf:"bytes,1,opt,name=id"`
	// secret of a ServiceDashboardClient
	Secret string `json:"secret, omitempty" protobuf:"bytes,2,opt,name=secret"`
	//redirect uri of a ServiceDashboardClient
	RedirectUri string `json:"redirect_uri, omitempty" protobuf:"bytes,3,opt,name=redirect_uri"`
}

// ServicePlan describe a ServicePlan
type ServicePlan struct {
	// name of a ServicePlan
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	//id of a ServicePlan
	Id string `json:"id" protobuf:"bytes,2,opt,name=id"`
	// description of a ServicePlan
	Description string `json:"description" protobuf:"bytes,3,opt,name=description"`
	// metadata of a ServicePlan
	Metadata ServicePlanMetadata `json:"metadata, omitempty" protobuf:"bytes,4,opt,name=metadata"`
	// is this plan free or not
	Free bool `json:"free, omitempty" protobuf:"varint,5,opt,name=free"`
}

// ServicePlanMetadata describe a ServicePlanMetadata
type ServicePlanMetadata struct {
	// bullets of a ServicePlanMetadata
	Bullets []string `json:"bullets, omitempty" protobuf:"bytes,1,rep,name=bullets"`
	// costs of a ServicePlanMetadata
	Costs []ServicePlanCost `json:"costs, omitempty" protobuf:"bytes,2,rep,name=costs"`
	// displayname of a ServicePlanMetadata
	DisplayName string `json:"displayName, omitempty" protobuf:"bytes,3,opt,name=displayName"`
}

//TODO amount should be a array object...

// ServicePlanCost describe a ServicePlanCost
type ServicePlanCost struct {
	// amount of a ServicePlanCost
	Amount map[string]float64 `json:"amount, omitempty" protobuf:"bytes,1,rep,name=monitors"`
	// unit of a ServicePlanCost
	Unit string `json:"unit, omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// ProjectStatus is information about the current status of a Project
type BackingServiceStatus struct {
	// phase is the current lifecycle phase of the servicebroker
	Phase BackingServicePhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase" description:"phase is the current lifecycle phase of the servicebroker"`
}

type BackingServicePhase string

const (
	BackingServicePhaseActive   BackingServicePhase = "Active"
	BackingServicePhaseInactive BackingServicePhase = "Inactive"
)
