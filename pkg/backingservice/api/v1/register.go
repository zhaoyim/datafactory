package v1

import (
	"k8s.io/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("v1",
		&BackingService{},
		&BackingServiceList{},
	)
}

func (*BackingService) IsAnAPIObject()     {}
func (*BackingServiceList) IsAnAPIObject() {}
