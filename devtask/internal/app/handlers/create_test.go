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

func Test_Create(t *testing.T) {
	t.Parallel()
	var (
		ctx      = context.Background()
		pvzModel = model.PVZ{
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
		s.mockArticles.EXPECT().AddInfo(gomock.Any(), pvzModel).Return(int64(1), nil)

		// act
		result, status := Add(s.srv, ctx, pvzRequestModel)

		//assert
		require.Equal(t, http.StatusOK, status)
		assert.Equal(t, "{\"ID\":1,\"name\":\"test\",\"address\":\"test\",\"contact\":\"+7900000000\"}", string(result))

	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()
			s.mockArticles.EXPECT().AddInfo(gomock.Any(), pvzModel).Return(int64(0), model.ErrNoRowsInResultSet)

			// act
			result, status := Add(s.srv, ctx, pvzRequestModel)

			//assert
			require.Equal(t, http.StatusInternalServerError, status)
			assert.Equal(t, "", string(result))
		})
	})
}
