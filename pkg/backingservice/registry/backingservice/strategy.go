package backingservice

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/validation/field"
	"github.com/openshift/origin/pkg/backingservice/api"
	"github.com/openshift/origin/pkg/backingservice/api/validation"
)

// sdnStrategy implements behavior for HostSubnets
type Strategy struct {
	runtime.ObjectTyper
}

// Strategy is the default logic that applies when creating and updating HostSubnet
// objects via the REST API.
var BsStrategy = Strategy{kapi.Scheme}
func (Strategy) Canonicalize(obj runtime.Object) {}

func (Strategy) PrepareForUpdate(obj, old runtime.Object) {}

// NamespaceScoped is false for sdns
func (Strategy) NamespaceScoped() bool {
	return true
}

func (Strategy) GenerateName(base string) string {
	return base
}

func (Strategy) PrepareForCreate(obj runtime.Object) {
}

// Validate validates a new sdn
func (Strategy) Validate(ctx kapi.Context, obj runtime.Object) field.ErrorList {
	bs:=obj.(*api.BackingService)
	return validation.ValidateBackingService(bs)
}

// AllowCreateOnUpdate is false for sdns
func (Strategy) AllowCreateOnUpdate() bool {
	return false
}

func (Strategy) AllowUnconditionalUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for a HostSubnet
func (Strategy) ValidateUpdate(ctx kapi.Context, obj, old runtime.Object) field.ErrorList {

	return validation.ValidateBackingServiceUpdate(obj.(*api.BackingService), old.(*api.BackingService))
}

// Matcher returns a generic matcher for a given label and field selector.
func Matcher(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{Label: label, Field: field, GetAttrs: getAttrs}
}

func getAttrs(obj runtime.Object) (objLabels labels.Set, objFields fields.Set, err error) {
	bs := obj.(*api.BackingService)
	return labels.Set(bs.Labels), api.BackingServiceToSelectableFields(bs), nil
}
