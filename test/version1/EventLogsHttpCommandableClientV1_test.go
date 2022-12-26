package test_clients1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-infrastructure2/client-eventlog-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type eventLogCommandableHttpClientV1Test struct {
	client  *version1.EventLogCommandableHttpClientV1
	fixture *EventLogClientV1Fixture
}

func newEventLogCommandableHttpClientV1Test() *eventLogCommandableHttpClientV1Test {
	return &eventLogCommandableHttpClientV1Test{}
}

func (c *eventLogCommandableHttpClientV1Test) setup(t *testing.T) *EventLogClientV1Fixture {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewEventLogCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewEventLogClientV1Fixture(c.client)

	return c.fixture
}

func (c *eventLogCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newEventLogCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
