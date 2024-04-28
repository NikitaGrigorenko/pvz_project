//go:build integration

package tests

import (
	"bytes"
	"context"
	"devtask/internal/app/handlers"
	"devtask/internal/infrastructure/cache/in_memory"
	"devtask/internal/model"
	"devtask/internal/service/pvz"
	"devtask/internal/storage/postgres"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHandle(t *testing.T) {
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

		pvzTest := `{"name":"test1","address":"test1","contact":"test1"}`

		handler := handlers.Create(pvzService)
		req := httptest.NewRequest("POST", "/pvz", bytes.NewBufferString(pvzTest))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		//act
		handler.ServeHTTP(w, req)

		//assert
		require.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("bad request", func(t *testing.T) {
		// arrange
		err := db.SetUp(t, "pvz")
		if err != nil {
			t.Fatal(err)
		}
		defer db.TearDown()

		body, _ := json.Marshal(1)

		handler := handlers.Create(pvzService)
		req := httptest.NewRequest("POST", "/pvz", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		//act
		handler.ServeHTTP(w, req)

		//assert
		require.Equal(t, http.StatusBadRequest, w.Code)
	})
}
