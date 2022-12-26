package build

import (
	"github.com/pip-services-infrastructure2/client-eventlog-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type EventLogClientFactory struct {
	cbuild.Factory
}

func NewEventLogClientFactory() *EventLogClientFactory {
	c := EventLogClientFactory{}
	c.Factory = *cbuild.NewFactory()

	nullClientDescriptor := cref.NewDescriptor("service-eventlog", "client", "null", "*", "1.0")
	httpClientDescriptor := cref.NewDescriptor("service-eventlog", "client", "commandable-http", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-eventlog", "client", "mock", "*", "1.0")

	c.RegisterType(nullClientDescriptor, version1.NewEventLogNullClientV1)
	c.RegisterType(httpClientDescriptor, version1.NewEventLogCommandableHttpClientV1)
	c.RegisterType(mockClientDescriptor, version1.NewEventLogMockClientV1)
	return &c
}
