package controller

import (
	"fmt"
	"strings"

	api "github.com/openshift/origin/pkg/application/api"
	"k8s.io/kubernetes/pkg/labels"
)

func getLabelSelectorByApplication(application *api.Application) (labels.Selector, error) {

	return labels.Parse(fmt.Sprintf("%s.application.%s=%s", application.Namespace, application.Name, application.Name))
}

func labelExistsOtherApplicationKey(label map[string]string, labelString string) bool {
	m := filterMapByKey(label, ".application.", strings.Contains)

	if len(m) > 1 {
		if _, exist := m[labelString]; exist {
			return exist
		}
	}

	return false
}

func labelExistsApplicationKey(label map[string]string, keyString string) bool {

	m := filterMapByKey(label, ".application.", strings.Contains)

	if _, exist := m[keyString]; exist {
		return exist
	}

	return false
}

func filterMapByKey(m map[string]string, filterStr string, filterFn func(string, string) bool) map[string]string {
	newMap := make(map[string]string)

	if m != nil {
		for key := range m {
			if filterFn(key, filterStr) {
				newMap[key] = m[key]
			}
		}
	}

	return newMap
}

func hasItem(items api.ItemList, item api.Item) bool {
	for i := range items {
		if items[i].Kind == item.Kind && items[i].Name == item.Name {
			return true
		}
	}

	return false
}

func getItemErrEventReason(appStatus api.ApplicationStatus, item api.Item) string {
	return fmt.Sprintf("Application[%s] Handle Resource %s=%s", appStatus.Phase, item.Kind, item.Name)
}

func addItemEvent(item api.Item) string {
	return fmt.Sprintf("Add %s=%s", item.Kind, item.Name)
}