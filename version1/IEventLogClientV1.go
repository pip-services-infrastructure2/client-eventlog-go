package version1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IEventLogClientV1 interface {
	GetEvents(ctx context.Context, correlationId string, filter *cdata.FilterParams,
		paging *cdata.PagingParams) (page cdata.DataPage[*SystemEventV1], err error)

	LogEvent(ctx context.Context, correlationId string, event *SystemEventV1) error
}
