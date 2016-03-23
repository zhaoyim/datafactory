package etcd

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	etcdgeneric "k8s.io/kubernetes/pkg/registry/generic/etcd"
	"k8s.io/kubernetes/pkg/storage"
	"k8s.io/kubernetes/pkg/watch"

	"github.com/openshift/origin/pkg/backingservice/api"
	"github.com/openshift/origin/pkg/backingservice/registry/backingservice"
	"k8s.io/kubernetes/pkg/runtime"
)

type REST struct {
	store *etcdgeneric.Etcd
}

// NewREST returns a new REST.
func NewREST(s storage.Interface) *REST {
	prefix := "/backingservices"
	store := &etcdgeneric.Etcd{
		NewFunc:     func() runtime.Object { return &api.BackingService{} },
		NewListFunc: func() runtime.Object { return &api.BackingServiceList{} },
		//KeyRootFunc: func(ctx kapi.Context) string {
		//	return prefix
		//},
		//KeyFunc: func(ctx kapi.Context, name string) (string, error) {
		//	return etcdgeneric.NoNamespaceKeyFunc(ctx, prefix, name)
		//},
		KeyRootFunc: func(ctx kapi.Context) string {
			return etcdgeneric.NamespaceKeyRootFunc(ctx, prefix)
		},
		KeyFunc: func(ctx kapi.Context, name string) (string, error) {
			return etcdgeneric.NamespaceKeyFunc(ctx, prefix, name)
		},
		
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*api.BackingService).Name, nil
		},
		PredicateFunc: func(label labels.Selector, field fields.Selector) generic.Matcher {
			return backingservice.Matcher(label, field)
		},

		QualifiedResource: api.Resource("backingservice"),

		CreateStrategy: backingservice.BsStrategy,
		UpdateStrategy: backingservice.BsStrategy,

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

func (r *REST) List(ctx kapi.Context, options *kapi.ListOptions) (runtime.Object, error) {
	return r.store.List(ctx, options)
}

// Create creates an image based on a specification.
func (r *REST) Create(ctx kapi.Context, obj runtime.Object) (runtime.Object, error) {
	if bs, ok := obj.(*api.BackingService); ok {
		bs.Status.Phase = api.BackingServicePhaseActive
	}
	return r.store.Create(ctx, obj)
}

// Update alters an existing image.
func (r *REST) Update(ctx kapi.Context, obj runtime.Object) (runtime.Object, bool, error) {
	return r.store.Update(ctx, obj)
}

// Delete deletes an existing image specified by its ID.
func (r *REST) Delete(ctx kapi.Context, name string, options *kapi.DeleteOptions) (runtime.Object, error) {
	return r.store.Delete(ctx, name, options)
}

func (r *REST) Watch(ctx kapi.Context,options *kapi.ListOptions) (watch.Interface, error) {
	return r.store.Watch(ctx, options)
}
