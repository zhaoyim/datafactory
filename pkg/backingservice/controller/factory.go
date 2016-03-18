package controller

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/cache"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/runtime"
	kutil "k8s.io/kubernetes/pkg/util"
	"k8s.io/kubernetes/pkg/watch"
	"time"

	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
	osclient "github.com/openshift/origin/pkg/client"
	controller "github.com/openshift/origin/pkg/controller"
	"k8s.io/kubernetes/pkg/client/record"
)

type BackingServiceControllerFactory struct {
	// Client is an OpenShift client.
	Client osclient.Interface
	// KubeClient is a Kubernetes client.
	KubeClient kclient.Interface
}

// Create creates a BackingServiceControllerFactory.
func (factory *BackingServiceControllerFactory) Create() controller.RunnableController {

	backingserviceLW := &cache.ListWatch{
		ListFunc: func() (runtime.Object, error) {

			return factory.Client.BackingServices(kapi.NamespaceAll).List(labels.Everything(), fields.Everything())

			//return factory.KubeClient.Namespaces().List(labels.Everything(), fields.Everything())
		},
		WatchFunc: func(resourceVersion string) (watch.Interface, error) {
			return factory.Client.BackingServices(kapi.NamespaceAll).Watch(labels.Everything(), fields.Everything(), resourceVersion)
			//return factory.KubeClient.Namespaces().Watch(labels.Everything(), fields.Everything(), resourceVersion)
		},
	}
	queue := cache.NewFIFO(cache.MetaNamespaceKeyFunc)
	cache.NewReflector(backingserviceLW, &backingserviceapi.BackingService{}, queue, 1*time.Minute).Run()

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(factory.KubeClient.Events(""))

	backingserviceController := &BackingServiceController{
		Client:     factory.Client,
		KubeClient: factory.KubeClient,
		recorder:   eventBroadcaster.NewRecorder(kapi.EventSource{Component: "BackingService"}),
	}

	return &controller.RetryController{
		Queue: queue,
		RetryManager: controller.NewQueueRetryManager(
			queue,
			cache.MetaNamespaceKeyFunc,
			func(obj interface{}, err error, retries controller.Retry) bool {
				kutil.HandleError(err)
				if _, isFatal := err.(fatalError); isFatal {
					return false
				}
				if retries.Count > 0 {
					return false
				}
				return true
			},
			kutil.NewTokenBucketRateLimiter(1, 10),
		),
		Handle: func(obj interface{}) error {
			backingservice := obj.(*backingserviceapi.BackingService)
			return backingserviceController.Handle(backingservice)
		},
	}
}

// buildConfigLW is a ListWatcher implementation for BuildConfigs.
type backingServiceLW struct {
	client osclient.Interface
}

// List lists all BuildConfigs.
func (lw *backingServiceLW) List() (runtime.Object, error) {
	return lw.client.BackingServices(kapi.NamespaceAll).List(labels.Everything(), fields.Everything())
}

// Watch watches all BuildConfigs.
func (lw *backingServiceLW) Watch(resourceVersion string) (watch.Interface, error) {
	return lw.client.BackingServices(kapi.NamespaceAll).Watch(labels.Everything(), fields.Everything(), resourceVersion)
}
