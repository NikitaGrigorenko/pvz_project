//go:generate mockgen -source ./common_interfaces.go -destination=./common_interfaces_mocks_test.go -package=handlers

package handlers

import (
	"context"
	"devtask/internal/model"
)

const QueryParamKey = "key"

type StoragePVZ interface {
	GetInfo(ctx context.Context, id int64) (*model.PVZ, error)
	AddInfo(ctx context.Context, pvz model.PVZ) (int64, error)
	UpdateInfo(ctx context.Context, pvz *model.PVZ, id int64) (int64, error)
	DeleteInfo(ctx context.Context, id int64) error
	ListInfo(ctx context.Context) ([]model.PVZ, error)
}
