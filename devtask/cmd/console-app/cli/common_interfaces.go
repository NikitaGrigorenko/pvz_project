package cli

import (
	"context"
	"devtask/internal/model"
)

type storagePVZ interface {
	GetInfo(ctx context.Context, id int64) (*model.PVZ, error)
	AddInfo(ctx context.Context, pvz model.PVZ) (int64, error)
}
