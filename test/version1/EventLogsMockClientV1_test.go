package test_clients1

import (
	"testing"

	"github.com/pip-services-infrastructure2/client-eventlog-go/version1"
)

type eventLogMockClientV1Test struct {
	client  *version1.EventLogMockClientV1
	fixture *EventLogClientV1Fixture
}

func neweventLogMockClientV1Test() *eventLogMockClientV1Test {
	return &eventLogMockClientV1Test{}
}

func (c *eventLogMockClientV1Test) setup(t *testing.T) *EventLogClientV1Fixture {
	c.client = version1.NewEventLogMockClientV1()

	c.fixture = NewEventLogClientV1Fixture(c.client)

	return c.fixture
}

func (c *eventLogMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := neweventLogMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
