package etcd

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/storage"
	"k8s.io/kubernetes/pkg/watch"
	kerrors "k8s.io/kubernetes/pkg/api/errors"
	etcdgeneric "k8s.io/kubernetes/pkg/registry/generic/etcd"


	servicebrokerapi "github.com/openshift/origin/pkg/servicebroker/api"
	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
	"github.com/openshift/origin/pkg/servicebroker/registry/servicebroker"
	oclient "github.com/openshift/origin/pkg/client"
	"errors"
	"fmt"
)

type REST struct {
	store     *etcdgeneric.Etcd
	bsClient  oclient.BackingServiceInterface
	bsiClient oclient.BackingServiceInstanceInterface
}

// NewREST returns a new REST.
func NewREST(s storage.Interface, oClient *oclient.Client) *REST {
	prefix := "/servicebrokers"
	store := &etcdgeneric.Etcd{
		NewFunc:     func() runtime.Object {
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
	return &REST{
		store: store,
		bsClient: oClient.BackingServices(backingserviceapi.BSNS),
		bsiClient: oClient.BackingServiceInstances(kapi.NamespaceAll),
	}
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
	if num, _ := countWorkingBackingServiceInstance(name, r.bsClient, r.bsiClient); num > 0 {
		return nil, errors.New(fmt.Sprintf("can not remove servicebroker %s, cause %d backingserviceinstances are using it", name, num))
	}

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
	return r.store.Watch(ctx,options)

}

func listBackingServiceByServiceBrokerName(name string, bsClient  oclient.BackingServiceInterface) (*backingserviceapi.BackingServiceList, error) {
	selector, _ := labels.Parse(servicebrokerapi.ServiceBrokerLabel + "=" + name)
	return  bsClient.List(kapi.ListOptions{LabelSelector : selector})
}

func listBackingServiceInstanceByBackingServiceName(name string, bsiClient oclient.BackingServiceInstanceInterface) (*backingserviceinstanceapi.BackingServiceInstanceList, error) {
	selector, _ := fields.ParseSelector("spec.provisioning.backingservice_name=" + name)
	return bsiClient.List(kapi.ListOptions{FieldSelector: selector})
}

func countWorkingBackingServiceInstance(name string, bsClient  oclient.BackingServiceInterface, bsiClient oclient.BackingServiceInstanceInterface) (int, error) {
	total := 0

	bsList , err := listBackingServiceByServiceBrokerName(name, bsClient)
	if err != nil {
		return total, err
	}

	for _, v := range bsList.Items {
		bsiList, err := listBackingServiceInstanceByBackingServiceName(v.Name, bsiClient)
		if err != nil && kerrors.IsNotFound(err) {
			continue
		}

		total += len(bsiList.Items)
	}

	return total, nil
}