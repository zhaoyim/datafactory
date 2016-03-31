package controller

import (
	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
	osclient "github.com/openshift/origin/pkg/client"
	"k8s.io/kubernetes/pkg/client/record"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	kapi "k8s.io/kubernetes/pkg/api"
)

// NamespaceController is responsible for participating in Kubernetes Namespace termination
// Use the NamespaceControllerFactory to create this controller.
type BackingServiceController struct {
	// Client is an OpenShift client.
	Client osclient.Interface
	// KubeClient is a Kubernetes client.
	KubeClient kclient.Interface
	recorder   record.EventRecorder
}

type fatalError string

func (e fatalError) Error() string {
	return "fatal error handling BackingServiceController: " + string(e)
}

// Handle processes a namespace and deletes content in origin if its terminating
func (c *BackingServiceController) Handle(bs *backingserviceapi.BackingService) (err error) {

	switch bs.Status.Phase {
	case backingserviceapi.BackingServicePhaseInactive:
		c.recorder.Eventf(bs, kapi.EventTypeNormal, "New", "'%s' is now %s!", bs.Name, bs.Status.Phase)
	case backingserviceapi.BackingServicePhaseActive:
	default:
		bs.Status.Phase = backingserviceapi.BackingServicePhaseActive

		c.recorder.Eventf(bs, kapi.EventTypeNormal, "New", "'%s' is now %s!", bs.Name, bs.Status.Phase)
		c.Client.BackingServices(bs.Namespace).Update(bs)
	}

	return nil
}
