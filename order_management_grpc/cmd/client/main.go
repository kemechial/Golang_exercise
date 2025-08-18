package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/order-management/proto"
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

	// Test the service
	testOrderService(client)
}

func testOrderService(client pb.OrderServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("=== Testing Order Management Service ===")

	// Test 1: Create an order
	log.Println("\n1. Creating an order...")
	createReq := &pb.CreateOrderRequest{
		CustomerId:   "customer-123",
		CustomerName: "John Doe",
		Items: []*pb.OrderItem{
			{
				ProductId:   "product-1",
				ProductName: "Laptop",
				Quantity:    1,
				UnitPrice:   999.99,
			},
			{
				ProductId:   "product-2",
				ProductName: "Mouse",
				Quantity:    2,
				UnitPrice:   25.50,
			},
		},
	}

	createResp, err := client.CreateOrder(ctx, createReq)
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}

	orderID := createResp.Order.Id
	log.Printf("Created order: %s (Total: $%.2f)", orderID, createResp.Order.TotalAmount)

	// Test 2: Get the order
	log.Println("\n2. Getting the order...")
	getResp, err := client.GetOrder(ctx, &pb.GetOrderRequest{OrderId: orderID})
	if err != nil {
		log.Fatalf("Failed to get order: %v", err)
	}
	log.Printf("Retrieved order: %s, Status: %s", getResp.Order.Id, getResp.Order.Status.String())

	// Test 3: Update order status
	log.Println("\n3. Updating order status...")
	updateResp, err := client.UpdateOrderStatus(ctx, &pb.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  pb.OrderStatus_CONFIRMED,
	})
	if err != nil {
		log.Fatalf("Failed to update order status: %v", err)
	}
	log.Printf("Updated order status to: %s", updateResp.Order.Status.String())

	// Test 4: List orders
	log.Println("\n4. Listing orders...")
	listResp, err := client.ListOrders(ctx, &pb.ListOrdersRequest{
		CustomerId: "customer-123",
		PageSize:   10,
	})
	if err != nil {
		log.Fatalf("Failed to list orders: %v", err)
	}
	log.Printf("Found %d orders for customer", len(listResp.Orders))
	for _, order := range listResp.Orders {
		log.Printf("  - Order %s: %s ($%.2f)", order.Id, order.Status.String(), order.TotalAmount)
	}

	// Test 5: Watch order status (server streaming)
	log.Println("\n5. Watching order status...")
	go watchOrderStatus(client, orderID)

	// Test 6: Update status again to trigger watcher
	time.Sleep(1 * time.Second)
	log.Println("\n6. Updating status to PREPARING...")
	_, err = client.UpdateOrderStatus(ctx, &pb.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  pb.OrderStatus_PREPARING,
	})
	if err != nil {
		log.Fatalf("Failed to update order status: %v", err)
	}

	// Test 7: Batch create orders (client streaming)
	log.Println("\n7. Batch creating orders...")
	testBatchCreateOrders(client)

	// Test 8: Cancel order
	log.Println("\n8. Cancelling order...")
	cancelResp, err := client.CancelOrder(ctx, &pb.CancelOrderRequest{OrderId: orderID})
	if err != nil {
		log.Fatalf("Failed to cancel order: %v", err)
	}
	log.Printf("Cancelled order: %s, Status: %s", cancelResp.Order.Id, cancelResp.Order.Status.String())

	// Wait a bit to see watcher output
	time.Sleep(2 * time.Second)
	log.Println("\n=== Testing completed ===")
}

func watchOrderStatus(client pb.OrderServiceClient, orderID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := client.WatchOrderStatus(ctx, &pb.GetOrderRequest{OrderId: orderID})
	if err != nil {
		log.Printf("Failed to watch order status: %v", err)
		return
	}

	log.Printf("Started watching order %s...", orderID)
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Watch stream error: %v", err)
			break
		}
		log.Printf("  [WATCH] Order %s status changed to: %s", order.Id, order.Status.String())
	}
}

func testBatchCreateOrders(client pb.OrderServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.BatchCreateOrders(ctx)
	if err != nil {
		log.Printf("Failed to create batch stream: %v", err)
		return
	}

	// Send multiple orders
	orders := []*pb.CreateOrderRequest{
		{
			CustomerId:   "customer-456",
			CustomerName: "Jane Smith",
			Items: []*pb.OrderItem{
				{ProductId: "product-3", ProductName: "Keyboard", Quantity: 1, UnitPrice: 89.99},
			},
		},
		{
			CustomerId:   "customer-789",
			CustomerName: "Bob Johnson",
			Items: []*pb.OrderItem{
				{ProductId: "product-4", ProductName: "Monitor", Quantity: 1, UnitPrice: 299.99},
				{ProductId: "product-5", ProductName: "Cable", Quantity: 3, UnitPrice: 12.99},
			},
		},
	}

	for _, order := range orders {
		if err := stream.Send(order); err != nil {
			log.Printf("Failed to send order: %v", err)
			return
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Failed to receive batch response: %v", err)
		return
	}

	log.Printf("Batch created %d orders", len(resp.Orders))
	for _, order := range resp.Orders {
		log.Printf("  - Created order %s for %s ($%.2f)", order.Id, order.CustomerName, order.TotalAmount)
	}
}
