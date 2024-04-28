package v1

import (
	"context"
	devtask_pvz_v1 "devtask/internal/pkg/pb/devtask/pvz/v1"
)

func (s *Service) GetInfo(ctx context.Context, req *devtask_pvz_v1.GetInfoRequest) (*devtask_pvz_v1.GetInfoResponse, error) {
	ctx, span := s.Tracer.Start(
		ctx,
		"GetInfo")

	span.AddEvent("GetInfo")
	defer span.End()
	pvzInfo, err := s.Storage.GetInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &devtask_pvz_v1.GetInfoResponse{Pvz: &devtask_pvz_v1.PVZInfo{
		Id:      pvzInfo.ID,
		Name:    pvzInfo.Name,
		Address: pvzInfo.Address,
		Contact: pvzInfo.Contact,
	}}, nil
}
