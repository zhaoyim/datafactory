package controller

import (
	"fmt"
	api "github.com/openshift/origin/pkg/application/api"
	kerrors "k8s.io/kubernetes/pkg/api/errors"
)

func errHandle(err error, a *api.Application, itemIndex int, resourceLabel map[string]string) {
	selectorKey := fmt.Sprintf("%s.application.%s", a.Namespace, a.Name)

	if err != nil {
		if kerrors.IsNotFound(err) && a.Spec.Items[itemIndex].Status != api.ApplicationItemDelete {
			a.Spec.Items[itemIndex].Status = api.ApplicationItemDelete
			a.Status.Phase = api.ApplicationChecking
		}
	} else {
		if !labelExistsApplicationKey(resourceLabel, selectorKey) {
			a.Spec.Items[itemIndex].Status = api.ApplicationItemLabelDelete
			a.Status.Phase = api.ApplicationChecking
		}
	}
}
