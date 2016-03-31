package servicebroker

import (
	"fmt"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/validation/field"

	"github.com/openshift/origin/pkg/servicebroker/api"
	"github.com/openshift/origin/pkg/servicebroker/api/validation"
)

// sdnStrategy implements behavior for HostSubnets
type Strategy struct {
	runtime.ObjectTyper
}

// Strategy is the default logic that applies when creating and updating HostSubnet
// objects via the REST API.
var SbStrategy = Strategy{kapi.Scheme}

func (Strategy) Canonicalize(obj runtime.Object) {}
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
func (Strategy) Validate(ctx kapi.Context, obj runtime.Object) field.ErrorList {
	return validation.ValidateServiceBroker(obj.(*api.ServiceBroker))
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
	return validation.ValidateServiceBrokerUpdate(obj.(*api.ServiceBroker),old.(*api.ServiceBroker))
}

// Matcher returns a generic matcher for a given label and field selector.
// Matcher returns a generic matcher for a given label and field selector.

func Matcher(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{
		Label: label,
		Field: field,
		GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			sb, ok := obj.(*api.ServiceBroker)
			if !ok {
				return nil, nil, fmt.Errorf("not a servicebroker")
			}
			return labels.Set(sb.ObjectMeta.Labels), api.ServiceBrokerToSelectableFields(sb), nil
		},
	}
}
/*
func Matcher(label labels.Selector, field fields.Selector) generic.Matcher {
	return generic.MatcherFunc(func(obj runtime.Object) (bool, error) {
		sb, ok := obj.(*api.ServiceBroker)
		if !ok {
			return false, fmt.Errorf("not a ServiceBroker")
		}
		return label.Matches(labels.Set(sb.Labels)) && field.Matches(api.ServiceBrokerToSelectableFields(sb)), nil
	})
}
*/


func getAttrs(obj runtime.Object) (objLabels labels.Set, objFields fields.Set, err error) {
	sb := obj.(*api.ServiceBroker)
	return labels.Set(sb.Labels), api.ServiceBrokerToSelectableFields(sb), nil
}
