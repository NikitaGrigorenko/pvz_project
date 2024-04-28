package event

import (
	mock_service "devtask/internal/app/event/mocks"
	"devtask/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_SendEvent(t *testing.T) {
	t.Parallel()
	t.Run("smoke test", func(t *testing.T) {
		t.Parallel()
		// arrange
		ctrl := gomock.NewController(t)
		mockProd := mock_service.NewMockProducer(ctrl)
		s, _ := NewKafkaSender(mockProd, "events")

		mockProd.EXPECT().SendSyncMessage(gomock.Any()).Return(int32(1), int64(1), nil)

		// act
		err := s.SendEvent(model.EventMessage{
			Method:    "GET",
			Body:      "",
			TimeStamp: time.Now(),
		})

		//assert
		assert.NoError(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		// arrange
		ctrl := gomock.NewController(t)
		mockProd := mock_service.NewMockProducer(ctrl)
		s, _ := NewKafkaSender(mockProd, "events")

		mockProd.EXPECT().SendSyncMessage(gomock.Any()).Return(int32(1), int64(1), assert.AnError)

		// act
		err := s.SendEvent(model.EventMessage{
			Method:    "GET",
			Body:      "",
			TimeStamp: time.Now(),
		})

		//assert
		assert.Equal(t, "send sync message error", err.Error())
	})
}
