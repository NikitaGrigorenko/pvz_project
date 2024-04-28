//go:build integration

package tests

import (
	"context"
	"devtask/internal/app"
	"devtask/internal/infrastructure/cache/in_memory"
	"devtask/internal/metrics"
	"devtask/internal/model"
	"devtask/internal/service/pvz"
	"devtask/internal/storage/postgres"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteHandle(t *testing.T) {
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
	promMetrics := metrics.RegisterMetrics()
	router := app.CreateRouter(pvzService, promMetrics)

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

		req := httptest.NewRequest("DELETE", fmt.Sprintf("/pvz/1"), nil)
		w := httptest.NewRecorder()

		//act
		router.ServeHTTP(w, req)

		//assert
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("not found", func(t *testing.T) {
		// arrange
		err := db.SetUp(t, "pvz")
		if err != nil {
			t.Fatal(err)
		}
		defer db.TearDown()

		req := httptest.NewRequest("DELETE", fmt.Sprintf("/pvz/1000"), nil)
		w := httptest.NewRecorder()

		//act
		router.ServeHTTP(w, req)

		//assert
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
