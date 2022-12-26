package version1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type EventLogNullClientV1 struct {
}

func NewEventLogNullClientV1() *EventLogNullClientV1 {
	return &EventLogNullClientV1{}
}

func (c *EventLogNullClientV1) GetEvents(ctx context.Context, correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (cdata.DataPage[*SystemEventV1], error) {
	return *cdata.NewEmptyDataPage[*SystemEventV1](), nil
}

func (c *EventLogNullClientV1) LogEvent(ctx context.Context, correlationId string, event *SystemEventV1) error {
	return nil
}
