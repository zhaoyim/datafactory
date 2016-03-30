package backingserviceinstance

import (
	"testing"
	//"time"

	kapi "k8s.io/kubernetes/pkg/api"
	//"k8s.io/kubernetes/pkg/api/unversioned"

	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
)

func TestBackingServiceInstanceStrategy(t *testing.T) {
	ctx := kapi.NewDefaultContext()
	if !BsiStrategy.NamespaceScoped() {
		t.Errorf("BackingServiceInstance is namespace scoped")
	}
	if BsiStrategy.AllowCreateOnUpdate() {
		t.Errorf("BackingServiceInstance should not allow create on update")
	}
	bsi := &backingserviceinstanceapi.BackingServiceInstance{
		ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
		Spec: backingserviceinstanceapi.BackingServiceInstanceSpec{
			InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
				BackingServiceName: "bs-1",
				BackingServicePlanGuid: "plan-abcde123",
			},
		},
	}
	BsiStrategy.PrepareForCreate(bsi)
	if len(bsi.Status.Phase) == 0 || bsi.Status.Phase != backingserviceinstanceapi.BackingServiceInstancePhaseProvisioning {
		t.Errorf("bsi phase is not provisioning")
	}
	errs := BsiStrategy.Validate(ctx, bsi)
	if len(errs) != 0 {
		t.Errorf("Unexpected error validating %v", errs)
	}

	bsi.ResourceVersion = "foo"
	errs = BsiStrategy.ValidateUpdate(ctx, bsi, bsi)
	if len(errs) != 0 {
		t.Errorf("Unexpected error validating %v", errs)
	}
	invalidBsi := &backingserviceinstanceapi.BackingServiceInstance{}
	errs = BsiStrategy.Validate(ctx, invalidBsi)
	if len(errs) == 0 {
		t.Errorf("Expected error validating")
	}
}