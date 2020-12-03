package test_clients1

import (
	"testing"

	clients1 "github.com/pip-services-samples/pip-clients-beacons-go/clients/version1"
)

type beaconsMemoryClientV1Test struct {
	client  *clients1.BeaconsMemoryClientV1
	fixture *BeaconsClientV1Fixture
}

func newBeaconsMemoryClientV1Test() *beaconsMemoryClientV1Test {

	client := clients1.NewBeaconsMemoryClientV1(nil)

	fixture := NewBeaconsClientV1Fixture(client)

	return &beaconsMemoryClientV1Test{
		client:  client,
		fixture: fixture,
	}
}

func (c *beaconsMemoryClientV1Test) setup(t *testing.T) {

	err := c.client.Open("")
	if err != nil {
		t.Error("Failed to open client", err)
	}
	err = c.client.Clear("")
	if err != nil {
		t.Error("Failed to open client", err)
	}

}

func (c *beaconsMemoryClientV1Test) teardown(t *testing.T) {
	err := c.client.Close("")
	if err != nil {
		t.Error("Failed to close client", err)
	}
}

func TestBeaconsMemoryClientV1(t *testing.T) {
	c := newBeaconsMemoryClientV1Test()

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("Calculate Positions", c.fixture.TestCalculatePosition)
	c.teardown(t)
}
