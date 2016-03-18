package v1

import (
	"k8s.io/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("v1",
		&ServiceBroker{},
		&ServiceBrokerList{},
	)
}

func (*ServiceBroker) IsAnAPIObject()     {}
func (*ServiceBrokerList) IsAnAPIObject() {}
