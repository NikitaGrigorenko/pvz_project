package v1

import (
	"context"
	devtask_pvz_v1 "devtask/internal/pkg/pb/devtask/pvz/v1"
)

func (s *Service) DeleteInfo(ctx context.Context, req *devtask_pvz_v1.DeleteInfoRequest) (*devtask_pvz_v1.DeleteInfoResponse, error) {
	err := s.Storage.DeleteInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
