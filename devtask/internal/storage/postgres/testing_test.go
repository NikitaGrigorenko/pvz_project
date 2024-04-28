package postgres

import (
	"devtask/internal/service/pvz"
	mock_pvz "devtask/internal/storage/postgres/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

type pvzsRepoFixtures struct {
	ctrl   *gomock.Controller
	repo   pvz.StoragePVZ
	mockDB *mock_pvz.MockDBops
}

func setUp(t *testing.T) pvzsRepoFixtures {
	ctrl := gomock.NewController(t)
	mockDB := mock_pvz.NewMockDBops(ctrl)
	repo := NewPVZs(mockDB)
	return pvzsRepoFixtures{
		ctrl:   ctrl,
		repo:   repo,
		mockDB: mockDB,
	}
}

func (a *pvzsRepoFixtures) tearDown() {
	a.ctrl.Finish()
}
