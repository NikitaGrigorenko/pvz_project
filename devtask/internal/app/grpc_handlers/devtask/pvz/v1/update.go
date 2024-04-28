package v1

import (
	"context"
	"devtask/internal/model"
	devtask_pvz_v1 "devtask/internal/pkg/pb/devtask/pvz/v1"
)

func (s *Service) UpdateInfo(ctx context.Context, req *devtask_pvz_v1.UpdateInfoRequest) (*devtask_pvz_v1.UpdateInfoResponse, error) {
	pvzRepo := &model.PVZ{
		Name:    req.Pvz.Name,
		Address: req.Pvz.Address,
		Contact: req.Pvz.Contact,
	}
	id, err := s.Storage.UpdateInfo(ctx, pvzRepo, req.Pvz.Id)
	if err != nil {
		return nil, err
	}
	return &devtask_pvz_v1.UpdateInfoResponse{Pvz: &devtask_pvz_v1.PVZInfo{
		Id:      id,
		Name:    pvzRepo.Name,
		Address: pvzRepo.Address,
		Contact: pvzRepo.Contact,
	}}, nil
}
