package model

import (
	"time"
)

type PackageType int

const (
	Packet PackageType = iota
	Box
	Tape
)

// Order struct using in ListRefunds and ListOrders
type Order struct {
	IdOrder     int64
	IdClient    int64
	IsGiven     bool
	IsReturned  bool
	IsRefunded  bool
	ExpireDate  time.Time
	WeightGrams int64
	PriceKopeck int64
	PackageType PackageType
}

// OrderInput struct using only in TakeOrder to give an access only for fields which are needed
type OrderInput struct {
	IdOrder     int64
	IdClient    int64
	ExpireDate  time.Time
	WeightGrams int64
	PriceKopeck int64
	PackageType PackageType
}
