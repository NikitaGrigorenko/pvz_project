//go:build integration

package tests

import (
	"bytes"
	"context"
	"devtask/internal/app/event"
	"devtask/internal/app/handlers"
	"devtask/internal/config"
	"devtask/internal/infrastructure/cache/in_memory"
	"devtask/internal/model"
	"devtask/internal/service/pvz"
	"devtask/internal/storage/postgres"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSender(t *testing.T) {
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

		configData, err := config.Read("test_config.json")

		//act
		handler.ServeHTTP(w, req)

		producer, err := event.NewProducer(configData.Brokers)

		sender, err := event.NewKafkaSender(producer, "events")
		err = sender.SendEvent(model.EventMessage{
			Method:    "POST",
			Body:      "",
			TimeStamp: time.Now(),
		})

		//assert
		assert.NoError(t, err)
	})
	t.Run("error brokers", func(t *testing.T) {
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

		_, err = event.NewProducer([]string{"321321"})

		//assert
		assert.Contains(t, err.Error(), "error with sync kafka-producer")
	})
}
