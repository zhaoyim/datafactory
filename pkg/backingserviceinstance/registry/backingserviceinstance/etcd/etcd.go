package etcd

import (
	"errors"
	"fmt"

	//"k8s.io/kubernetes/pkg/api/rest"

	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	etcdgeneric "k8s.io/kubernetes/pkg/registry/generic/etcd"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/storage"
	"k8s.io/kubernetes/pkg/watch"
	//"k8s.io/kubernetes/pkg/util"

	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api/unversioned"

	//backingserviceregistry "github.com/openshift/origin/pkg/backingservice/registry"
	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
	backingserviceinstanceregistry "github.com/openshift/origin/pkg/backingserviceinstance/registry/backingserviceinstance"
	//backingserviceinstancecontroller "github.com/openshift/origin/pkg/backingserviceinstance/controller"
	deployconfigregistry "github.com/openshift/origin/pkg/deploy/registry/deployconfig"
)

const BackingServiceInstancePath = "/backingserviceinstances"

type BackingServiceInstanceStorage struct {
	Instance *REST
	Binding  *BindingREST
}

type REST struct {
	store *etcdgeneric.Etcd
}

func NewREST(s storage.Interface) *REST {
	store := &etcdgeneric.Etcd{
		NewFunc: func() runtime.Object {
			return &backingserviceinstanceapi.BackingServiceInstance{}
		},
		NewListFunc: func() runtime.Object {
			return &backingserviceinstanceapi.BackingServiceInstanceList{}
		},
		KeyRootFunc: func(ctx kapi.Context) string {
			return etcdgeneric.NamespaceKeyRootFunc(ctx, BackingServiceInstancePath)
		},
		KeyFunc: func(ctx kapi.Context, id string) (string, error) {
			return etcdgeneric.NamespaceKeyFunc(ctx, BackingServiceInstancePath, id)
		},
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*backingserviceinstanceapi.BackingServiceInstance).Name, nil
		},
		PredicateFunc: func(label labels.Selector, field fields.Selector) generic.Matcher {
			return backingserviceinstanceregistry.Matcher(label, field)
		},
		QualifiedResource: backingserviceinstanceapi.Resource("backingserviceinstance"),


		CreateStrategy: backingserviceinstanceregistry.BsiStrategy,
		UpdateStrategy: backingserviceinstanceregistry.BsiStrategy,

		ReturnDeletedObject: false,

		Storage: s,
	}

	return &REST{store: store}
}

func (r *REST) New() runtime.Object {
	return r.store.NewFunc()
}

func (r *REST) NewList() runtime.Object {
	return r.store.NewListFunc()
}

func (r *REST) Get(ctx kapi.Context, name string) (runtime.Object, error) {
	return r.store.Get(ctx, name)
}

func (r *REST) List(ctx kapi.Context, options *kapi.ListOptions) (runtime.Object, error) {
	return r.store.List(ctx,options)
}

func (r *REST) Create(ctx kapi.Context, obj runtime.Object) (runtime.Object, error) {
	return r.store.Create(ctx, obj)
}

func (r *REST) Update(ctx kapi.Context, obj runtime.Object) (runtime.Object, bool, error) {
	return r.store.Update(ctx, obj)
}

func (r *REST) Delete(ctx kapi.Context, name string, options *kapi.DeleteOptions) (runtime.Object, error) {
	bsiObj, err := r.Get(ctx, name)
	if err != nil {
		return nil, err
	}
	bsi := bsiObj.(*backingserviceinstanceapi.BackingServiceInstance)


	if len(bsi.Annotations) > 0 {
		for dc, bound := range bsi.Annotations {
			if bound == backingserviceinstanceapi.BindDeploymentConfigBound {
				return nil, fmt.Errorf("'%s' is bound to this instance, unbind it first.", dc)
			}
		}
	}

	if bsi.DeletionTimestamp.IsZero() {
		now := unversioned.Now()
		bsi.DeletionTimestamp = &now

		bsi.Status.Action = backingserviceinstanceapi.BackingServiceInstanceActionToDelete

		result, _, err := r.store.Update(ctx, bsi)
		return result, err
	}

	return r.store.Delete(ctx, name, options)
}

func (r *REST) Watch(ctx kapi.Context, options *kapi.ListOptions) (watch.Interface, error) {
	return r.store.Watch(ctx, options)
}

//============================================

func NewBindingREST(bsir backingserviceinstanceregistry.Registry, dcr deployconfigregistry.Registry) *BindingREST {
	return &BindingREST{
		backingServiceInstanceRegistry: bsir,
		deployConfigRegistry:           dcr,
	}
}

type BindingREST struct {
	backingServiceInstanceRegistry backingserviceinstanceregistry.Registry
	deployConfigRegistry           deployconfigregistry.Registry
	store *etcdgeneric.Etcd
}

func (r *BindingREST) New() runtime.Object {
	return &backingserviceinstanceapi.BindingRequestOptions{}
}

func (r *BindingREST) Get(ctx kapi.Context, name string) (runtime.Object, error) {
	return r.store.Get(ctx, name)
}

func (r *BindingREST) List(ctx kapi.Context, options *kapi.ListOptions) (runtime.Object, error) {
	return r.store.List(ctx,options)
}

func (r *BindingREST) Create(ctx kapi.Context, obj runtime.Object) (runtime.Object, error) {
	glog.Infoln("to create a bsi binding.")

	//if err := rest.BeforeCreate(backingserviceinstanceregistry.BsiStrategy, ctx, obj); err != nil {
	//	return nil, err
	//}

	bro, ok := obj.(*backingserviceinstanceapi.BindingRequestOptions)
	if !ok || bro.BindKind != backingserviceinstanceapi.BindKind_DeploymentConfig {
		return nil, fmt.Errorf("unsupported bind type: %s", bro.BindKind)
	}
	// todo: check bro.BindResourceVersion

	//kapi.FillObjectMetaSystemFields(ctx, &bro.ObjectMeta)

	bsi, err := r.backingServiceInstanceRegistry.GetBackingServiceInstance(ctx, bro.Name)
	if err != nil {
		return nil, err
	}

	if bound := bsi.Annotations[bro.ResourceName]; bound == backingserviceinstanceapi.BindDeploymentConfigBound {
		return nil, fmt.Errorf("'%s' is already bound to this instance.", bro.ResourceName)
	}
	/*
		if bsi.Status.Phase != backingserviceinstanceapi.BackingServiceInstancePhaseUnbound {
			return nil, errors.New("back service instance is not in unbound phase")
		}
	*/

	//bs, err := r.backingServiceRegistry.GetBackingService(ctx, bsi.Spec.BackingServiceName)
	//if err != nil {
	//	return nil, err
	//}

	dc, err := r.deployConfigRegistry.GetDeploymentConfig(ctx, bro.ResourceName)
	if err != nil {
		return nil, err
	}
	_ = dc

	// update bsi

	//need debug....bsi.Spec.BindDeploymentConfig = bro.ResourceName // dc.Name

	if bsi.Annotations == nil {
		bsi.Annotations = map[string]string{}
	}
	bsi.Annotations[bro.ResourceName] = backingserviceinstanceapi.BindDeploymentConfigBinding

	bsi.Status.Action = backingserviceinstanceapi.BackingServiceInstanceActionToBind

	bsi, err = r.backingServiceInstanceRegistry.UpdateBackingServiceInstance(ctx, bsi)
	if err != nil {
		return nil, err
	}

	// ...

	return bsi, nil
}

func (r *BindingREST) Update(ctx kapi.Context, obj runtime.Object) (runtime.Object, bool, error) {
	bro, ok := obj.(*backingserviceinstanceapi.BindingRequestOptions)
	if !ok || bro.BindKind != backingserviceinstanceapi.BindKind_DeploymentConfig {
		return nil, false, fmt.Errorf("unsupported bind type: '%s'", bro.BindKind)
	}

	bsi, err := r.backingServiceInstanceRegistry.GetBackingServiceInstance(ctx, bro.Name)
	if err != nil {
		return nil, false, err
	}

	if bound, ok := bsi.Annotations[bro.ResourceName]; !ok || bound == "unbound"/*unbound should never happen.*/ {
		return nil, false, fmt.Errorf("DeploymentConfig '%s' not bound to this instance yet.", bro.ResourceName)
	} else {
		bsi.Annotations[bro.ResourceName] = backingserviceinstanceapi.BindDeploymentConfigUnbinding
		bsi.Status.Action = backingserviceinstanceapi.BackingServiceInstanceActionToUnbind
	}
	bsi, err = r.backingServiceInstanceRegistry.UpdateBackingServiceInstance(ctx, bsi)
	if err != nil {
		return nil, false, err
	}
	return bsi, true, nil
}

func (r *BindingREST) Delete(ctx kapi.Context, name string, options *kapi.DeleteOptions) (runtime.Object, error) {
	glog.Infoln("to delete a bsi binding")

	bsi, err := r.backingServiceInstanceRegistry.GetBackingServiceInstance(ctx, name)
	if err != nil {
		return nil, err
	}

	if bsi.Status.Phase != backingserviceinstanceapi.BackingServiceInstancePhaseBound {
		return nil, errors.New("back service instance is not bound yet")
	}

	bsi.Status.Action = backingserviceinstanceapi.BackingServiceInstanceActionToUnbind
	//bsi.Annotations[]

	bsi, err = r.backingServiceInstanceRegistry.UpdateBackingServiceInstance(ctx, bsi)
	if err != nil {
		return nil, err
	}

	return bsi, nil
}
