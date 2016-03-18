package etcd

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	etcdgeneric "k8s.io/kubernetes/pkg/registry/generic/etcd"
	"k8s.io/kubernetes/pkg/storage"
	"k8s.io/kubernetes/pkg/watch"

	"github.com/openshift/origin/pkg/application/api"
	application "github.com/openshift/origin/pkg/application/registry/application"
	"k8s.io/kubernetes/pkg/runtime"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	oclient "github.com/openshift/origin/pkg/client"
)

type REST struct {
	store *etcdgeneric.Etcd
}

// NewREST returns a new REST.
func NewREST(s storage.Interface, oClient *oclient.Client, kClient *kclient.Client) *REST {
	prefix := "/applications"
	application.AppStrategy.OClient = oClient
	application.AppStrategy.KClient = kClient
	store := &etcdgeneric.Etcd{
		NewFunc:     func() runtime.Object { return &api.Application{} },
		NewListFunc: func() runtime.Object { return &api.ApplicationList{} },
		KeyRootFunc: func(ctx kapi.Context) string {
			return prefix
		},
		KeyFunc: func(ctx kapi.Context, name string) (string, error) {
			return etcdgeneric.NoNamespaceKeyFunc(ctx, prefix, name)
		},
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*api.Application).Name, nil
		},
		PredicateFunc: func(label labels.Selector, field fields.Selector) generic.Matcher {
			return application.Matcher(label, field)
		},
		EndpointName: "application",

		CreateStrategy:      application.AppStrategy,
		UpdateStrategy:      application.AppStrategy,
		ReturnDeletedObject: false,

		Storage: s,
	}
	return &REST{store: store}
}

/// New returns a new object
func (r *REST) New() runtime.Object {
	return r.store.NewFunc()
}

// NewList returns a new list object
func (r *REST) NewList() runtime.Object {
	return r.store.NewListFunc()
}

// Get gets a specific image specified by its ID.
func (r *REST) Get(ctx kapi.Context, name string) (runtime.Object, error) {
	return r.store.Get(ctx, name)
}

func (r *REST) List(ctx kapi.Context, label labels.Selector, field fields.Selector) (runtime.Object, error) {
	return r.store.List(ctx, label, field)
}

// Create creates an image based on a specification.
func (r *REST) Create(ctx kapi.Context, obj runtime.Object) (runtime.Object, error) {
	app, ok := obj.(*api.Application)
	if ok {
		app.Status.Phase = api.ApplicationNew
	}
	return r.store.Create(ctx, obj)
}

// Update alters an existing image.
func (r *REST) Update(ctx kapi.Context, obj runtime.Object) (runtime.Object, bool, error) {
	newApp, ok := obj.(*api.Application)
	if ok {
		switch newApp.Status.Phase {
		case api.ApplicationChecking:
			newApp.Status.Phase = api.ApplicationActive
			return r.store.Update(ctx, obj)
		case api.ApplicationTerminating:
			return r.store.Update(ctx, obj)
		}

		if newApp.Spec.Destory == true {
			newApp.Status.Phase = api.ApplicationTerminating
			return r.store.Update(ctx, obj)
		}

		oldObj, _ := r.store.Get(ctx, newApp.Name)
		if oldApp, ok := oldObj.(*api.Application); ok {
			switch oldApp.Status.Phase {
			case api.ApplicationActiveUpdate:
				newApp.Status.Phase = api.ApplicationActive
			case api.ApplicationActive:
				newApp.Status.Phase = api.ApplicationActiveUpdate
			}
		}
	}

	return r.store.Update(ctx, obj)
}

// Delete deletes an existing image specified by its ID.
func (r *REST) Delete(ctx kapi.Context, name string, options *kapi.DeleteOptions) (runtime.Object, error) {
	appObj, err := r.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	application := appObj.(*api.Application)

	if application.Status.Phase == api.ApplicationTerminating {
		return r.store.Delete(ctx, name, options)
	}
	if application.Status.Phase == api.ApplicationTerminatingLabel {
		return r.store.Delete(ctx, name, options)
	}

	if application.DeletionTimestamp.IsZero() {
		now := unversioned.Now()
		application.DeletionTimestamp = &now
		application.Status.Phase = api.ApplicationTerminatingLabel
		result, _, err := r.store.Update(ctx, application)
		return result, err
	}

	return r.store.Delete(ctx, name, options)
}

func (r *REST) Watch(ctx kapi.Context, label labels.Selector, field fields.Selector, resourceVersion string) (watch.Interface, error) {
	return r.store.Watch(ctx, label, field, resourceVersion)
}
