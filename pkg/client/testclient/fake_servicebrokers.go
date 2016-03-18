package testclient

import (
	ktestclient "k8s.io/kubernetes/pkg/client/unversioned/testclient"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"

	servicebrokerapi "github.com/openshift/origin/pkg/servicebroker/api"
)

// FakeServiceBrokers implements ServiceBrokerInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeServiceBrokers struct {
	Fake *Fake
}

func (c *FakeServiceBrokers) Get(name string) (*servicebrokerapi.ServiceBroker, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootGetAction("servicebroker", name), &servicebrokerapi.ServiceBroker{})
	if obj == nil {
		return nil, err
	}

	return obj.(*servicebrokerapi.ServiceBroker), err
}

func (c *FakeServiceBrokers) List(label labels.Selector, field fields.Selector) (*servicebrokerapi.ServiceBrokerList, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootListAction("servicebroker", label, field), &servicebrokerapi.ServiceBrokerList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*servicebrokerapi.ServiceBrokerList), err
}

func (c *FakeServiceBrokers) Create(inObj *servicebrokerapi.ServiceBroker) (*servicebrokerapi.ServiceBroker, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootCreateAction("servicebroker", inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*servicebrokerapi.ServiceBroker), err
}

func (c *FakeServiceBrokers) Update(inObj *servicebrokerapi.ServiceBroker) (*servicebrokerapi.ServiceBroker, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewRootUpdateAction("servicebroker", inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*servicebrokerapi.ServiceBroker), err
}

func (c *FakeServiceBrokers) Delete(name string) error {
	_, err := c.Fake.Invokes(ktestclient.NewRootDeleteAction("servicebroker", name), &servicebrokerapi.ServiceBroker{})
	return err
}
