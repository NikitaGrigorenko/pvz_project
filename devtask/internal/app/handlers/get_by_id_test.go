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

func Test_GetByID(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = int64(1)
	)
	t.Run("smoke test", func(t *testing.T) {
		t.Parallel()
		// arrange
		s := setUp(t)
		defer s.tearDown()
		s.mockArticles.EXPECT().GetInfo(gomock.Any(), id).Return(&model.PVZ{ID: 1,
			Name:    "test",
			Address: "test",
			Contact: "+7900000000"}, nil)

		// act
		result, status := Get(s.srv, ctx, id)

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
			s.mockArticles.EXPECT().GetInfo(gomock.Any(), id).Return(&model.PVZ{}, model.ErrObjectNotFound)

			// act
			_, status := Get(s.srv, ctx, id)

			//assert
			require.Equal(t, http.StatusNotFound, status)
			assert.Equal(t, nil, nil)
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()
			s.mockArticles.EXPECT().GetInfo(gomock.Any(), id).Return(&model.PVZ{}, model.ErrNoRowsInResultSet)

			// act
			_, status := Get(s.srv, ctx, id)

			//assert
			require.Equal(t, http.StatusInternalServerError, status)
			assert.Equal(t, nil, nil)
		})
	})
}

func Test_validateGetByID(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		result := ValidateGetByID(1)
		assert.True(t, result)
	})
	t.Run("fail", func(t *testing.T) {
		result := ValidateGetByID(-1)
		assert.False(t, result)
	})
}
