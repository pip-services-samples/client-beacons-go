package test_clients

import (
	"testing"

	bclients "github.com/pip-services-samples/pip-client-microservice-go/clients/version1"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

func TestBeaconsHttpClientExtV1(t *testing.T) {

	var client *bclients.BeaconsHttpClientV1
	var fixture *BeaconsClientV1Fixture

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "8080",
		"connection.host", "localhost",
	)

	client = bclients.NewBeaconsHttpClientV1()
	client.Configure(httpConfig)

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("beacons", "client", "http", "default", "1.0"), client,
	)

	client.SetReferences(references)

	fixture = NewBeaconsClientV1Fixture(client)

	client.Open("")

	defer client.Close("")

	t.Run("BeaconsHttpClientV1:CRUD Operations", fixture.TestCrudOperations)

	//t.Run("BeaconsHttpClientV1:1Calculate Positions", fixture.TestCalculatePosition)

}
