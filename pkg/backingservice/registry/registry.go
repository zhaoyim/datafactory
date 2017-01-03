package backingservice

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/rest"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/watch"

	"github.com/openshift/origin/pkg/backingservice/api"
)

// Registry is an interface for things that know how to store ImageStream objects.
type Registry interface {
	// ListImageStreams obtains a list of image streams that match a selector.
	ListBackingServices(ctx kapi.Context) (*api.BackingServiceList, error)
	// GetImageStream retrieves a specific image stream.
	GetBackingService(ctx kapi.Context, name string) (*api.BackingService, error)
	// CreateImageStream creates a new image stream.
	CreateBackingService(ctx kapi.Context, bs *api.BackingService) (*api.BackingService, error)
	// UpdateImageStream updates an image stream.
	UpdateBackingService(ctx kapi.Context, bs *api.BackingService) (*api.BackingService, error)
	// UpdateImageStreamSpec updates an image stream's spec.
	UpdateBackingServiceSpec(ctx kapi.Context, bs *api.BackingService) (*api.BackingService, error)
	// UpdateImageStreamStatus updates an image stream's status.
	UpdateBackingServiceStatus(ctx kapi.Context, bs *api.BackingService) (*api.BackingService, error)
	// DeleteImageStream deletes an image stream.
	DeleteBackingService(ctx kapi.Context, name string) (*unversioned.Status, error)
	// WatchImageStreams watches for new/changed/deleted image streams.
	WatchBackingServices(ctx kapi.Context, options *kapi.ListOptions) (watch.Interface, error)
}

// storage puts strong typing around storage calls
type storage struct {
	rest.StandardStorage
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched
// types will panic.
func NewRegistry(s rest.StandardStorage) Registry {
	return &storage{s}
}

func (s *storage) ListBackingServices(ctx kapi.Context) (*api.BackingServiceList, error) {
	obj, err := s.List(ctx, &kapi.ListOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingServiceList), nil
}

func (s *storage) GetBackingService(ctx kapi.Context, name string) (*api.BackingService, error) {
	obj, err := s.Get(ctx, name)
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingService), nil
}

func (s *storage) CreateBackingService(ctx kapi.Context, backingservice *api.BackingService) (*api.BackingService, error) {
	obj, err := s.Create(ctx, backingservice)
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingService), nil
}

func (s *storage) UpdateBackingService(ctx kapi.Context, backingService *api.BackingService) (*api.BackingService, error) {
	obj, _, err := s.internal.Update(ctx, backingService)
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingService), nil
}

func (s *storage) UpdateBackingServiceSpec(ctx kapi.Context, backingService *api.BackingService) (*api.BackingService, error) {
	obj, _, err := s.Update(ctx, backingService)
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingService), nil
}

func (s *storage) UpdateBackingServiceStatus(ctx kapi.Context, backingService *api.BackingService) (*api.BackingService, error) {
	obj, _, err := s.status.Update(ctx, backingService)
	if err != nil {
		return nil, err
	}
	return obj.(*api.BackingService), nil
}

func (s *storage) DeleteBackingService(ctx kapi.Context, backingServiceID string) (*unversioned.Status, error) {
	obj, err := s.Delete(ctx, backingServiceID, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*unversioned.Status), nil
}

func (s *storage) WatchBackingServices(ctx kapi.Context, options *kapi.ListOptions) (watch.Interface, error) {
	return s.Watch(ctx, options)
}
