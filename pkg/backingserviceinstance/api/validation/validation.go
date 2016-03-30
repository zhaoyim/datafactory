package validation

import (
	kapi "k8s.io/kubernetes/pkg/api"
	oapi "github.com/openshift/origin/pkg/api"
	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
	"k8s.io/kubernetes/pkg/api/validation"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

func BackingServiceInstanceName(name string, prefix bool) (bool, string) {
	if ok, reason := oapi.MinimalNameRequirements(name, prefix); !ok {
		return ok, reason
	}

	if len(name) < 2 {
		return false, "must be at least 2 characters long"
	}

	return true, ""
}

func validateBackingServiceInstanceSpec(spec *backingserviceinstanceapi.BackingServiceInstanceSpec) field.ErrorList {
	allErrs := field.ErrorList{}
	
	specPath := field.NewPath("spec")
	if spec.BackingServiceName == "" {
		allErrs = append(allErrs, field.Invalid(specPath.Child("backingservice_name"), spec.BackingServiceName, "BackingServiceName must be specified"))
	}
	if spec.BackingServicePlanGuid == "" {
		allErrs = append(allErrs, field.Invalid(specPath.Child("backingservice_plan_guid"), spec.BackingServicePlanGuid, "BackingServicePlanGuid must be specified"))
	}
	
	return allErrs
}

// ValidateBackingServiceInstance tests required fields for a BackingServiceInstance.
func ValidateBackingServiceInstance(bsi *backingserviceinstanceapi.BackingServiceInstance) field.ErrorList {
	allErrs := validation.ValidateObjectMeta(&bsi.ObjectMeta, true, BackingServiceInstanceName, field.NewPath("metadata"))
	
	allErrs = append(allErrs, validateBackingServiceInstanceSpec(&bsi.Spec)...)
	return allErrs
}

// ValidateBuildRequest validates a BuildRequest object
func ValidateBackingServiceInstanceUpdate(bsi *backingserviceinstanceapi.BackingServiceInstance, older *backingserviceinstanceapi.BackingServiceInstance) field.ErrorList {
 	allErrs := validation.ValidateObjectMetaUpdate(&bsi.ObjectMeta, &older.ObjectMeta, field.NewPath("metadata"))
	
 	allErrs = append(allErrs, ValidateBackingServiceInstance(bsi)...)

	statusPath := field.NewPath("status")
	if older.Status.Phase != bsi.Status.Phase {
		allErrs = append(allErrs, field.Invalid(statusPath.Child("phase"), bsi.Status.Phase, "phase cannot be updated from external"))
	}
	
	specPath := field.NewPath("spec")
	if !kapi.Semantic.DeepEqual(bsi.Spec, older.Spec) {
		allErrs = append(allErrs, field.Invalid(specPath, "content of spec is not printed out, please refer to the \"details\"", "spec is immutable"))
	}

	return allErrs
}

//==========================================

//func ValidateBackingServiceInstanceBindingRequest(bi *BindingRequest) field.ErrorList {
//	allErrs := field.ErrorList{}
//	// todo
//	return allErrs
//}

func ValidateBackingServiceInstanceBindingRequestOptions(o *backingserviceinstanceapi.BindingRequestOptions) field.ErrorList {
	return validation.ValidateObjectMeta(&o.ObjectMeta, true, oapi.MinimalNameRequirements, field.NewPath("metadata"))
}

func ValidateBackingServiceInstanceBindingRequestOptionsUpdate(o *backingserviceinstanceapi.BindingRequestOptions, older *backingserviceinstanceapi.BindingRequestOptions) field.ErrorList {
	return ValidateBackingServiceInstanceBindingRequestOptions(o)
}



