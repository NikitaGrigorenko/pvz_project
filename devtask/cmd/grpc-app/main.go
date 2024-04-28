package main

import (
	"context"
	"devtask/internal/app/grpc_handlers/devtask/pvz/v1"
	"devtask/internal/config"
	"devtask/internal/infrastructure/cache/in_memory"
	"devtask/internal/model"
	"devtask/internal/pkg/db"
	devtask_pvz_v1 "devtask/internal/pkg/pb/devtask/pvz/v1"
	"devtask/internal/service/pvz"
	"devtask/internal/storage/postgres"
	"devtask/internal/tracing"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
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

	inMemoryCache := in_memory.NewInMemoryCache[model.PVZ]()

	pvzService := pvz.NewService(pvzsRepo, inMemoryCache, database)

	err = tracing.NewTracer(ctx, configData.AddrInfo.AddrJaeger)
	if err != nil {
		log.Fatal("error while creating new tracer", err)
	}

	tracer := otel.Tracer("server")

	grpcService := v1.NewGrpcService(pvzService, tracer)

	grpcRun(grpcService)

	fmt.Println("Работа сервера завершена!")
}

func grpcRun(service *v1.Service) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 10000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	devtask_pvz_v1.RegisterPVZServer(grpcServer, service)

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-quit
	fmt.Println("Завершение работы сервера!")
	grpcServer.GracefulStop()
}
