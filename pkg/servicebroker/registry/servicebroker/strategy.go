package servicebroker

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/fielderrors"

	"github.com/openshift/origin/pkg/servicebroker/api"
)

// sdnStrategy implements behavior for HostSubnets
type Strategy struct {
	runtime.ObjectTyper
}

// Strategy is the default logic that applies when creating and updating HostSubnet
// objects via the REST API.
var SbStrategy = Strategy{kapi.Scheme}

func (Strategy) PrepareForUpdate(obj, old runtime.Object) {}

// NamespaceScoped is false for sdns
func (Strategy) NamespaceScoped() bool {
	return false
}

func (Strategy) GenerateName(base string) string {
	return base
}

func (Strategy) PrepareForCreate(obj runtime.Object) {
}

// Validate validates a new sdn
func (Strategy) Validate(ctx kapi.Context, obj runtime.Object) fielderrors.ValidationErrorList {
	return fielderrors.ValidationErrorList{}
}

// AllowCreateOnUpdate is false for sdns
func (Strategy) AllowCreateOnUpdate() bool {
	return false
}

func (Strategy) AllowUnconditionalUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for a HostSubnet
func (Strategy) ValidateUpdate(ctx kapi.Context, obj, old runtime.Object) fielderrors.ValidationErrorList {
	return fielderrors.ValidationErrorList{}
}

// Matcher returns a generic matcher for a given label and field selector.
func Matcher(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{Label: label, Field: field, GetAttrs: getAttrs}
}

func getAttrs(obj runtime.Object) (objLabels labels.Set, objFields fields.Set, err error) {
	sb := obj.(*api.ServiceBroker)
	return labels.Set(sb.Labels), api.ServiceBrokerToSelectableFields(sb), nil
}
