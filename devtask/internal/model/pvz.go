package model

import (
	"errors"
)

type PVZ struct {
	ID      int64  `json:"ID"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Contact string `json:"contact"`
}

type PVZRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Contact string `json:"contact"`
}

var ErrObjectNotFound = errors.New("not found")
var ErrNoRowsInResultSet = errors.New("no rows in result set")
