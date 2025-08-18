package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"order_management_grpc/internal/service"
	"order_management_grpc/internal/storage"
	pb "order_management_grpc/proto/orderpb/proto"

	"google.golang.org/grpc/metadata"
)

func TestWatchOrderStatus_ServerStreaming(t *testing.T) {
	store := storage.NewInMemoryOrderStore()
	orderSvc := service.NewOrderService(store)

	green := "\033[32m"
	red := "\033[31m"
	reset := "\033[0m"
	check := "✅"
	fail := "❌"
	arrow := "➡️"
	pendingIcon := "🟡"
	confirmedIcon := "🟢"

	// Create order
	items := []*pb.OrderItem{{ProductId: "p1", ProductName: "Item1", Quantity: 1, UnitPrice: 50.0}}
	createResp, err := orderSvc.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		CustomerId:   "c3",
		CustomerName: "Stream Tester",
		Items:        items,
	})
	if err != nil {
		t.Fatalf("CreateOrder failed: %v", err)
	}
	orderID := createResp.Order.Id

	// Start watcher in goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream := &mockOrderStatusStream{updates: make(chan *pb.Order, 10), ctx: ctx}
	go func() {
		_ = orderSvc.WatchOrderStatus(&pb.GetOrderRequest{OrderId: orderID}, stream)
	}()

	// Wait for initial status
	time.Sleep(50 * time.Millisecond)
	var receivedUpdates []*pb.Order
	collect := func() {
		for {
			select {
			case update := <-stream.updates:
				receivedUpdates = append(receivedUpdates, update)
			default:
				return
			}
		}
	}
	collect()
	if len(receivedUpdates) == 0 {
		t.Errorf("%s%s no initial update received on stream%s", red, fail, reset)
	} else {
		statusIcon := pendingIcon
		if receivedUpdates[0].Status == pb.OrderStatus_CONFIRMED {
			statusIcon = confirmedIcon
		}
		t.Logf("%s%s Initial status: %v %s%s", green, check, receivedUpdates[0].Status, statusIcon, reset)
	}

	// Update status
	_, err = orderSvc.UpdateOrderStatus(context.Background(), &pb.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  pb.OrderStatus_CONFIRMED,
	})
	if err != nil {
		t.Fatalf("UpdateOrderStatus failed: %v", err)
	}

	// Wait for update
	time.Sleep(100 * time.Millisecond)
	receivedUpdates = receivedUpdates[:0]
	var summaryRows []string
	collect()
	if len(receivedUpdates) == 0 {
		t.Errorf("%s%s no update received after status change%s", red, fail, reset)
	} else {
		for i, update := range receivedUpdates {
			statusIcon := pendingIcon
			if update.Status == pb.OrderStatus_CONFIRMED {
				statusIcon = confirmedIcon
			}
			t.Logf("%s%s Update %d: status=%v %s%s", green, arrow, i, update.Status, statusIcon, reset)
			summaryRows = append(summaryRows, fmt.Sprintf("| %d | %v | %s |", i, update.Status, statusIcon))
		}
		if receivedUpdates[len(receivedUpdates)-1].Status != pb.OrderStatus_CONFIRMED {
			t.Errorf("%s%s expected last status CONFIRMED, got %v%s", red, fail, receivedUpdates[len(receivedUpdates)-1].Status, reset)
		} else {
			t.Logf("%s%s Status transition to CONFIRMED successful%s", green, check, reset)
		}
	}

	// Print summary table
	t.Log("\nStatus Transition Summary:")
	t.Log("| Step | Status         | Icon |")
	t.Log("|------|---------------|------|")
	for _, row := range summaryRows {
		t.Log(row)
	}
}

type mockOrderStatusStream struct {
	updates chan *pb.Order
	ctx     context.Context
}

func (m *mockOrderStatusStream) Send(order *pb.Order) error {
	m.updates <- order
	return nil
}
func (m *mockOrderStatusStream) Context() context.Context        { return m.ctx }
func (m *mockOrderStatusStream) RecvMsg(interface{}) error       { return nil }
func (m *mockOrderStatusStream) SendMsg(interface{}) error       { return nil }
func (m *mockOrderStatusStream) SendHeader(md metadata.MD) error { return nil }
func (m *mockOrderStatusStream) SetHeader(md metadata.MD) error  { return nil }
func (m *mockOrderStatusStream) SetTrailer(md metadata.MD)       {}
