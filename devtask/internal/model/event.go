package model

import "time"

type EventMessage struct {
	Method    string
	Body      string
	TimeStamp time.Time
}
