package main

import (
	"context"
	"devtask/internal/app"
	"devtask/internal/app/event"
	"devtask/internal/config"
	rediscache "devtask/internal/infrastructure/cache/redis"
	"devtask/internal/infrastructure/kafka"
	"devtask/internal/metrics"
	"devtask/internal/model"
	"devtask/internal/pkg/db"
	"devtask/internal/service/pvz"
	"devtask/internal/storage/postgres"
	"devtask/internal/tracing"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.NewDb(ctx, "config/config.json")
	if err != nil {
		log.Fatal("error while creating db", err)
	}

	defer database.GetPool(ctx).Close()

	configData, err := config.Read("config/config.json")
	if err != nil {
		log.Fatal("error while config reading", err)
	}
	pvzsRepo := postgres.NewPVZs(database)

	//inMemoryCache := in_memory.NewInMemoryCache[model.PVZ]()

	redisCache := rediscache.NewRedis[model.PVZ](&redis.Options{
		Addr:     configData.RedisInfo.Addr,
		Password: configData.RedisInfo.Password,
		DB:       configData.RedisInfo.DB,
	})

	pvzService := pvz.NewService(pvzsRepo, redisCache, database)

	producer, err := event.NewProducer(configData.Brokers)
	if err != nil {
		log.Fatal("error while creating new kafka producer", err)
	}

	sender, err := event.NewKafkaSender(producer, configData.Topic)
	if err != nil {
		log.Fatal("error while creating new kafka sender", err)
	}

	recv, err := kafka.NewConsumer(configData.Brokers, configData.Group)
	if err != nil {
		log.Fatal("error while creating new kafka consumer", err)
	}

	promMetrics := metrics.RegisterMetrics()

	err = tracing.NewTracer(ctx, configData.AddrInfo.AddrJaeger)
	if err != nil {
		log.Fatal("error while creating new tracer", err)
	}

	app.RunHTTP(ctx, pvzService, configData, sender, recv, promMetrics)
}
