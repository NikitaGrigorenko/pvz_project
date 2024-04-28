package v1

import (
	"context"
	"devtask/internal/model"
	"devtask/internal/pkg/pb/devtask/pvz/v1"
	"go.opentelemetry.io/otel/trace"
)

type StoragePVZ interface {
	GetInfo(ctx context.Context, id int64) (*model.PVZ, error)
	AddInfo(ctx context.Context, pvz model.PVZ) (int64, error)
	UpdateInfo(ctx context.Context, pvz *model.PVZ, id int64) (int64, error)
	DeleteInfo(ctx context.Context, id int64) error
	ListInfo(ctx context.Context) ([]model.PVZ, error)
}

type Service struct {
	devtask_pvz_v1.UnimplementedPVZServer
	Storage StoragePVZ
	Tracer  trace.Tracer
}

func NewGrpcService(storage StoragePVZ, tracer trace.Tracer) *Service {
	return &Service{Storage: storage, Tracer: tracer}
}
