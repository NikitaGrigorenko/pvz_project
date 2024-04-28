//go:build integration

package tests

import (
	"bytes"
	"context"
	"devtask/internal/app"
	"devtask/internal/config"
	redis2 "devtask/internal/infrastructure/cache/redis"
	"devtask/internal/metrics"
	"devtask/internal/model"
	"devtask/internal/service/pvz"
	"devtask/internal/storage/postgres"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateHandle(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := NewFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	defer db.DB.GetPool(ctx).Close()

	pvzsRepo := postgres.NewPVZs(db.DB)

	configData, err := config.Read("test_config.json")
	redisCache := redis2.NewRedis[model.PVZ](&redis.Options{
		Addr:     configData.RedisInfo.Addr,
		Password: configData.RedisInfo.Password,
		DB:       configData.RedisInfo.DB,
	})

	pvzService := pvz.NewService(pvzsRepo, redisCache, db.DB)
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

		pvzTest := `{"name":"test1","address":"test1","contact":"test1"}`

		req := httptest.NewRequest("PUT", "/pvz/1", bytes.NewBufferString(pvzTest))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		//act
		router.ServeHTTP(w, req)

		//assert
		require.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("not found", func(t *testing.T) {
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

		pvz1 := &model.PVZ{Name: "test2", Address: "test2", Contact: "name2"}
		body, _ := json.Marshal(pvz1)

		req := httptest.NewRequest("PUT", "/pvz/2", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		//act
		router.ServeHTTP(w, req)

		//assert
		require.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("bad request", func(t *testing.T) {
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

		body, _ := json.Marshal(1)

		req := httptest.NewRequest("PUT", "/pvz/2", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		//act
		router.ServeHTTP(w, req)

		//assert
		require.Equal(t, http.StatusBadRequest, w.Code)
	})
}
