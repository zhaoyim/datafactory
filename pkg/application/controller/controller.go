package controller

import (
	"errors"
	api "github.com/openshift/origin/pkg/application/api"
	osclient "github.com/openshift/origin/pkg/client"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	errutil "k8s.io/kubernetes/pkg/util/errors"
)

// NamespaceController is responsible for participating in Kubernetes Namespace termination
// Use the NamespaceControllerFactory to create this controller.
type ApplicationController struct {
	// Client is an OpenShift client.
	Client osclient.Interface
	// KubeClient is a Kubernetes client.
	KubeClient kclient.Interface
}

type fatalError string

func (e fatalError) Error() string {
	return "fatal error handling ApplicationController: " + string(e)
}

// Handle processes a namespace and deletes content in origin if its terminating
func (c *ApplicationController) Handle(application *api.Application) (err error) {

	switch application.Status.Phase {
	case api.ApplicationTerminating:
		fallthrough

	case api.ApplicationTerminatingLabel:
		if err := c.handleAllLabel(application); err != nil {
			return err
		}

	case api.ApplicationActive:
		c.unifyDaemon(application)
		return nil

	case api.ApplicationActiveUpdate:
		if err := c.preHandleAllLabel(application); err != nil {
			return err
		}

		fallthrough

	default:
		if err := c.handleAllLabel(application); err != nil {
			return err
		}

		application.Status.Phase = api.ApplicationActive
		c.Client.Applications(application.Namespace).Update(application)

	}

	return nil
}

func (c *ApplicationController) unifyDaemon(application *api.Application) {
	for i := range application.Spec.Items {
		switch application.Spec.Items[i].Kind {
		case "ServiceBroker":
			resource, err := c.Client.ServiceBrokers().Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)

		case "BackingServiceInstance":
			resource, err := c.Client.BackingServiceInstances(application.Namespace).Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)

		case "Build":
			resource, err := c.Client.Builds(application.Namespace).Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)

		case "BuildConfig":
			resource, err := c.Client.BuildConfigs(application.Namespace).Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)

		case "DeploymentConfig":
			resource, err := c.Client.DeploymentConfigs(application.Namespace).Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)

		case "ReplicationController":
			resource, err := c.KubeClient.ReplicationControllers(application.Namespace).Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)

		case "ImageStream":
			resource, err := c.Client.ImageStreams(application.Namespace).Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)

		case "Node":
			resource, err := c.KubeClient.Nodes().Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)

		case "Pod":
			resource, err := c.KubeClient.Pods(application.Namespace).Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)

		case "Service":
			resource, err := c.KubeClient.Services(application.Namespace).Get(application.Spec.Items[i].Name)
			errHandle(err, application, i, resource.Labels)
		}

	}

	if application.Status.Phase == api.ApplicationChecking {
		c.Client.Applications(application.Namespace).Update(application)
	}
}

func (c *ApplicationController) preHandleAllLabel(application *api.Application) error {

	selector, err := getLabelSelectorByApplication(application)
	if err != nil {
		return err
	}

	errs := []error{}

	if err := unloadServiceBrokerLabel(c.Client, application, selector); err != nil {
		errs = append(errs, err)
	}

	if err := unloadBackingServiceInstanceLabel(c.Client, application, selector); err != nil {
		errs = append(errs, err)
	}

	if err := unloadBuildLabel(c.Client, application, selector); err != nil {
		errs = append(errs, err)
	}

	if err := unloadBuildConfigLabel(c.Client, application, selector); err != nil {
		errs = append(errs, err)
	}

	if err := unloadDeploymentConfigLabel(c.Client, application, selector); err != nil {
		errs = append(errs, err)
	}

	if err := unloadImageStreamLabel(c.Client, application, selector); err != nil {
		errs = append(errs, err)
	}

	if err := unloadReplicationControllerLabel(c.KubeClient, application, selector); err != nil {
		errs = append(errs, err)
	}

	if err := unloadNodeLabel(c.KubeClient, application, selector); err != nil {
		errs = append(errs, err)
	}

	if err := unloadPodLabel(c.KubeClient, application, selector); err != nil {
		errs = append(errs, err)
	}

	if err := unloadServiceLabel(c.KubeClient, application, selector); err != nil {
		errs = append(errs, err)
	}

	return errutil.NewAggregate(errs)
}

func (c *ApplicationController) handleAllLabel(app *api.Application) error {
	if c.destoryApplication(app) {
		return nil
	}

	errs := []error{}
	oldLength := len(app.Spec.Items)
	for i, item := range app.Spec.Items {
		newLength := len(app.Spec.Items)
		deleteNum := oldLength - newLength
		i = i - deleteNum

		switch item.Kind {
		case "ServiceBroker":
			if err := c.handleServiceBrokerLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		case "BackingServiceInstance":
			if err := c.handleBackingServiceInstanceLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		case "Build":
			if err := c.handleBuildLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		case "BuildConfig":
			if err := c.handleBuildConfigLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		case "DeploymentConfig":
			if err := c.handleDeploymentConfigLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		case "ReplicationController":
			if err := c.handleReplicationControllerLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		case "ImageStream":
			if err := c.handleImageStreamLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		case "Node":
			if err := c.handleNodeLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		case "Pod":
			if err := c.handlePodLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		case "Service":
			if err := c.handleServiceLabel(app, i); err != nil {
				errs = append(errs, err)
			}
		default:
			errs = append(errs, errors.New("unknown resource "+item.Kind+"="+item.Name))
		}
	}

	return errutil.NewAggregate(errs)
}

func (c *ApplicationController) destoryApplication(app *api.Application) bool {
	if len(app.Spec.Items) == 0 {
		if app.Status.Phase == api.ApplicationTerminatingLabel || app.Status.Phase == api.ApplicationTerminating {
			c.Client.Applications(app.Namespace).Delete(app.Name)
			return true
		}
	}
	return false
}
