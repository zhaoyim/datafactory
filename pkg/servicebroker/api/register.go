package api

import (
	"k8s.io/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("",
		&ServiceBroker{},
		&ServiceBrokerList{},
	)
}

func (*ServiceBroker) IsAnAPIObject()     {}
func (*ServiceBrokerList) IsAnAPIObject() {}
