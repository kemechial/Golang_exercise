package tests

import (
	"context"
	"testing"
	"time"

	"order_management_grpc/internal/models"
	"order_management_grpc/internal/service"
	"order_management_grpc/internal/storage"
	pb "order_management_grpc/proto/orderpb/proto"
)

func TestCreateOrder_Valid(t *testing.T) {
	store := storage.NewInMemoryOrderStore()
	orderSvc := service.NewOrderService(store)

	items := []*pb.OrderItem{
		{ProductId: "p1", ProductName: "Item1", Quantity: 2, UnitPrice: 10.0},
		{ProductId: "p2", ProductName: "Item2", Quantity: 1, UnitPrice: 20.0},
	}
	resp, err := orderSvc.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		CustomerId:   "c1",
		CustomerName: "Test Customer",
		Items:        items,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.Order.TotalAmount != 40.0 {
		t.Errorf("expected total 40.0, got %v", resp.Order.TotalAmount)
	}
}

func TestCreateOrder_Invalid(t *testing.T) {
	store := storage.NewInMemoryOrderStore()
	orderSvc := service.NewOrderService(store)

	// Missing customer ID
	_, err := orderSvc.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		CustomerName: "Test Customer",
		Items:        []*pb.OrderItem{{ProductId: "p1", ProductName: "Item1", Quantity: 1, UnitPrice: 10.0}},
	})
	if err == nil {
		t.Error("expected error for missing customer ID")
	}
}

func TestOrderStatusTransition(t *testing.T) {
	order := &models.Order{
		ID:           "o1",
		CustomerID:   "c1",
		CustomerName: "Test Customer",
		Status:       models.StatusPending,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if !order.IsValidStatusTransition(models.StatusConfirmed) {
		t.Error("should allow transition from PENDING to CONFIRMED")
	}
	if order.IsValidStatusTransition(models.StatusDelivered) {
		t.Error("should not allow transition from PENDING to DELIVERED")
	}
}
