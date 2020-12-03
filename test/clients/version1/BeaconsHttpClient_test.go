package test_clients1

import (
	"testing"

	clients1 "github.com/pip-services-samples/pip-clients-beacons-go/clients/version1"
	blogic "github.com/pip-services-samples/pip-services-beacons-go/logic"
	bpersist "github.com/pip-services-samples/pip-services-beacons-go/persistence"
	bservices "github.com/pip-services-samples/pip-services-beacons-go/services/version1"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

func TestBeaconsHttpClientV1(t *testing.T) {

	var persistence *bpersist.BeaconsMemoryPersistence
	var controller *blogic.BeaconsController
	var service *bservices.BeaconsCommandableHttpServiceV1
	var client *clients1.BeaconsCommandableHttpClientV1
	var fixture *BeaconsClientV1Fixture

	persistence = bpersist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = blogic.NewBeaconsController()
	controller.Configure(cconf.NewEmptyConfigParams())

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3000",
		"connection.host", "localhost",
	)

	service = bservices.NewBeaconsCommandableHttpServiceV1()
	service.Configure(httpConfig)

	client = clients1.NewBeaconsCommandableHttpClientV1()
	client.Configure(httpConfig)

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("beacons", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("beacons", "client", "http", "default", "1.0"), client,
	)
	controller.SetReferences(references)
	service.SetReferences(references)
	client.SetReferences(references)

	fixture = NewBeaconsClientV1Fixture(client)

	opnErr := persistence.Open("")
	if opnErr != nil {
		panic("TestBeaconsHttpClientV1:Error open persistence!")
	}

	opnErr = service.Open("")
	if opnErr != nil {
		panic("TestBeaconsHttpClientV1:Error open service!")
	}

	client.Open("")

	defer client.Close("")
	defer service.Close("")
	defer persistence.Close("")

	t.Run("BeaconsHttpClientV1:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("BeaconsHttpClientV1:1Calculate Positions", fixture.TestCalculatePosition)

}
