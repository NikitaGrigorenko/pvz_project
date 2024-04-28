package clients

import (
	devtask_pvz_v1 "devtask/internal/pkg/pb/devtask/pvz/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewGrpcClient(addr string) devtask_pvz_v1.PVZClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	client := devtask_pvz_v1.NewPVZClient(conn)
	return client
}
