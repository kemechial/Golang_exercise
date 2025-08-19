package tests

import (
	"context"
	"testing"

	"order_management_grpc/internal/service"
	"order_management_grpc/internal/storage"
	pb "order_management_grpc/proto/orderpb"
)

func TestOrderService_Integration(t *testing.T) {
	store := storage.NewInMemoryOrderStore()
	orderSvc := service.NewOrderService(store)

	// Create order
	items := []*pb.OrderItem{{ProductId: "p1", ProductName: "Item1", Quantity: 1, UnitPrice: 100.0}}
	createResp, err := orderSvc.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		CustomerId:   "c2",
		CustomerName: "Integration Tester",
		Items:        items,
	})
	if err != nil {
		t.Fatalf("CreateOrder failed: %v", err)
	}
	orderID := createResp.Order.Id

	// Get order
	getResp, err := orderSvc.GetOrder(context.Background(), &pb.GetOrderRequest{OrderId: orderID})
	if err != nil {
		t.Fatalf("GetOrder failed: %v", err)
	}
	if getResp.Order.CustomerName != "Integration Tester" {
		t.Errorf("expected customer name Integration Tester, got %s", getResp.Order.CustomerName)
	}

	// Update status
	_, err = orderSvc.UpdateOrderStatus(context.Background(), &pb.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  pb.OrderStatus_CONFIRMED,
	})
	if err != nil {
		t.Fatalf("UpdateOrderStatus failed: %v", err)
	}

	// List orders
	listResp, err := orderSvc.ListOrders(context.Background(), &pb.ListOrdersRequest{CustomerId: "c2", PageSize: 10})
	if err != nil {
		t.Fatalf("ListOrders failed: %v", err)
	}
	if len(listResp.Orders) != 1 {
		t.Errorf("expected 1 order, got %d", len(listResp.Orders))
	}

	// Cancel order
	_, err = orderSvc.CancelOrder(context.Background(), &pb.CancelOrderRequest{OrderId: orderID})
	if err != nil {
		t.Fatalf("CancelOrder failed: %v", err)
	}
}
