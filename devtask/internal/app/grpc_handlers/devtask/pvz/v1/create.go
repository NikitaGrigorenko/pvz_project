package v1

import (
	"context"
	"devtask/internal/model"
	devtask_pvz_v1 "devtask/internal/pkg/pb/devtask/pvz/v1"
)

func (s *Service) AddInfo(ctx context.Context, req *devtask_pvz_v1.AddInfoRequest) (*devtask_pvz_v1.AddInfoResponse, error) {
	ctx, span := s.Tracer.Start(
		ctx,
		"AddInfo")

	span.AddEvent("AddInfo")
	defer span.End()

	pvzRepo := &model.PVZ{
		Name:    req.Pvz.Name,
		Address: req.Pvz.Address,
		Contact: req.Pvz.Contact,
	}
	id, err := s.Storage.AddInfo(ctx, *pvzRepo)
	if err != nil {
		return nil, err
	}

	return &devtask_pvz_v1.AddInfoResponse{Pvz: &devtask_pvz_v1.PVZInfo{
		Id:      id,
		Name:    pvzRepo.Name,
		Address: pvzRepo.Address,
		Contact: pvzRepo.Contact,
	}}, nil
}
