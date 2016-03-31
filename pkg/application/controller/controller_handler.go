package controller

import (
	"fmt"
	"github.com/openshift/origin/pkg/application/api"
	kerrors "k8s.io/kubernetes/pkg/api/errors"
	deployapi "github.com/openshift/origin/pkg/deploy/api"
	deployutil "github.com/openshift/origin/pkg/deploy/util"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	kapi "k8s.io/kubernetes/pkg/api"
)

func (c *ApplicationController) handleServiceBrokerLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.Client.ServiceBrokers()

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app,kapi.EventTypeWarning, addItemEvent(app.Spec.Items[itemIndex]), "get servicebroker has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app, kapi.EventTypeWarning,addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app, kapi.EventTypeNormal, "Application", addItemEvent(app.Spec.Items[itemIndex]) + "success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete servicebroker has error: %s", err.Error())
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update servicebroker has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update servicebroker has error: %s", err.Error())
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app,kapi.EventTypeWarning,  "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) handleBackingServiceInstanceLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.Client.BackingServiceInstances(app.Namespace)

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app,kapi.EventTypeWarning,  getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "get backingserviceinstance has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app,kapi.EventTypeWarning,  addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app,kapi.EventTypeWarning,  "Application", addItemEvent(app.Spec.Items[itemIndex])+"success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete backingservice has error: %s", err.Error())
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update backingservice has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) handleBuildLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.Client.Builds(app.Namespace)

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app,kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "get build has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app,kapi.EventTypeNormal, "Application", addItemEvent(app.Spec.Items[itemIndex])+"success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				c.Recorder.Eventf(app,kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete build has error: %s", err.Error())
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app,kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update build has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app,kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) handleBuildConfigLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.Client.BuildConfigs(app.Namespace)

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app, kapi.EventTypeWarning,getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "get buildconfig has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app,kapi.EventTypeWarning, addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app, kapi.EventTypeWarning,"Application", addItemEvent(app.Spec.Items[itemIndex])+"success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning,getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete buildconfig has error: %s", err.Error())
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning,getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update buildconfig has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app,kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app,kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) handleDeploymentConfigLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.Client.DeploymentConfigs(app.Namespace)

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "get deploymentconfig has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name

			for _, trigger := range resource.Spec.Triggers {
				if trigger.Type == deployapi.DeploymentTriggerOnConfigChange {
					trigger.Type = deployapi.DeploymentTriggerManual
				}
			}

		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app,kapi.EventTypeWarning, "Application", addItemEvent(app.Spec.Items[itemIndex])+ "success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {

			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete deploymentconfig has error: %s", err.Error())
				return err
			}

			sel := deployutil.ConfigSelector(app.Spec.Items[itemIndex].Name)
			existingDeployments, err := c.KubeClient.ReplicationControllers(app.Namespace).List(kapi.ListOptions{LabelSelector: sel, FieldSelector:fields.Everything()})
			if err != nil && !kerrors.IsNotFound(err) {
				fmt.Printf("delete application dc %s err ", app.Spec.Items[itemIndex].Name, err)
			}

			for _, v := range existingDeployments.Items {
				if err := c.KubeClient.ReplicationControllers(app.Namespace).Delete(v.Name); err != nil {
					c.Recorder.Eventf(app,kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete deploymentconfig-replicationcontroller %s has error: %v", v.Name, err)
				}
				c.Recorder.Eventf(app,kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete deploymentconfig-replicationcontroller %s has error", v.Name)
			}

			existingPods, err := c.KubeClient.Pods(app.Namespace).List(kapi.ListOptions{LabelSelector: labels.Set{deployapi.DeploymentConfigLabel: app.Spec.Items[itemIndex].Name}.AsSelector(), FieldSelector: fields.Everything()})
			if err != nil && !kerrors.IsNotFound(err) {
				fmt.Printf("delete application dc %s err ", app.Spec.Items[itemIndex].Name, err)
			}

			for _, v := range existingPods.Items {
				if err := c.KubeClient.Pods(app.Namespace).Delete(v.Name, nil); err != nil {
					c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete deploymentconfig-pod %s has error: %v", v.Name, err)
				}
				c.Recorder.Eventf(app,kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete deploymentconfig-pod %s has error", v.Name)
			}

		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update deploymentconfig has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) handleImageStreamLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.Client.ImageStreams(app.Namespace)

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "get imagestream has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app, kapi.EventTypeWarning, "Application", addItemEvent(app.Spec.Items[itemIndex])+ "success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete imagestream has error: %s", err.Error())
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update imagestream has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) handleReplicationControllerLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.KubeClient.ReplicationControllers(app.Namespace)

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "get replicationcontroller has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app, kapi.EventTypeWarning, "Application", addItemEvent(app.Spec.Items[itemIndex])+ "success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete replicationcontroller has error: %s", err.Error())
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update replicationcontroller has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) handleNodeLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.KubeClient.Nodes()

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "get node has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app, kapi.EventTypeWarning, "Application", addItemEvent(app.Spec.Items[itemIndex])+ "success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete node has error: %s", err.Error())
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update replicationcontroller has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) handlePodLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.KubeClient.Pods(app.Namespace)

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "get pod has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app, "Application", addItemEvent(app.Spec.Items[itemIndex]), "success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name, nil); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete pod has error: %s", err.Error())
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update pod has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) handleServiceLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.KubeClient.Services(app.Namespace)

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "get service has error: %s", err.Error())
			c.deleteApplicationItem(app, itemIndex)
			return nil
		}
		return err
	}

	switch app.Status.Phase {
	case api.ApplicationActiveUpdate:
		if _, exists := resource.Labels[labelSelectorStr]; exists {
			//Active正常状态,当有新的更新时,如果这个label不存在,则新建
			return nil
		}
		fallthrough
	case api.ApplicationNew:
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}

		resource.Labels[labelSelectorStr] = app.Name
		if _, err := client.Update(resource); err != nil {
			c.Recorder.Eventf(app, kapi.EventTypeWarning, addItemEvent(app.Spec.Items[itemIndex]), "error: %s", err.Error())
			return err
		}
		c.Recorder.Event(app, kapi.EventTypeWarning, "Application", addItemEvent(app.Spec.Items[itemIndex])+ "success")

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "delete service has error: %s", err.Error())
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, getItemErrEventReason(app.Status, app.Spec.Items[itemIndex]), "update service has error: %s", err.Error())
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning,  "Clean Application", "delete application has error: %s", err.Error())
			}
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			if err := c.Client.Applications(app.Namespace).Delete(app.Name); err != nil {
				c.Recorder.Eventf(app, kapi.EventTypeWarning, "Clean Application", "delete application has error: %s", err.Error())
			}
		}
	}

	return nil
}

func (c *ApplicationController) deleteApplicationItem(app *api.Application, itemIndex int) {
	app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)
	c.Client.Applications(app.Namespace).Delete(app.Name)
}