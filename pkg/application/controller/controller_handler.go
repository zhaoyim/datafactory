package controller

import (
	"fmt"
	api "github.com/openshift/origin/pkg/application/api"
	kerrors "k8s.io/kubernetes/pkg/api/errors"
)

func (c *ApplicationController) handleServiceBrokerLabel(app *api.Application, itemIndex int) error {
	labelSelectorStr := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	client := c.Client.ServiceBrokers()

	resource, err := client.Get(app.Spec.Items[itemIndex].Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name, nil); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
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
			return err
		}

	case api.ApplicationTerminating:
		if !labelExistsOtherApplicationKey(resource.Labels, labelSelectorStr) {
			if err := client.Delete(app.Spec.Items[itemIndex].Name); err != nil {
				return err
			}
		} else {
			delete(resource.Labels, labelSelectorStr)
			if _, err := client.Update(resource); err != nil {
				return err
			}
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}

	case api.ApplicationTerminatingLabel:
		delete(resource.Labels, labelSelectorStr)
		if _, err := client.Update(resource); err != nil {
			return err
		}

		app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)

		if len(app.Spec.Items) == 0 {
			c.Client.Applications(app.Namespace).Delete(app.Name)
		}
	}

	return nil
}

func (c *ApplicationController) deleteApplicationItem(app *api.Application, itemIndex int) {
	app.Spec.Items = append(app.Spec.Items[:itemIndex], app.Spec.Items[itemIndex + 1:]...)
	c.Client.Applications(app.Namespace).Delete(app.Name)
}
