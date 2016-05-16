package api

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
)

const BSNS = "openshift"

type BackingService struct {
	unversioned.TypeMeta
	kapi.ObjectMeta

	// Spec defines the behavior of the Namespace.
	Spec BackingServiceSpec

	// Status describes the current status of a Namespace
	Status BackingServiceStatus
}

type BackingServiceList struct {
	unversioned.TypeMeta
	unversioned.ListMeta

	// Items is a list of routes
	Items []BackingService
}

type BackingServiceSpec struct {
	Name           string
	Id             string
	Description    string
	Bindable       bool
	PlanUpdateable bool
	Tags           []string
	Requires       []string

	//Metadata        ServiceMetadata
	Metadata        map[string]string
	Plans           []ServicePlan
	DashboardClient map[string]string
	//DashboardClient ServiceDashboardClient
}

type ServiceMetadata struct {
	DisplayName         string
	ImageUrl            string
	LongDescription     string
	ProviderDisplayName string
	DocumentationUrl    string
	SupportUrl          string
}

type ServiceDashboardClient struct {
	Id          string
	Secret      string
	RedirectUri string
}

type ServicePlan struct {
	Name        string
	Id          string
	Description string
	Metadata    ServicePlanMetadata
	Free        bool
}

type ServicePlanMetadata struct {
	Bullets     []string
	Costs       []ServicePlanCost
	DisplayName string
}

//TODO amount should be a array object...
type ServicePlanCost struct {
	Amount map[string]float64
	Unit   string
}

// ProjectStatus is information about the current status of a Project
type BackingServiceStatus struct {
	Phase BackingServicePhase
}

type BackingServicePhase string

const (
	BackingServicePhaseActive   BackingServicePhase = "Active"
	BackingServicePhaseInactive BackingServicePhase = "Inactive"
)
