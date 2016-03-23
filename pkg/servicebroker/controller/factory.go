package controller

import (
	"k8s.io/kubernetes/pkg/client/cache"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
	kutil "k8s.io/kubernetes/pkg/util"
	utilruntime "k8s.io/kubernetes/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/watch"
	"time"
	kapi "k8s.io/kubernetes/pkg/api"
	osclient "github.com/openshift/origin/pkg/client"
	controller "github.com/openshift/origin/pkg/controller"
	servicebrokerapi "github.com/openshift/origin/pkg/servicebroker/api"
	servicebrokerclient "github.com/openshift/origin/pkg/servicebroker/client"
)

type ServiceBrokerControllerFactory struct {
	// Client is an OpenShift client.
	Client osclient.Interface
	// KubeClient is a Kubernetes client.
	KubeClient kclient.Interface
}

// Create creates a ServiceBrokerControllerFactory.
func (factory *ServiceBrokerControllerFactory) Create() controller.RunnableController {

	servicebrokerLW := &cache.ListWatch{
		ListFunc: func(options kapi.ListOptions) (runtime.Object, error) {

			return factory.Client.ServiceBrokers().List(options)

			//return factory.KubeClient.Namespaces().List(labels.Everything(), fields.Everything())
		},
		WatchFunc: func(options kapi.ListOptions) (watch.Interface, error) {
			return factory.Client.ServiceBrokers().Watch(options)
			//return factory.KubeClient.Namespaces().Watch(labels.Everything(), fields.Everything(), resourceVersion)
		},
	}
	queue := cache.NewFIFO(cache.MetaNamespaceKeyFunc)
	cache.NewReflector(servicebrokerLW, &servicebrokerapi.ServiceBroker{}, queue, 10 * time.Second).Run()

	servicebrokerController := &ServiceBrokerController{
		Client:              factory.Client,
		KubeClient:          factory.KubeClient,
		ServiceBrokerClient: servicebrokerclient.NewServiceBrokerClient(),
	}

	return &controller.RetryController{
		Queue: queue,
		RetryManager: controller.NewQueueRetryManager(
			queue,
			cache.MetaNamespaceKeyFunc,
			func(obj interface{}, err error, retries controller.Retry) bool {
				utilruntime.HandleError(err)
				if _, isFatal := err.(fatalError); isFatal {
					return false
				}
				if retries.Count > 0 {
					return false
				}
				return true
			},
			kutil.NewTokenBucketRateLimiter(10, 1),
		),
		Handle: func(obj interface{}) error {
			servicebroker := obj.(*servicebrokerapi.ServiceBroker)
			return servicebrokerController.Handle(servicebroker)
		},
	}
}

// buildConfigLW is a ListWatcher implementation for BuildConfigs.
type backingServiceLW struct {
	client osclient.Interface
}

// List lists all BuildConfigs.
func (lw *backingServiceLW) List() (runtime.Object, error) {
	return lw.client.ServiceBrokers().List(kapi.ListOptions{})
}

// Watch watches all BuildConfigs.
func (lw *backingServiceLW) Watch(resourceVersion string) (watch.Interface, error) {
	return lw.client.ServiceBrokers().Watch(kapi.ListOptions{})
}
