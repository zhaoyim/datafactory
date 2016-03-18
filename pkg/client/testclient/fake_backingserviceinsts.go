package testclient

import (
	ktestclient "k8s.io/kubernetes/pkg/client/unversioned/testclient"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"

	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
)

// FakeBackingServiceInstances implements BackingServiceInstanceInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeBackingServiceInstances struct {
	Fake *Fake
}

func (c *FakeBackingServiceInstances) Get(name string) (*backingserviceinstanceapi.BackingServiceInstance, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootGetAction("backingserviceinstances", name), &backingserviceinstanceapi.BackingServiceInstance{})
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceinstanceapi.BackingServiceInstance), err
}

func (c *FakeBackingServiceInstances) List(label labels.Selector, field fields.Selector) (*backingserviceinstanceapi.BackingServiceInstanceList, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootListAction("backingserviceinstances", label, field), &backingserviceinstanceapi.BackingServiceInstanceList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceinstanceapi.BackingServiceInstanceList), err
}

func (c *FakeBackingServiceInstances) Create(inObj *backingserviceinstanceapi.BackingServiceInstance) (*backingserviceinstanceapi.BackingServiceInstance, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootCreateAction("backingserviceinstances", inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceinstanceapi.BackingServiceInstance), err
}

func (c *FakeBackingServiceInstances) Update(inObj *backingserviceinstanceapi.BackingServiceInstance) (*backingserviceinstanceapi.BackingServiceInstance, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootUpdateAction("backingserviceinstances", inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceinstanceapi.BackingServiceInstance), err
}

func (c *FakeBackingServiceInstances) Delete(name string) error {
	_, err := c.Fake.Invokes(ktestclient.NewRootDeleteAction("backingserviceinstances", name), &backingserviceinstanceapi.BackingServiceInstance{})
	return err
}
