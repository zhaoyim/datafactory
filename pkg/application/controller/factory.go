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

	applicationapi "github.com/openshift/origin/pkg/application/api"
	osclient "github.com/openshift/origin/pkg/client"
	controller "github.com/openshift/origin/pkg/controller"
)

type ApplicationControllerFactory struct {
	// Client is an OpenShift client.
	Client osclient.Interface
	// KubeClient is a Kubernetes client.
	KubeClient kclient.Interface
}

// Create creates a ApplicationControllerFactory.
func (factory *ApplicationControllerFactory) Create() controller.RunnableController {

	applicationLW := &cache.ListWatch{
		ListFunc: func() (runtime.Object, error) {
			return factory.Client.Applications(kapi.NamespaceAll).List(labels.Everything(), fields.Everything())

		},
		WatchFunc: func(resourceVersion string) (watch.Interface, error) {
			return factory.Client.Applications(kapi.NamespaceAll).Watch(labels.Everything(), fields.Everything(), resourceVersion)
		},
	}
	queue := cache.NewFIFO(cache.MetaNamespaceKeyFunc)
	cache.NewReflector(applicationLW, &applicationapi.Application{}, queue, 10 * time.Second).Run()

	applicationController := &ApplicationController{
		Client:     factory.Client,
		KubeClient: factory.KubeClient,
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
			kutil.NewTokenBucketRateLimiter(10, 1),
		),
		Handle: func(obj interface{}) error {
			application := obj.(*applicationapi.Application)
			return applicationController.Handle(application)
		},
	}
}
