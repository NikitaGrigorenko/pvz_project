//go:build integration

package tests

import (
	"context"
	"devtask/internal/app/handlers"
	"devtask/internal/infrastructure/cache/in_memory"
	"devtask/internal/model"
	"devtask/internal/service/pvz"
	"devtask/internal/storage/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListHandle(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := NewFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	defer db.DB.GetPool(ctx).Close()

	pvzsRepo := postgres.NewPVZs(db.DB)

	cacheInMemory := in_memory.NewInMemoryCache[model.PVZ]()

	pvzService := pvz.NewService(pvzsRepo, cacheInMemory, db.DB)

	t.Run("smoke test", func(t *testing.T) {
		// arrange
		err := db.SetUp(t, "pvz")
		if err != nil {
			t.Fatal(err)
		}
		defer db.TearDown()

		_, err = pvzsRepo.AddPVZ(ctx, &model.PVZ{
			Name:    "test1",
			Address: "test1",
			Contact: "test1",
		})
		if err != nil {
			return
		}

		expectedBody := `[{"ID":1,"name":"test1","address":"test1","contact":"test1"}]`

		handler := handlers.List(pvzService)
		req := httptest.NewRequest("GET", "/pvz/", nil)

		w := httptest.NewRecorder()

		//act
		handler.ServeHTTP(w, req)

		//assert
		require.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, w.Body.String(), expectedBody)
	})
	t.Run("not found", func(t *testing.T) {
		// arrange
		err := db.SetUp(t, "pvz")
		if err != nil {
			t.Fatal(err)
		}
		defer db.TearDown()

		handler := handlers.List(pvzService)
		req := httptest.NewRequest("GET", "/pvz/", nil)

		w := httptest.NewRecorder()

		//act
		handler.ServeHTTP(w, req)

		//assert
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
