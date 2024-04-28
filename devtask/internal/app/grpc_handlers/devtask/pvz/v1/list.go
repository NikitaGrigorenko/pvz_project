package v1

import (
	"context"
	devtask_pvz_v1 "devtask/internal/pkg/pb/devtask/pvz/v1"
)

func (s *Service) ListInfo(ctx context.Context, _ *devtask_pvz_v1.ListInfoRequest) (*devtask_pvz_v1.ListInfoResponse, error) {
	pvzList, err := s.Storage.ListInfo(ctx)
	if err != nil {
		return nil, err
	}

	var pvzListPb []*devtask_pvz_v1.PVZInfo
	for _, pvz := range pvzList {
		pvzListPb = append(pvzListPb, &devtask_pvz_v1.PVZInfo{
			Id:      pvz.ID,
			Name:    pvz.Name,
			Address: pvz.Address,
			Contact: pvz.Contact,
		})
	}
	pvzListPb = append(pvzListPb, &devtask_pvz_v1.PVZInfo{})
	return &devtask_pvz_v1.ListInfoResponse{Pvz: pvzListPb}, nil
}
