//go:generate mockgen -source ./service.go -destination=./mocks/service.go -package=mock_service
package pvz

import (
	"context"
	"devtask/internal/model"
	"github.com/jackc/pgx/v4"
)

type StoragePVZ interface {
	AddPVZ(ctx context.Context, pvz *model.PVZ) (int64, error)
	GetPVZ(ctx context.Context, id int64) (*model.PVZ, error)
	Update(ctx context.Context, pvz *model.PVZ, id int64) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]model.PVZ, error)
}

type Transactor interface {
	InTransaction(ctx context.Context, pgx pgx.TxAccessMode, f func(ctxTX context.Context) error) error
}

type Cache[T any] interface {
	Get(ctx context.Context, id int64) (T, error)
	Set(ctx context.Context, id int64, item T) error
	Delete(ctx context.Context, id int64)
}

type Service struct {
	storage    StoragePVZ
	cache      Cache[model.PVZ]
	transactor Transactor
}

func NewService(storage StoragePVZ, cache Cache[model.PVZ], transactor Transactor) *Service {
	return &Service{
		storage:    storage,
		cache:      cache,
		transactor: transactor,
	}
}

func (s Service) GetInfo(ctx context.Context, id int64) (*model.PVZ, error) {
	pvzInfo, err := s.cache.Get(ctx, id)
	if err == nil {
		return &pvzInfo, nil
	}

	pvz, err := s.storage.GetPVZ(ctx, id)

	if err != nil {
		return nil, err
	}

	if err := s.cache.Set(ctx, pvz.ID, *pvz); err != nil {
		return nil, err
	}

	return pvz, err
}

func (s Service) AddInfo(ctx context.Context, pvz model.PVZ) (id int64, err error) {
	err = s.transactor.InTransaction(ctx, pgx.ReadWrite, func(ctxTX context.Context) error {
		id, err = s.storage.AddPVZ(ctx, &pvz)
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (s Service) UpdateInfo(ctx context.Context, pvz *model.PVZ, id int64) (idRet int64, err error) {
	err = s.transactor.InTransaction(ctx, pgx.ReadOnly, func(ctxTX context.Context) error {
		idRet, err = s.storage.Update(ctx, pvz, id)
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (s Service) DeleteInfo(ctx context.Context, id int64) (err error) {
	err = s.transactor.InTransaction(ctx, pgx.ReadOnly, func(ctxTX context.Context) error {
		err := s.storage.Delete(ctx, id)
		if err != nil {
			s.cache.Delete(ctx, id)
			return err
		}
		return nil
	})
	return err
}

func (s Service) ListInfo(ctx context.Context) ([]model.PVZ, error) {
	pvzs, err := s.storage.List(ctx)
	return pvzs, err
}
