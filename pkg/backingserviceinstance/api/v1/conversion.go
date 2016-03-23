package v1

import (

	"k8s.io/kubernetes/pkg/runtime"

	oapi "github.com/openshift/origin/pkg/api"
	"github.com/openshift/origin/pkg/backingserviceinstance/api"
)

func addConversionFuncs(scheme *runtime.Scheme) {
	if err := scheme.AddFieldLabelConversionFunc("v1", "BackingServiceInstance",
		oapi.GetFieldLabelConversionFunc(api.BackingServiceInstanceToSelectableFields(&api.BackingServiceInstance{}), nil),
	); err != nil {
		panic(err)
	}
}

