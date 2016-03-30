package controller

import (
	"fmt"

	api "github.com/openshift/origin/pkg/application/api"
	osclient "github.com/openshift/origin/pkg/client"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	kapi "k8s.io/kubernetes/pkg/api"
)

func unloadServiceBrokerLabel(client osclient.Interface, application *api.Application, labelSelector labels.Selector) error {
	resourceList, _ := client.ServiceBrokers().List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "ServiceBroker", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.ServiceBrokers().Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}

func unloadBackingServiceInstanceLabel(client osclient.Interface, application *api.Application, labelSelector labels.Selector) error {

	resourceList, _ := client.BackingServiceInstances(application.Namespace).List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "BackingServiceInstance", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.BackingServiceInstances(application.Namespace).Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}

func unloadBuildLabel(client osclient.Interface, application *api.Application, labelSelector labels.Selector) error {
	resourceList, _ := client.Builds(application.Namespace).List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "Build", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.Builds(application.Namespace).Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}

func unloadBuildConfigLabel(client osclient.Interface, application *api.Application, labelSelector labels.Selector) error {

	resourceList, _ := client.BuildConfigs(application.Namespace).List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "BuildConfig", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.BuildConfigs(application.Namespace).Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}

func unloadDeploymentConfigLabel(client osclient.Interface, application *api.Application, labelSelector labels.Selector) error {

	resourceList, _ := client.DeploymentConfigs(application.Namespace).List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "DeploymentConfig", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.DeploymentConfigs(application.Namespace).Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}

func unloadReplicationControllerLabel(client kclient.Interface, application *api.Application, labelSelector labels.Selector) error {

	resourceList, _ := client.ReplicationControllers(application.Namespace).List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "ReplicationController", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.ReplicationControllers(application.Namespace).Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}

func unloadImageStreamLabel(client osclient.Interface, application *api.Application, labelSelector labels.Selector) error {

	resourceList, _ := client.ImageStreams(application.Namespace).List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "ImageStream", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.ImageStreams(application.Namespace).Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}

func unloadNodeLabel(client kclient.Interface, application *api.Application, labelSelector labels.Selector) error {

	resourceList, _ := client.Nodes().List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "Node", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.Nodes().Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}

func unloadPodLabel(client kclient.Interface, application *api.Application, labelSelector labels.Selector) error {

	resourceList, _ := client.Pods(application.Namespace).List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "Pod", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.Pods(application.Namespace).Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}

func unloadServiceLabel(client kclient.Interface, application *api.Application, labelSelector labels.Selector) error {

	resourceList, _ := client.Services(application.Namespace).List(kapi.ListOptions{LabelSelector:labelSelector,FieldSelector: fields.Everything()})
	errs := []error{}
	for _, resource := range resourceList.Items {
		if !hasItem(application.Spec.Items, api.Item{Kind: "Service", Name: resource.Name}) {
			delete(resource.Labels, fmt.Sprintf("%s.application.%s", application.Namespace, application.Name))
			if _, err := client.Services(application.Namespace).Update(&resource); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}
