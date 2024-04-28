package handlers

import (
	"context"
	"devtask/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Delete(t *testing.T) {
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
		s.mockArticles.EXPECT().DeleteInfo(gomock.Any(), id).Return(nil)

		// act
		status := Del(s.srv, ctx, id)

		//assert
		assert.Equal(t, http.StatusOK, status)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()
			s.mockArticles.EXPECT().DeleteInfo(gomock.Any(), id).Return(model.ErrObjectNotFound)

			// act
			status := Del(s.srv, ctx, id)

			//assert
			assert.Equal(t, http.StatusNotFound, status)
		})
		t.Run("not found 2", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()
			s.mockArticles.EXPECT().DeleteInfo(gomock.Any(), id).Return(model.ErrNoRowsInResultSet)

			// act
			status := Del(s.srv, ctx, id)

			//assert
			assert.Equal(t, http.StatusNotFound, status)
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()
			s.mockArticles.EXPECT().DeleteInfo(gomock.Any(), id).Return(assert.AnError)

			// act
			status := Del(s.srv, ctx, id)

			//assert
			assert.Equal(t, http.StatusInternalServerError, status)
		})
	})
}

func Test_validateDelete(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		result := ValidateDelete(1)
		assert.True(t, result)
	})
	t.Run("fail", func(t *testing.T) {
		result := ValidateDelete(-1)
		assert.False(t, result)
	})
}
