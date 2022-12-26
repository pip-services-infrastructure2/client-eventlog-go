package version1

import (
	"context"
	"os"
	"time"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	clients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type EventLogCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewEventLogCommandableHttpClientV1() *EventLogCommandableHttpClientV1 {
	return &EventLogCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/eventlog"),
	}
}

func (c *EventLogCommandableHttpClientV1) GetEvents(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (cdata.DataPage[*SystemEventV1], error) {
	params := cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_events", correlationId, params)
	if err != nil {
		return *cdata.NewEmptyDataPage[*SystemEventV1](), err
	}

	return clients.HandleHttpResponse[cdata.DataPage[*SystemEventV1]](res, correlationId)
}

func (c *EventLogCommandableHttpClientV1) LogEvent(ctx context.Context, correlationId string, event *SystemEventV1) error {
	event.Time = time.Now()
	if event.Source == "" {
		event.Source, _ = os.Hostname()
	}

	params := cdata.NewAnyValueMapFromTuples(
		"event", event,
	)

	_, err := c.CallCommand(ctx, "log_event", correlationId, params)
	return err
}
