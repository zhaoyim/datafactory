package api

import (
	"k8s.io/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("",
		&BackingService{},
		&BackingServiceList{},
	)
}

func (*BackingService) IsAnAPIObject()     {}
func (*BackingServiceList) IsAnAPIObject() {}
