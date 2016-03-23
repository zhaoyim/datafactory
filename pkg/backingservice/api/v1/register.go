package v1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
)

const GroupName = ""

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1"}

func AddToScheme(scheme *runtime.Scheme) {
	addKnownTypes(scheme)
	addConversionFuncs(scheme)
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&BackingService{},
		&BackingServiceList{},
	)
}

func (obj *BackingService) GetObjectKind() unversioned.ObjectKind     { return &obj.TypeMeta }
func (obj *BackingServiceList) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }

