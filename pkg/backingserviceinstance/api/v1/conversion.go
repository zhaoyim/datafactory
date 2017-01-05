package v1

import (
	"k8s.io/kubernetes/pkg/runtime"

	oapi "github.com/openshift/origin/pkg/api"
	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
)

func addConversionFuncs(scheme *runtime.Scheme) {
	if err := scheme.AddFieldLabelConversionFunc("v1", "BackingServiceInstance",
		oapi.GetFieldLabelConversionFunc(backingserviceinstanceapi.BackingServiceInstanceToSelectableFields(&backingserviceinstanceapi.BackingServiceInstance{}), nil),
	); err != nil {
		panic(err)
	}
}
