package backingserviceinstance

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/rest"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/watch"

	"github.com/openshift/origin/pkg/backingserviceinstance/api"
)

// Registry is an interface for things that know how to store ImageStream objects.
type Registry interface {
	// ListImageStreams obtains a list of image streams that match a selector.
	ListBackingServiceInstances(ctx kapi.Context, options *kapi.ListOptions) (*api.BackingServiceInstanceList, error)
	// GetImageStream retrieves a specific image stream.
	GetBackingServiceInstance(ctx kapi.Context, id string) (*api.BackingServiceInstance, error)
	// CreateImageStream creates a new image stream.
	CreateBackingServiceInstance(ctx kapi.Context, repo *api.BackingServiceInstance) (*api.BackingServiceInstance, error)
	// UpdateImageStream updates an image stream.
	UpdateBackingServiceInstance(ctx kapi.Context, repo *api.BackingServiceInstance) (*api.BackingServiceInstance, error)
	// UpdateImageStreamSpec updates an image stream's spec.
	//UpdateBackingServiceInstanceSpec(ctx kapi.Context, repo *api.BackingServiceInstance) (*api.BackingServiceInstance, error)
	// UpdateImageStreamStatus updates an image stream's status.
	//UpdateBackingServiceInstanceStatus(ctx kapi.Context, repo *api.BackingServiceInstance) (*api.BackingServiceInstance, error)
	// DeleteImageStream deletes an image stream.
	DeleteBackingServiceInstance(ctx kapi.Context, id string) (*unversioned.Status, error)
	// WatchImageStreams watches for new/changed/deleted image streams.
	WatchBackingServiceInstances(ctx kapi.Context, options *kapi.ListOptions) (watch.Interface, error)
}

// Storage is an interface for a standard REST Storage backend
type Storage interface {
	rest.GracefulDeleter
	rest.Lister
	rest.Getter
	rest.Watcher

	Create(ctx kapi.Context, obj runtime.Object) (runtime.Object, error)
	Update(ctx kapi.Context, obj runtime.Object) (runtime.Object, bool, error)
}

// storage puts strong typing around storage calls
type storage struct {
	Storage
	//status   rest.Updater
	//internal rest.Updater
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched
// types will panic.
//func NewRegistry(s Storage, status, internal rest.Updater) Registry {
//	return &storage{Storage: s, status: status, internal: internal}
//}
func NewRegistry(s Storage) Registry {
	return &storage{Storage: s}
}

func (s *storage) ListBackingServiceInstances(ctx kapi.Context,options *kapi.ListOptions) (*api.BackingServiceInstanceList, error) {
	obj, err := s.List(ctx,options)
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingServiceInstanceList), nil
}

func (s *storage) GetBackingServiceInstance(ctx kapi.Context, backingServiceInstanceID string) (*api.BackingServiceInstance, error) {
	obj, err := s.Get(ctx, backingServiceInstanceID)
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingServiceInstance), nil
}

func (s *storage) CreateBackingServiceInstance(ctx kapi.Context, backingserviceinstance *api.BackingServiceInstance) (*api.BackingServiceInstance, error) {
	obj, err := s.Create(ctx, backingserviceinstance)
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingServiceInstance), nil
}

//func (s *storage) UpdateBackingServiceInstance(ctx kapi.Context, backingServiceInstance *api.BackingServiceInstance) (*api.BackingServiceInstance, error) {
//	obj, _, err := s.internal.Update(ctx, backingServiceInstance)
//	if err != nil {
//		return nil, err
//	}
//	return obj.(*api.BackingServiceInstance), nil
//}

func (s *storage) UpdateBackingServiceInstance(ctx kapi.Context, backingServiceInstance *api.BackingServiceInstance) (*api.BackingServiceInstance, error) {
	obj, _, err := s.Update(ctx, backingServiceInstance)
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingServiceInstance), nil
}

//func (s *storage) UpdateBackingServiceInstanceSpec(ctx kapi.Context, backingServiceInstance *api.BackingServiceInstance) (*api.BackingServiceInstance, error) {
//	obj, _, err := s.Update(ctx, backingServiceInstance)
//	if err != nil {
//		return nil, err
//	}
//	return obj.(*api.BackingServiceInstance), nil
//}

//func (s *storage) UpdateBackingServiceInstanceStatus(ctx kapi.Context, backingServiceInstance *api.BackingServiceInstance) (*api.BackingServiceInstance, error) {
//	obj, _, err := s.status.Update(ctx, backingServiceInstance)
//	if err != nil {
//		return nil, err
//	}
//	return obj.(*api.BackingServiceInstance), nil
//}

func (s *storage) DeleteBackingServiceInstance(ctx kapi.Context, backingServiceInstanceID string) (*unversioned.Status, error) {
	obj, err := s.Delete(ctx, backingServiceInstanceID, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*unversioned.Status), nil
}

func (s *storage) WatchBackingServiceInstances(ctx kapi.Context, options *kapi.ListOptions) (watch.Interface, error) {
	return s.Watch(ctx, options)
}
