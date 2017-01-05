package v1

import (
	"k8s.io/kubernetes/pkg/runtime"

	oapi "github.com/openshift/origin/pkg/api"
	"github.com/openshift/origin/pkg/application/api"
)

func addConversionFuncs(scheme *runtime.Scheme) {
	if err := scheme.AddFieldLabelConversionFunc("v1", "Application",
		oapi.GetFieldLabelConversionFunc(api.ApplicationToSelectableFields(&api.Application{}), nil),
	); err != nil {
		panic(err)
	}
}
