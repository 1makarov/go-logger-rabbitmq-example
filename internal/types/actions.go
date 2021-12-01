package types

import "time"

type LogItem struct {
	Entity    string    `bson:"entity" json:"entity"`
	Action    string    `bson:"action" json:"action"`
	UserID    int64     `bson:"user_id" json:"user_id"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}
