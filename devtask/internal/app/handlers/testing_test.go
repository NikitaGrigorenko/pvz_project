package handlers

import (
	"github.com/golang/mock/gomock"
	"testing"
)

type PvzsRepoFixtures struct {
	ctrl         *gomock.Controller
	srv          StoragePVZ
	mockArticles MockStoragePVZ
}

func setUp(t *testing.T) PvzsRepoFixtures {
	ctrl := gomock.NewController(t)
	mockArticles := NewMockStoragePVZ(ctrl)
	srv := mockArticles
	return PvzsRepoFixtures{
		ctrl:         ctrl,
		mockArticles: *mockArticles,
		srv:          srv,
	}
}

func (a *PvzsRepoFixtures) tearDown() {
	a.ctrl.Finish()
}
