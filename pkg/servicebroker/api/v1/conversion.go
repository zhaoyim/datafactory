package v1

import (
	"k8s.io/kubernetes/pkg/runtime"

	oapi "github.com/openshift/origin/pkg/api"
	"github.com/openshift/origin/pkg/servicebroker/api"
)

func addConversionFuncs(scheme *runtime.Scheme) {
	if err := scheme.AddFieldLabelConversionFunc("v1", "ServiceBroker",
		oapi.GetFieldLabelConversionFunc(api.ServiceBrokerToSelectableFields(&api.ServiceBroker{}), nil),
	); err != nil {
		panic(err)
	}

}
