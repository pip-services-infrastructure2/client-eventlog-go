package version1

import (
	"time"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type SystemEventV1 struct {
	Id            string               `json:"id" bson:"_id"`
	Time          time.Time            `json:"time" bson:"time"`
	CorrelationId string               `json:"correlation_id" bson:"correlation_id"`
	Source        string               `json:"source" bson:"source"`
	Type          string               `json:"type" bson:"type"`
	Severity      int                  `json:"severity" bson:"severity"`
	Message       string               `json:"message" bson:"message"`
	Details       cdata.StringValueMap `json:"details" bson:"details"`
}
