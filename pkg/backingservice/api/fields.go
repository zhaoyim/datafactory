package api

import "k8s.io/kubernetes/pkg/fields"

// BackingServiceToSelectableFields returns a label set that represents the object
func BackingServiceToSelectableFields(backingService *BackingService) fields.Set {
	return fields.Set{
		"metadata.name": backingService.Name,
	}
}
