package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "order_management_grpc/proto/orderpb"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewOrderServiceClient(conn)

	// Simple check: try listing orders (should succeed even if empty)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = client.ListOrders(ctx, &pb.ListOrdersRequest{PageSize: 1})
	if err != nil {
		log.Fatalf("gRPC server health check failed: %v", err)
	}
	log.Println("gRPC server is reachable and OrderService is available.")

	// TODO: Implement production client logic or CLI here
	log.Println("OrderService gRPC client is ready. Implement CLI or client logic as needed.")
}
