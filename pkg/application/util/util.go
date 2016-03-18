package util

import (
	"errors"
	applicationapi "github.com/openshift/origin/pkg/application/api"
	"strings"
)

func Contains(arr []string, str string) bool {
	for _, ele := range arr {
		if ele == str {
			return true
		}
	}

	return false
}

func Parse(items string) (applicationapi.ItemList, error) {
	list := applicationapi.ItemList{}
	arr := strings.Split(items, ",")

	for _, s := range arr {
		item := strings.Split(strings.TrimSpace(s), "=")
		if len(item) != 2 {
			return nil, errors.New("items wrong format")
		}

		list = append(list, applicationapi.Item{
			Kind: expandKindShortcut(item[0]),
			Name: item[1],
		})
	}

	return list, nil
}

//todo make elegant
func expandKindShortcut(kind string) string {
	shortForms := map[string]string{
		"dc":      "DeploymentConfig",
		"bc":      "BuildConfig",
		"is":      "ImageStream",
		"istag":   "ImageStreamTag",
		"isimage": "ImageStreamImage",
		"pv":      "PersistentVolume",
		"pvc":     "PersistentVolumeClaim",
		"rc":      "ReplicationController",
		"no":      "Node",
		"po":      "Pod",
		"svc":     "Service",
		"ev":      "Event",
		"bs":      "BackingService",
		"sb":      "ServiceBroker",
		"bsi":     "BackingServiceInstance",
	}
	if expanded, ok := shortForms[kind]; ok {
		return expanded
	}
	return kind
}