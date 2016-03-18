package validation

import (
	"reflect"

	"k8s.io/kubernetes/pkg/api/validation"
	"k8s.io/kubernetes/pkg/util/fielderrors"

	oapi "github.com/openshift/origin/pkg/api"
	servicebrokerapi "github.com/openshift/origin/pkg/servicebroker/api"
)

func ValidateServiceBrokerName(name string, prefix bool) (bool, string) {
	if ok, reason := oapi.MinimalNameRequirements(name, prefix); !ok {
		return ok, reason
	}

	if len(name) < 2 {
		return false, "must be at least 2 characters long"
	}

	return true, ""
}

// ValidateServiceBroker tests required fields for a ServiceBroker.
// This should only be called when creating a servicebroker (not on update),
// since its name validation is more restrictive than default namespace name validation
func ValidateServiceBroker(servicebroker *servicebrokerapi.ServiceBroker) fielderrors.ValidationErrorList {
	result := fielderrors.ValidationErrorList{}
	result = append(result, validation.ValidateObjectMeta(&servicebroker.ObjectMeta, false, ValidateServiceBrokerName).Prefix("metadata")...)

	return result
}

// ValidateServiceBrokerUpdate tests to make sure a servicebroker update can be applied.  Modifies newServiceBroker with immutable fields.
func ValidateServiceBrokerUpdate(newServiceBroker *servicebrokerapi.ServiceBroker, oldServiceBroker *servicebrokerapi.ServiceBroker) fielderrors.ValidationErrorList {
	allErrs := fielderrors.ValidationErrorList{}
	allErrs = append(allErrs, validation.ValidateObjectMetaUpdate(&newServiceBroker.ObjectMeta, &oldServiceBroker.ObjectMeta).Prefix("metadata")...)
	allErrs = append(allErrs, ValidateServiceBroker(newServiceBroker)...)

	if !reflect.DeepEqual(newServiceBroker.Spec.Finalizers, oldServiceBroker.Spec.Finalizers) {
		allErrs = append(allErrs, fielderrors.NewFieldInvalid("spec.finalizers", oldServiceBroker.Spec.Finalizers, "field is immutable"))
	}
	if !reflect.DeepEqual(newServiceBroker.Status, oldServiceBroker.Status) {
		allErrs = append(allErrs, fielderrors.NewFieldInvalid("status", oldServiceBroker.Spec.Finalizers, "field is immutable"))
	}

	for name, value := range newServiceBroker.Labels {
		if value != oldServiceBroker.Labels[name] {
			allErrs = append(allErrs, fielderrors.NewFieldInvalid("metadata.labels["+name+"]", value, "field is immutable, , try updating the namespace"))
		}
	}
	for name, value := range oldServiceBroker.Labels {
		if _, inNew := newServiceBroker.Labels[name]; !inNew {
			allErrs = append(allErrs, fielderrors.NewFieldInvalid("metadata.labels["+name+"]", value, "field is immutable, try updating the namespace"))
		}
	}

	return allErrs
}
