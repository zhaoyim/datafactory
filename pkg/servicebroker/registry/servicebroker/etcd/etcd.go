package etcd

import (
	servicebrokerapi "github.com/openshift/origin/pkg/servicebroker/api"
	"github.com/openshift/origin/pkg/servicebroker/registry/servicebroker"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	etcdgeneric "k8s.io/kubernetes/pkg/registry/generic/etcd"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/storage"
	"k8s.io/kubernetes/pkg/watch"
)

type REST struct {
	store *etcdgeneric.Etcd
}

// NewREST returns a new REST.
func NewREST(s storage.Interface) *REST {
	prefix := "/servicebrokers"
	store := &etcdgeneric.Etcd{
		NewFunc: func() runtime.Object {
			return &servicebrokerapi.ServiceBroker{}
		},
		NewListFunc: func() runtime.Object {
			return &servicebrokerapi.ServiceBrokerList{}
		},
		KeyRootFunc: func(ctx kapi.Context) string {
			return prefix
		},
		KeyFunc: func(ctx kapi.Context, name string) (string, error) {
			return etcdgeneric.NoNamespaceKeyFunc(ctx, prefix, name)
		},
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*servicebrokerapi.ServiceBroker).Name, nil
		},
		PredicateFunc: func(label labels.Selector, field fields.Selector) generic.Matcher {
			return servicebroker.Matcher(label, field)
		},

		QualifiedResource: servicebrokerapi.Resource("servicebroker"),

		CreateStrategy: servicebroker.SbStrategy,
		UpdateStrategy: servicebroker.SbStrategy,

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

	servicebroker := obj.(*servicebrokerapi.ServiceBroker)
	servicebroker.Status.Phase = servicebrokerapi.ServiceBrokerNew

	return r.store.Create(ctx, obj)
}

// Update alters an existing image.
func (r *REST) Update(ctx kapi.Context, obj runtime.Object) (runtime.Object, bool, error) {
	return r.store.Update(ctx, obj)
}

// Delete deletes an existing image specified by its ID.
func (r *REST) Delete(ctx kapi.Context, name string, options *kapi.DeleteOptions) (runtime.Object, error) {

	sbObj, err := r.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	servicebroker := sbObj.(*servicebrokerapi.ServiceBroker)

	if servicebroker.DeletionTimestamp.IsZero() {
		now := unversioned.Now()
		servicebroker.DeletionTimestamp = &now
		servicebroker.Status.Phase = servicebrokerapi.ServiceBrokerDeleting
		result, _, err := r.store.Update(ctx, servicebroker)
		return result, err
	}

	return r.store.Delete(ctx, name, options)
}

func (r *REST) Watch(ctx kapi.Context, options *kapi.ListOptions) (watch.Interface, error) {
	return r.store.Watch(ctx, options)

}
