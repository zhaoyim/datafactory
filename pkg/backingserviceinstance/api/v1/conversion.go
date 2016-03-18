package v1

import (
	kapi "k8s.io/kubernetes/pkg/api"

	//"k8s.io/kubernetes/pkg/registry/namespace"

	oapi "github.com/openshift/origin/pkg/api"
	newer "github.com/openshift/origin/pkg/backingserviceinstance/api"
)

func init() {
	if err := kapi.Scheme.AddFieldLabelConversionFunc("v1", "BackingServiceInstance",
		oapi.GetFieldLabelConversionFunc(newer.BackingServiceInstanceToSelectableFields(&newer.BackingServiceInstance{}), nil),
	); err != nil {
		panic(err)
	}
}
