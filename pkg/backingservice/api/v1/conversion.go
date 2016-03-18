package v1

import (
	kapi "k8s.io/kubernetes/pkg/api"

	//"k8s.io/kubernetes/pkg/registry/namespace"

	oapi "github.com/openshift/origin/pkg/api"
	newer "github.com/openshift/origin/pkg/backingservice/api"
)

func init() {
	if err := kapi.Scheme.AddFieldLabelConversionFunc("v1", "BackingService",
		oapi.GetFieldLabelConversionFunc(newer.BackingServiceToSelectableFields(&newer.BackingService{}), nil),
	); err != nil {
		panic(err)
	}
}
