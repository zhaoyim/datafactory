package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_BackingServiceInstance = map[string]string{
	"":         "BackingServiceInstance describe a BackingServiceInstance",
	"metadata": "Standard object's metadata.",
	"spec":     "Spec defines the behavior of the Namespace.",
	"status":   "Status describes the current status of a Namespace",
}

func (BackingServiceInstance) SwaggerDoc() map[string]string {
	return map_BackingServiceInstance
}

var map_BackingServiceInstanceList = map[string]string{
	"":         "BackingServiceInstanceList describe a list of BackingServiceInstance",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of routes",
}

func (BackingServiceInstanceList) SwaggerDoc() map[string]string {
	return map_BackingServiceInstanceList
}

var map_BackingServiceInstanceSpec = map[string]string{
	"":             "BackingServiceInstanceSpec describes the attributes on a BackingServiceInstance",
	"provisioning": "description of an instance.",
	"binding":      "bindings of an instance",
	"bound":        "binding number of an instance",
	"instance_id":  "id of an instance",
	"tags":         "tags of an instance",
}

func (BackingServiceInstanceSpec) SwaggerDoc() map[string]string {
	return map_BackingServiceInstanceSpec
}

var map_BackingServiceInstanceStatus = map[string]string{
	"":               "BackingServiceInstanceStatus describe the status of a BackingServiceInstance",
	"phase":          "phase is the current lifecycle phase of the instance",
	"action":         "action is the action of the instance",
	"last_operation": "last operation  of a instance provisioning",
}

func (BackingServiceInstanceStatus) SwaggerDoc() map[string]string {
	return map_BackingServiceInstanceStatus
}

var map_BindingRequestOptions = map[string]string{
	"":                    "BindingRequestOptions describe a BindingRequestOptions.",
	"metadata":            "Standard object's metadata.",
	"bindKind":            "bind kind is bindking of an instance binding",
	"bindResourceVersion": "bindResourceVersion is bindResourceVersion of an instance binding.",
	"resourceName":        "resourceName of an instance binding",
}

func (BindingRequestOptions) SwaggerDoc() map[string]string {
	return map_BindingRequestOptions
}

var map_InstanceBinding = map[string]string{
	"":                      "InstanceBinding describe an instance binding.",
	"bound_time":            "bound time of an instance binding",
	"bind_uuid":             "bind uid of an instance binding",
	"bind_deploymentconfig": "deploymentconfig of an binding.",
	"credentials":           "credentials of an instance binding",
}

func (InstanceBinding) SwaggerDoc() map[string]string {
	return map_InstanceBinding
}

var map_InstanceProvisioning = map[string]string{
	"":                         "InstanceProvisioning describe an InstanceProvisioning detail",
	"dashboard_url":            "dashboard url of an instance",
	"backingservice_name":      "bs name of an instance",
	"backingservice_spec_id":   "bs id of an instance",
	"backingservice_plan_guid": "bs plan id of an instance",
	"backingservice_plan_name": "bs plan name of an instance",
	"parameters":               "parameters of an instance",
}

func (InstanceProvisioning) SwaggerDoc() map[string]string {
	return map_InstanceProvisioning
}

var map_LastOperation = map[string]string{
	"":                            "LastOperation describe last operation of an instance provisioning",
	"state":                       "state of last operation",
	"description":                 "description of last operation",
	"async_poll_interval_seconds": "async_poll_interval_seconds of a last operation",
}

func (LastOperation) SwaggerDoc() map[string]string {
	return map_LastOperation
}
