package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/order-management/internal/service"
	"github.com/order-management/internal/storage"
	pb "github.com/order-management/proto"
)

const (
	port = ":50051"
)

func main() {
	// Create storage
	store := storage.NewInMemoryOrderStore()

	// Create service
	orderService := service.NewOrderService(store)

	// Create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server
	s := grpc.NewServer()

	// Register service
	pb.RegisterOrderServiceServer(s, orderService)

	// Register reflection service (for grpcurl and debugging)
	reflection.Register(s)

	log.Printf("Order Management gRPC server listening on port %s", port)
	log.Printf("Use grpcurl for testing: grpcurl -plaintext localhost%s list", port)

	// Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
