package application

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/validation/field"

	"github.com/openshift/origin/pkg/application/api"
	"github.com/openshift/origin/pkg/application/api/validation"

	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	oclient "github.com/openshift/origin/pkg/client"

)

type Strategy struct {
	runtime.ObjectTyper
	KClient *kclient.Client
	OClient *oclient.Client
}

// Strategy is the default logic that applies when creating and updating HostSubnet
// objects via the REST API.
var AppStrategy = Strategy{kapi.Scheme, nil, nil}
func (s Strategy) Canonicalize(obj runtime.Object) {}
func (s Strategy) PrepareForUpdate(obj, old runtime.Object) {}

func (s Strategy) NamespaceScoped() bool {
	return true
}

func (s Strategy) GenerateName(base string) string {
	return base
}

func (s Strategy) PrepareForCreate(obj runtime.Object) {
}

func (s Strategy) Validate(ctx kapi.Context, obj runtime.Object) field.ErrorList {
	application := obj.(*api.Application)

	return validation.ValidateApplication(application, s.OClient, s.KClient)
}

func (s Strategy) AllowCreateOnUpdate() bool {
	return false
}

func (s Strategy) AllowUnconditionalUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for a HostSubnet
func (s Strategy) ValidateUpdate(ctx kapi.Context, obj, old runtime.Object) field.ErrorList {
	return validation.ValidateApplicationUpdate(obj.(*api.Application),old.(*api.Application))
}

// Matcher returns a generic matcher for a given label and field selector.
func Matcher(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{Label: label, Field: field, GetAttrs: getAttrs}
}

func getAttrs(obj runtime.Object) (objLabels labels.Set, objFields fields.Set, err error) {
	app := obj.(*api.Application)
	return labels.Set(app.Labels), api.ApplicationToSelectableFields(app), nil
}
