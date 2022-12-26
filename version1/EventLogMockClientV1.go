package version1

import (
	"context"
	"strings"
	"time"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type EventLogMockClientV1 struct {
	events []*SystemEventV1
}

func NewEventLogMockClientV1() *EventLogMockClientV1 {
	return &EventLogMockClientV1{
		events: make([]*SystemEventV1, 0),
	}
}

func (c *EventLogMockClientV1) GetEvents(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page cdata.DataPage[*SystemEventV1], err error) {
	items := make([]*SystemEventV1, 0)
	filterFunc := c.composeFilter(filter)

	for _, v := range c.events {
		item := v
		if filterFunc(*item) {
			items = append(items, item)
		}
	}
	return *cdata.NewDataPage(items, len(c.events)), nil
}

func (c *EventLogMockClientV1) LogEvent(ctx context.Context, correlationId string, event *SystemEventV1) error {
	if event.Severity < 0 {
		event.Severity = Informational
	}

	if event.Time.UnixMilli() == (time.Time{}).UnixMilli() {
		event.Time = time.Now()
	}

	c.events = append(c.events, event)

	return nil
}

func (c *EventLogMockClientV1) composeFilter(filter *cdata.FilterParams) func(item SystemEventV1) bool {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	id := filter.GetAsString("id")
	typee := filter.GetAsString("type")
	source := filter.GetAsString("source")
	correlationId := filter.GetAsString("correlation_id")
	fromCreateTime, fromCreateTimeOK := filter.GetAsNullableDateTime("from_create_time")
	toCreateTime, toCreateTimeOk := filter.GetAsNullableDateTime("to_create_time")
	severity, severityOk := filter.GetAsNullableInteger("severity")

	ids := make([]string, 0)

	if idsStr := filter.GetAsString("ids"); idsStr != "" {
		ids = strings.Split(idsStr, ",")
	}

	return func(item SystemEventV1) bool {
		if id != "" && id != item.Id {
			return false
		}
		if typee != "" && typee != item.Type {
			return false
		}
		if source != "" && source != item.Source {
			return false
		}
		if correlationId != "" && correlationId != item.CorrelationId {
			return false
		}
		if severityOk && severity != item.Severity {
			return false
		}
		if fromCreateTimeOK && item.Time.Nanosecond() >= fromCreateTime.Nanosecond() {
			return false
		}
		if toCreateTimeOk && item.Time.Nanosecond() < toCreateTime.Nanosecond() {
			return false
		}
		if len(ids) > 0 && !contains(ids, item.Id) {
			return false
		}
		return true
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
