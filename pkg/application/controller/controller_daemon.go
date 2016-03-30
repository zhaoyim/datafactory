package controller

import (
	"fmt"
	api "github.com/openshift/origin/pkg/application/api"
	kerrors "k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/runtime"
)

func errHandle(err error, app *api.Application, itemIndex int, resourceLabel map[string]string, eventf func(object runtime.Object, eventtype, reason, messageFmt string, args ...interface{})) {
	selectorKey := fmt.Sprintf("%s.application.%s", app.Namespace, app.Name)

	if err != nil {
		if kerrors.IsNotFound(err) && app.Spec.Items[itemIndex].Status != api.ApplicationItemDelete {
			eventf(app, "Health Check", "Application", "Application Unhealthy", "resource %s=%s has been deleted", app.Spec.Items[itemIndex].Kind, app.Spec.Items[itemIndex].Name)
			app.Spec.Items[itemIndex].Status = api.ApplicationItemDelete
			app.Status.Phase = api.ApplicationChecking
		}
	} else {
		if !labelExistsApplicationKey(resourceLabel, selectorKey) {
			eventf(app, "Health Check", "Application", "Application Unhealthy", "resource %s=%s label has been deleted", app.Spec.Items[itemIndex].Kind, app.Spec.Items[itemIndex].Name)
			app.Spec.Items[itemIndex].Status = api.ApplicationItemLabelDelete
			app.Status.Phase = api.ApplicationChecking
		}
	}
}
