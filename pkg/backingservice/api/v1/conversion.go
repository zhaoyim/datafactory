package v1

import (
	"k8s.io/kubernetes/pkg/runtime"

	oapi "github.com/openshift/origin/pkg/api"
	"github.com/openshift/origin/pkg/backingservice/api"
)

func addConversionFuncs(scheme *runtime.Scheme) {
	if err := scheme.AddFieldLabelConversionFunc("v1", "BackingService",
		oapi.GetFieldLabelConversionFunc(api.BackingServiceToSelectableFields(&api.BackingService{}), nil),
	); err != nil {
		panic(err)
	}

}
