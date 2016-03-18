package v1

import (
	"k8s.io/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("v1",
		&BackingServiceInstance{},
		&BackingServiceInstanceList{},
		//&BindingRequest{},
		&BindingRequestOptions{},
	)
}

func (*BackingServiceInstance) IsAnAPIObject()     {}
func (*BackingServiceInstanceList) IsAnAPIObject() {}
//func (*BindingRequest) IsAnAPIObject()             {}
func (*BindingRequestOptions) IsAnAPIObject()      {}
