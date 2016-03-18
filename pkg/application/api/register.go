package api

import (
	"k8s.io/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("",
		&Application{},
		&ApplicationList{},
	)
}

func (*Application) IsAnAPIObject()     {}
func (*ApplicationList) IsAnAPIObject() {}
