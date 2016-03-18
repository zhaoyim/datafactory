package api

import "k8s.io/kubernetes/pkg/fields"

func ServiceBrokerToSelectableFields(serviceBroker *ServiceBroker) fields.Set {
	return fields.Set{
		"metadata.name": serviceBroker.Name,
	}
}
