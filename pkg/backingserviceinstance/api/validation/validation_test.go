package validation

import (
	//"strings"
	"testing"

	kapi "k8s.io/kubernetes/pkg/api"
	//"k8s.io/kubernetes/pkg/util/fielderrors"

	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
)



func TestBuildValidationSuccess(t *testing.T) {
	bsi := &backingserviceinstanceapi.BackingServiceInstance {
		ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
		Spec: backingserviceinstanceapi.BackingServiceInstanceSpec {
			InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
				BackingServiceName:     "bsi-1",
				BackingServicePlanGuid: "plan-abc123",
			},
		},
		Status: backingserviceinstanceapi.BackingServiceInstanceStatus {
			Phase: backingserviceinstanceapi.BackingServiceInstancePhaseProvisioning,
		},
	}
	if result := ValidateBackingServiceInstance(bsi); len(result) > 0 {
		t.Errorf("Unexpected validation error returned %v", result)
	}
}

func TestBuildValidationFailure(t *testing.T) {
	build := &backingserviceinstanceapi.BackingServiceInstance {
	}
	if result := ValidateBackingServiceInstance(build); len(result) == 0 {
		t.Errorf("Unexpected validation error returned %v", result)
	}
}

func TestValidateBuildUpdate(t *testing.T) {
	errs := ValidateBackingServiceInstanceUpdate(
		&backingserviceinstanceapi.BackingServiceInstance {
			ObjectMeta: kapi.ObjectMeta{Namespace: kapi.NamespaceDefault, Name: "my-build", ResourceVersion: "1"},
			Spec:        backingserviceinstanceapi.BackingServiceInstanceSpec {
				InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
					BackingServiceName:     "bsi-1",
					BackingServicePlanGuid: "plan-abc123",
				},
			},
			Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
				Phase: backingserviceinstanceapi.BackingServiceInstancePhaseBound,
				Action: "",
			},
		},
		&backingserviceinstanceapi.BackingServiceInstance {
			ObjectMeta: kapi.ObjectMeta{Namespace: kapi.NamespaceDefault, Name: "my-build", ResourceVersion: "1"},
			Spec:        backingserviceinstanceapi.BackingServiceInstanceSpec {
				InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
					BackingServiceName:     "bsi-1",
					BackingServicePlanGuid: "plan-abc123",
				},
			},
			Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
				Phase: backingserviceinstanceapi.BackingServiceInstancePhaseBound,
				Action: backingserviceinstanceapi.BackingServiceInstanceActionToBind,
			},
		},
	)
	if len(errs) != 0 {
		t.Errorf("expected success: %v", errs)
	}
	
	// ...

	errorCases := map[string]struct {
		Old    *backingserviceinstanceapi.BackingServiceInstance
		Update *backingserviceinstanceapi.BackingServiceInstance
	}{
		"changed Spec.Status.Phase": {
			Old: &backingserviceinstanceapi.BackingServiceInstance {
				ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
				Spec:       backingserviceinstanceapi.BackingServiceInstanceSpec {
					InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
						BackingServiceName:     "bsi-1",
						BackingServicePlanGuid: "plan-abc123",
					},
				},
				Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
					Phase: backingserviceinstanceapi.BackingServiceInstancePhaseBound,
				},
			},
			Update: &backingserviceinstanceapi.BackingServiceInstance {
				ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
				Spec:       backingserviceinstanceapi.BackingServiceInstanceSpec {
					InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
						BackingServiceName:     "bsi-1",
						BackingServicePlanGuid: "plan-abc123",
					},
				},
				Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
					Phase: backingserviceinstanceapi.BackingServiceInstancePhaseUnbound,
				},
			},
		},
		"changed Spec.BackingServiceName": {
			Old: &backingserviceinstanceapi.BackingServiceInstance {
				ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
				Spec:       backingserviceinstanceapi.BackingServiceInstanceSpec {
					InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
						BackingServiceName:     "bsi-1",
						BackingServicePlanGuid: "plan-abc123",
					},
				},
				Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
					Phase: backingserviceinstanceapi.BackingServiceInstancePhaseBound,
				},
			},
			Update: &backingserviceinstanceapi.BackingServiceInstance {
				ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
				Spec:       backingserviceinstanceapi.BackingServiceInstanceSpec {
					InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
						BackingServiceName:     "bsi-2",
						BackingServicePlanGuid: "plan-abc123",
					},
				},
				Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
					Phase: backingserviceinstanceapi.BackingServiceInstancePhaseBound,
				},
			},
		},
		"changed Spec.BackingServicePlanGuid": {
			Old: &backingserviceinstanceapi.BackingServiceInstance {
				ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
				Spec:       backingserviceinstanceapi.BackingServiceInstanceSpec {
					InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
						BackingServiceName:     "bsi-1",
						BackingServicePlanGuid: "plan-abc123",
					},
				},
				Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
					Phase: backingserviceinstanceapi.BackingServiceInstancePhaseBound,
				},
			},
			Update: &backingserviceinstanceapi.BackingServiceInstance {
				ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
				Spec:       backingserviceinstanceapi.BackingServiceInstanceSpec {
					InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
						BackingServiceName:     "bsi-1",
						BackingServicePlanGuid: "plan-xxxxx",
					},
				},
				Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
					Phase: backingserviceinstanceapi.BackingServiceInstancePhaseBound,
				},
			},
		},
		"changed Spec.InstanceID": {
			Old: &backingserviceinstanceapi.BackingServiceInstance {
				ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
				Spec:       backingserviceinstanceapi.BackingServiceInstanceSpec {
					InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
						BackingServiceName:     "bsi-1",
						BackingServicePlanGuid: "plan-abc123",
					},
					InstanceID: "a1b2c3d4",
				},
				Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
					Phase: backingserviceinstanceapi.BackingServiceInstancePhaseBound,
				},
			},
			Update: &backingserviceinstanceapi.BackingServiceInstance {
				ObjectMeta: kapi.ObjectMeta{Name: "bsi-1", Namespace: "default"},
				Spec:       backingserviceinstanceapi.BackingServiceInstanceSpec {
					InstanceProvisioning: backingserviceinstanceapi.InstanceProvisioning {
						BackingServiceName:     "bsi-1",
						BackingServicePlanGuid: "plan-abc123",
					},
					InstanceID: "xxxxxx",
				},
				Status:     backingserviceinstanceapi.BackingServiceInstanceStatus {
					Phase: backingserviceinstanceapi.BackingServiceInstancePhaseBound,
				},
			},
		},
	}

	for k, v := range errorCases {
		errs := ValidateBackingServiceInstanceUpdate(v.Update, v.Old)
		if len(errs) == 0 {
			t.Errorf("expected failure %s for %v", k, v.Update)
			continue
		}
	}
}