package etcd

import (
	kapi "k8s.io/kubernetes/pkg/api"
	// "k8s.io/kubernetes/pkg/api/rest"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/registry/generic/registry"
	"k8s.io/kubernetes/pkg/runtime"

	"github.com/openshift/origin/pkg/application/api"
	"github.com/openshift/origin/pkg/application/registry/application"
	"github.com/openshift/origin/pkg/util/restoptions"
)

type REST struct {
	*registry.Store
}

// NewStorage returns a RESTStorage object that will work against nodes.
func NewREST(optsGetter restoptions.Getter) (*REST, error) {
	prefix := "/applications"

	store := &registry.Store{
		NewFunc:     func() runtime.Object { return &api.Application{} },
		NewListFunc: func() runtime.Object { return &api.ApplicationList{} },

		KeyRootFunc: func(ctx kapi.Context) string {
			return registry.NamespaceKeyRootFunc(ctx, prefix)
		},
		KeyFunc: func(ctx kapi.Context, id string) (string, error) {
			return registry.NamespaceKeyFunc(ctx, prefix, id)
		},

		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*api.Application).Name, nil
		},
		PredicateFunc: func(label labels.Selector, field fields.Selector) generic.Matcher {
			return application.Matcher(label, field)
		},

		QualifiedResource: api.Resource("applications"),

		CreateStrategy:      application.AppStrategy,
		UpdateStrategy:      application.AppStrategy,
		DeleteStrategy:      application.AppStrategy,
		ReturnDeletedObject: false,
	}

	if err := restoptions.ApplyOptions(optsGetter, store, prefix); err != nil {
		return nil, err
	}
	return &REST{store}, nil
}

// /// New returns a new object
// func (r *REST) New() runtime.Object {
// 	return r.store.NewFunc()
// }

// // NewList returns a new list object
// func (r *REST) NewList() runtime.Object {
// 	return r.store.NewListFunc()
// }

// // Get gets a specific image specified by its ID.
// func (r *REST) Get(ctx kapi.Context, name string) (runtime.Object, error) {
// 	return r.store.Get(ctx, name)
// }

// func (r *REST) List(ctx kapi.Context, options *kapi.ListOptions) (runtime.Object, error) {
// 	return r.store.List(ctx, options)
// }

// // Create creates an image based on a specification.
// func (r *REST) Create(ctx kapi.Context, obj runtime.Object) (runtime.Object, error) {
// 	app, ok := obj.(*api.Application)
// 	if ok {
// 		app.Status.Phase = api.ApplicationNew
// 	}
// 	return r.store.Create(ctx, obj)
// }

// func (r *REST) Watch(ctx kapi.Context, options *kapi.ListOptions) (watch.Interface, error) {
// 	return r.store.Watch(ctx, options)
// }

// // Update alters an existing image.
// func (r *REST) Update(ctx kapi.Context, obj runtime.Object) (runtime.Object, bool, error) {
// 	newApp, ok := obj.(*api.Application)
// 	if ok {
// 		switch newApp.Status.Phase {
// 		case api.ApplicationChecking:
// 			newApp.Status.Phase = api.ApplicationActive
// 			return r.store.Update(ctx, obj)
// 		case api.ApplicationTerminating:
// 			return r.store.Update(ctx, obj)
// 		}

// 		if newApp.Spec.Destory == true {
// 			newApp.Status.Phase = api.ApplicationTerminating
// 			return r.store.Update(ctx, obj)
// 		}

// 		oldObj, _ := r.store.Get(ctx, newApp.Name)
// 		if oldApp, ok := oldObj.(*api.Application); ok {
// 			switch oldApp.Status.Phase {
// 			case api.ApplicationActiveUpdate:
// 				newApp.Status.Phase = api.ApplicationActive
// 			case api.ApplicationActive:
// 				newApp.Status.Phase = api.ApplicationActiveUpdate
// 			}
// 		}
// 	}

// 	return r.store.Update(ctx, obj)
// }

// // Delete deletes an existing image specified by its ID.
// func (r *REST) Delete(ctx kapi.Context, name string, options *kapi.DeleteOptions) (runtime.Object, error) {
// 	appObj, err := r.Get(ctx, name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	application := appObj.(*api.Application)

// 	if application.Status.Phase == api.ApplicationTerminating {
// 		return r.store.Delete(ctx, name, options)
// 	}
// 	if application.Status.Phase == api.ApplicationTerminatingLabel {
// 		return r.store.Delete(ctx, name, options)
// 	}

// 	if application.DeletionTimestamp.IsZero() {
// 		now := unversioned.Now()
// 		application.DeletionTimestamp = &now
// 		application.Status.Phase = api.ApplicationTerminatingLabel
// 		result, _, err := r.store.Update(ctx, application)
// 		return result, err
// 	}

// 	return r.store.Delete(ctx, name, options)
// }

// func (r *REST) Watch(ctx kapi.Context, options *kapi.ListOptions) (watch.Interface, error) {
// 	return r.store.Watch(ctx, options)
// }
