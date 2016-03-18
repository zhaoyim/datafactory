package validation

import (
	oapi "github.com/openshift/origin/pkg/api"
	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
	"k8s.io/kubernetes/pkg/api/validation"
	"k8s.io/kubernetes/pkg/util/fielderrors"
)

func BackingServicetName(name string, prefix bool) (bool, string) {
	if ok, reason := oapi.MinimalNameRequirements(name, prefix); !ok {
		return ok, reason
	}

	if len(name) < 2 {
		return false, "must be at least 2 characters long"
	}

	return true, ""
}

// ValidateBackingService tests required fields for a BackingService.
func ValidateBackingService(bs *backingserviceapi.BackingService) fielderrors.ValidationErrorList {
	allErrs := fielderrors.ValidationErrorList{}
	allErrs = append(allErrs, validation.ValidateObjectMeta(&bs.ObjectMeta, true, BackingServicetName).Prefix("metadata")...)
	//allErrs = append(allErrs, validateBuildSpec(&build.Spec).Prefix("spec")...)
	return allErrs
}

// ValidateBuildRequest validates a BuildRequest object
func ValidateBackingServiceUpdate(bs *backingserviceapi.BackingService, older *backingserviceapi.BackingService) fielderrors.ValidationErrorList {
	allErrs := fielderrors.ValidationErrorList{}
	allErrs = append(allErrs, validation.ValidateObjectMetaUpdate(&bs.ObjectMeta, &older.ObjectMeta).Prefix("metadata")...)

	allErrs = append(allErrs, ValidateBackingService(bs)...)

	if older.Status.Phase != bs.Status.Phase {
		allErrs = append(allErrs, fielderrors.NewFieldInvalid("status.Phase", bs.Status.Phase, "phase cannot be updated from a terminal state"))
	}
	return allErrs
}
