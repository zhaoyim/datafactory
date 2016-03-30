package v1

import (
	"testing"

	//kapi "k8s.io/kubernetes/pkg/api"

	api "github.com/openshift/origin/pkg/backingserviceinstance/api"
	testutil "github.com/openshift/origin/test/util/api"
)

func TestFieldSelectorConversions(t *testing.T) {
	testutil.CheckFieldLabelConversions(t, "v1", "BackingServiceInstance",
		// Ensure all currently returned labels are supported
		api.BackingServiceInstanceToSelectableFields(&backingserviceinstanceapi.BackingServiceInstance{}),
		// Ensure previously supported labels have conversions. DO NOT REMOVE THINGS FROM THIS LIST
		"spec.provisioning.backingservice_name",
	)
}