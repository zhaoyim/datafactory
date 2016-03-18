package api

import "k8s.io/kubernetes/pkg/fields"

func ApplicationToSelectableFields(application *Application) fields.Set {
	return fields.Set{
		"metadata.name": application.Name,
	}
}
