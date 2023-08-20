package audit

import (
	"time"
)

type LogItem struct {
	Entity    string    `json:"entity" bson:"entity"`
	Action    string    `json:"action" bson:"action"`
	EntityID  int64     `json:"entity_id" bson:"entity_id"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
