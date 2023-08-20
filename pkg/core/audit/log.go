package audit

import (
	"time"
)

const (
	ENTITY_USER      = "USER"
	ENTITY_PETITION  = "PETITION"
	ENTITY_SIGNATURE = "SIGNATURE"

	ACTION_CREATE   = "CREATE"
	ACTION_UPDATE   = "UPDATE"
	ACTION_GET      = "GET"
	ACTION_DELETE   = "DELETE"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
	ACTION_LOGOUT   = "LOGOUT"
)

type LogItem struct {
	Entity    string    `json:"entity" bson:"entity"`
	Action    string    `json:"action" bson:"action"`
	EntityID  int64     `json:"entity_id" bson:"entity_id"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
