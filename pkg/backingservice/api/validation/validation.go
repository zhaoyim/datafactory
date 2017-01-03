package validation

import (
	oapi "github.com/openshift/origin/pkg/api"
	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
	"k8s.io/kubernetes/pkg/api/validation"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

func BackingServicetName(name string, prefix bool) []string {
	if reasons := oapi.MinimalNameRequirements(name, prefix); len(reasons) != 0 {
		return reasons
	}

	if len(name) < 2 {
		return []string{"must be at least 2 characters long"}
	}

	return nil
}

// ValidateBackingService tests required fields for a BackingService.
func ValidateBackingService(bs *backingserviceapi.BackingService) field.ErrorList {

	allErrs := validation.ValidateObjectMeta(&bs.ObjectMeta, true, BackingServicetName, field.NewPath("metadata"))

	return allErrs
}

// ValidateBuildRequest validates a BuildRequest object
func ValidateBackingServiceUpdate(bs *backingserviceapi.BackingService, older *backingserviceapi.BackingService) field.ErrorList {

	allErrs := validation.ValidateObjectMetaUpdate(&bs.ObjectMeta, &older.ObjectMeta, field.NewPath("metadata"))

	allErrs = append(allErrs, ValidateBackingService(bs)...)

	return allErrs
}
