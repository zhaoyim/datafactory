package backingserviceinstance

import (
	"fmt"
	
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/validation/field"

	"github.com/openshift/origin/pkg/backingserviceinstance/api"
	"github.com/openshift/origin/pkg/backingserviceinstance/api/validation"
)

// sdnStrategy implements behavior for HostSubnets
type Strategy struct {
	runtime.ObjectTyper
	kapi.NameGenerator
}

// Strategy is the default logic that applies when creating and updating HostSubnet
// objects via the REST API.
var BsiStrategy = Strategy{kapi.Scheme, kapi.SimpleNameGenerator}

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

func (Strategy) Validate(ctx kapi.Context, obj runtime.Object) field.ErrorList {
	bsi:=obj.(*api.BackingServiceInstance)
	return validation.ValidateBackingServiceInstance(bsi)
}

// AllowCreateOnUpdate is false for sdns
func (Strategy) AllowCreateOnUpdate() bool {
	return false
}

func (Strategy) AllowUnconditionalUpdate() bool {
	return false
}

// CheckGracefulDelete allows a backingserviceinstance to be gracefully deleted.
func (Strategy) CheckGracefulDelete(obj runtime.Object, options *kapi.DeleteOptions) bool {
	return false
}

// ValidateUpdate is the default update validation for a HostSubnet
func (Strategy) ValidateUpdate(ctx kapi.Context, obj, old runtime.Object) field.ErrorList {
	 oldbsi := old.(*api.BackingServiceInstance)
	 bsi := obj.(*api.BackingServiceInstance)
	 return validation.ValidateBackingServiceInstanceUpdate(bsi, oldbsi)
}

// Matcher returns a generic matcher for a given label and field selector.
func Matcher(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{Label: label, Field: field, GetAttrs: getAttrs}
}

func getAttrs(obj runtime.Object) (objLabels labels.Set, objFields fields.Set, err error) {
	bsi, ok := obj.(*api.BackingServiceInstance)
	if !ok {
		return nil, nil, fmt.Errorf("not a BackingServiceInstance")
	}
	return labels.Set(bsi.Labels), api.BackingServiceInstanceToSelectableFields(bsi), nil
}
