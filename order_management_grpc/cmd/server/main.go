package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "order_management_grpc/proto/orderpb/proto"

	"order_management_grpc/internal/service"
	"order_management_grpc/internal/storage"
)

const (
	port = ":50051"
)

func getPostgresConn() (*pgx.Conn, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("missing database environment variables")
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	return pgx.Connect(context.Background(), connStr)
}

func main() {
	_ = godotenv.Load()

	var store storage.OrderStore
	conn, err := getPostgresConn()
	if err == nil {
		log.Println("Using PostgreSQL for order storage.")
		store = storage.NewPostgresOrderStore(conn)
	} else {
		log.Printf("PostgreSQL not available (%v), using in-memory store.", err)
		store = storage.NewInMemoryOrderStore()
	}

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
