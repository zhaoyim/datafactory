package testclient

import (
	kapi "k8s.io/kubernetes/pkg/api"
	ktestclient "k8s.io/kubernetes/pkg/client/unversioned/testclient"
	//"k8s.io/kubernetes/pkg/fields"
	//"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/watch"

	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
)

// FakeBackingServices implements BackingServiceInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeBackingServices struct {
	Fake *Fake
}

func (c *FakeBackingServices) Get(name string) (*backingserviceapi.BackingService, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootGetAction("backingservices", name), &backingserviceapi.BackingService{})
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceapi.BackingService), err
}

func (c *FakeBackingServices) List(opts kapi.ListOptions) (*backingserviceapi.BackingServiceList, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootListAction("backingservices", opts), &backingserviceapi.BackingServiceList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceapi.BackingServiceList), err
}

func (c *FakeBackingServices) Create(inObj *backingserviceapi.BackingService) (*backingserviceapi.BackingService, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootCreateAction("backingservices", inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceapi.BackingService), err
}

func (c *FakeBackingServices) Update(inObj *backingserviceapi.BackingService) (*backingserviceapi.BackingService, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootUpdateAction("backingservices", inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceapi.BackingService), err
}

func (c *FakeBackingServices) Delete(name string) error {
	_, err := c.Fake.Invokes(ktestclient.NewRootDeleteAction("backingservices", name), &backingserviceapi.BackingService{})
	return err
}

func (c *FakeBackingServices) Watch(opts kapi.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(ktestclient.NewRootWatchAction("backingservices", opts))
}
