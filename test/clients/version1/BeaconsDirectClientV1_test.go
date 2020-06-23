package test_clients

import (
	"testing"

	bclients "github.com/pip-services-samples/pip-client-microservice-go/clients/version1"
	blogic "github.com/pip-services-samples/pip-data-microservice-go/logic"
	bpersist "github.com/pip-services-samples/pip-data-microservice-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

func TestBeaconsDirectClientV1(t *testing.T) {

	var persistence *bpersist.BeaconsMemoryPersistence
	var controller *blogic.BeaconsController
	var client *bclients.BeaconsDirectClientV1
	var fixture *BeaconsClientV1Fixture

	persistence = bpersist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = blogic.NewBeaconsController()
	controller.Configure(cconf.NewEmptyConfigParams())
	client = bclients.NewBeaconsDirectClientV1()

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("beacons", "client", "direct", "default", "1.0"), client,
	)

	controller.SetReferences(references)
	client.SetReferences(references)
	fixture = NewBeaconsClientV1Fixture(client)

	persistence.Open("")
	defer persistence.Close("")

	t.Run("TestBeaconsDirectClientV1:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("TestBeaconsDirectClientV:1Calculate Positions", fixture.TestCalculatePosition)

}
