package handlers

import (
	"context"
	"devtask/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_Update(t *testing.T) {
	t.Parallel()
	var (
		ctx      = context.Background()
		id       = int64(1)
		pvzModel = &model.PVZ{
			Name:    "test",
			Address: "test",
			Contact: "+7900000000"}
		pvzRequestModel = model.PVZRequest{
			Name:    "test",
			Address: "test",
			Contact: "+7900000000"}
	)
	t.Run("smoke test", func(t *testing.T) {
		t.Parallel()
		// arrange
		s := setUp(t)
		defer s.tearDown()
		s.mockArticles.EXPECT().UpdateInfo(gomock.Any(), pvzModel, id).Return(int64(1), nil)

		// act
		result, status := Upd(s.srv, ctx, pvzRequestModel, id)

		//assert
		require.Equal(t, http.StatusOK, status)
		assert.Equal(t, "{\"ID\":1,\"name\":\"test\",\"address\":\"test\",\"contact\":\"+7900000000\"}", string(result))

	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()
			s.mockArticles.EXPECT().UpdateInfo(gomock.Any(), pvzModel, id).Return(int64(0), model.ErrNoRowsInResultSet)

			// act
			result, status := Upd(s.srv, ctx, pvzRequestModel, id)

			//assert
			require.Equal(t, http.StatusNotFound, status)
			assert.Equal(t, "", string(result))
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()
			s.mockArticles.EXPECT().UpdateInfo(gomock.Any(), pvzModel, id).Return(int64(0), assert.AnError)

			// act
			result, status := Upd(s.srv, ctx, pvzRequestModel, id)

			//assert
			require.Equal(t, http.StatusInternalServerError, status)
			assert.Equal(t, "", string(result))
		})
	})
}
