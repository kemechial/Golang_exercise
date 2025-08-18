package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/order-management/internal/models"
	"github.com/order-management/internal/storage"
	pb "github.com/order-management/proto"
)

// OrderService implements the gRPC OrderService
type OrderService struct {
	pb.UnimplementedOrderServiceServer
	store     storage.OrderStore
	watchers  map[string][]chan *models.Order
	watcherMu sync.RWMutex
}

// NewOrderService creates a new order service
func NewOrderService(store storage.OrderStore) *OrderService {
	return &OrderService{
		store:    store,
		watchers: make(map[string][]chan *models.Order),
	}
}

// CreateOrder creates a new order
func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Validate request
	if req.CustomerId == "" {
		return nil, status.Error(codes.InvalidArgument, "customer_id is required")
	}
	if req.CustomerName == "" {
		return nil, status.Error(codes.InvalidArgument, "customer_name is required")
	}
	if len(req.Items) == 0 {
		return nil, status.Error(codes.InvalidArgument, "at least one item is required")
	}

	// Create order
	order := &models.Order{
		ID:           uuid.New().String(),
		CustomerID:   req.CustomerId,
		CustomerName: req.CustomerName,
		Status:       models.StatusPending,
	}

	// Convert items
	for _, pbItem := range req.Items {
		if pbItem.Quantity <= 0 {
			return nil, status.Error(codes.InvalidArgument, "item quantity must be positive")
		}
		if pbItem.UnitPrice <= 0 {
			return nil, status.Error(codes.InvalidArgument, "item unit price must be positive")
		}

		item := models.OrderItem{
			ProductID:   pbItem.ProductId,
			ProductName: pbItem.ProductName,
			Quantity:    pbItem.Quantity,
			UnitPrice:   pbItem.UnitPrice,
		}
		order.Items = append(order.Items, item)
	}

	// Calculate total
	order.CalculateTotal()

	// Store order
	if err := s.store.Create(order); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create order: %v", err))
	}

	// Notify watchers
	s.notifyWatchers(order.ID, order)

	return &pb.CreateOrderResponse{
		Order: order.ToProto(),
	}, nil
}

// GetOrder retrieves an order by ID
func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	if req.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "order_id is required")
	}

	order, err := s.store.GetByID(req.OrderId)
	if err != nil {
		if err == storage.ErrOrderNotFound {
			return nil, status.Error(codes.NotFound, "order not found")
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get order: %v", err))
	}

	return &pb.GetOrderResponse{
		Order: order.ToProto(),
	}, nil
}

// UpdateOrderStatus updates the status of an order
func (s *OrderService) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.UpdateOrderStatusResponse, error) {
	if req.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "order_id is required")
	}

	// Get existing order
	order, err := s.store.GetByID(req.OrderId)
	if err != nil {
		if err == storage.ErrOrderNotFound {
			return nil, status.Error(codes.NotFound, "order not found")
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get order: %v", err))
	}

	// Validate status transition
	newStatus := models.OrderStatus(req.Status)
	if !order.IsValidStatusTransition(newStatus) {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid status transition from %s to %s", order.Status.String(), newStatus.String()))
	}

	// Update status
	order.Status = newStatus
	order.UpdatedAt = time.Now()

	if err := s.store.Update(order); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update order: %v", err))
	}

	// Notify watchers
	s.notifyWatchers(order.ID, order)

	return &pb.UpdateOrderStatusResponse{
		Order: order.ToProto(),
	}, nil
}

// ListOrders lists orders with optional filtering and pagination
func (s *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10 // default page size
	}
	if pageSize > 100 {
		pageSize = 100 // max page size
	}

	offset := 0
	if req.PageToken != "" {
		var err error
		offset, err = strconv.Atoi(req.PageToken)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid page token")
		}
	}

	orders, total, err := s.store.List(req.CustomerId, offset, int(pageSize))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to list orders: %v", err))
	}

	// Convert to proto
	pbOrders := make([]*pb.Order, len(orders))
	for i, order := range orders {
		pbOrders[i] = order.ToProto()
	}

	// Calculate next page token
	nextPageToken := ""
	if offset+int(pageSize) < total {
		nextPageToken = strconv.Itoa(offset + int(pageSize))
	}

	return &pb.ListOrdersResponse{
		Orders:        pbOrders,
		NextPageToken: nextPageToken,
		TotalCount:    int32(total),
	}, nil
}

// CancelOrder cancels an order
func (s *OrderService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	if req.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "order_id is required")
	}

	// Get existing order
	order, err := s.store.GetByID(req.OrderId)
	if err != nil {
		if err == storage.ErrOrderNotFound {
			return nil, status.Error(codes.NotFound, "order not found")
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get order: %v", err))
	}

	// Check if order can be cancelled
	if !order.IsValidStatusTransition(models.StatusCancelled) {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("order cannot be cancelled in %s status", order.Status.String()))
	}

	// Update status to cancelled
	order.Status = models.StatusCancelled
	order.UpdatedAt = time.Now()

	if err := s.store.Update(order); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to cancel order: %v", err))
	}

	// Notify watchers
	s.notifyWatchers(order.ID, order)

	return &pb.CancelOrderResponse{
		Order: order.ToProto(),
	}, nil
}

// WatchOrderStatus streams order status updates
func (s *OrderService) WatchOrderStatus(req *pb.GetOrderRequest, stream pb.OrderService_WatchOrderStatusServer) error {
	if req.OrderId == "" {
		return status.Error(codes.InvalidArgument, "order_id is required")
	}

	// Check if order exists
	order, err := s.store.GetByID(req.OrderId)
	if err != nil {
		if err == storage.ErrOrderNotFound {
			return status.Error(codes.NotFound, "order not found")
		}
		return status.Error(codes.Internal, fmt.Sprintf("failed to get order: %v", err))
	}

	// Send current order state
	if err := stream.Send(order.ToProto()); err != nil {
		return err
	}

	// Create a channel for this watcher
	ch := make(chan *models.Order, 10)
	s.addWatcher(req.OrderId, ch)
	defer s.removeWatcher(req.OrderId, ch)

	// Stream updates
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case updatedOrder := <-ch:
			if err := stream.Send(updatedOrder.ToProto()); err != nil {
				return err
			}
		}
	}
}

// BatchCreateOrders creates multiple orders from a client stream
func (s *OrderService) BatchCreateOrders(stream pb.OrderService_BatchCreateOrdersServer) error {
	var createdOrders []*models.Order

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// End of stream, return the created orders
			pbOrders := make([]*pb.Order, len(createdOrders))
			for i, order := range createdOrders {
				pbOrders[i] = order.ToProto()
			}

			return stream.SendAndClose(&pb.ListOrdersResponse{
				Orders:     pbOrders,
				TotalCount: int32(len(createdOrders)),
			})
		}
		if err != nil {
			return err
		}

		// Create order (reuse logic from CreateOrder)
		order := &models.Order{
			ID:           uuid.New().String(),
			CustomerID:   req.CustomerId,
			CustomerName: req.CustomerName,
			Status:       models.StatusPending,
		}

		// Convert items
		for _, pbItem := range req.Items {
			item := models.OrderItem{
				ProductID:   pbItem.ProductId,
				ProductName: pbItem.ProductName,
				Quantity:    pbItem.Quantity,
				UnitPrice:   pbItem.UnitPrice,
			}
			order.Items = append(order.Items, item)
		}

		order.CalculateTotal()

		if err := s.store.Create(order); err != nil {
			return status.Error(codes.Internal, fmt.Sprintf("failed to create order: %v", err))
		}

		createdOrders = append(createdOrders, order)
		s.notifyWatchers(order.ID, order)
	}
}

// ProcessOrdersStream processes order status updates bidirectionally
func (s *OrderService) ProcessOrdersStream(stream pb.OrderService_ProcessOrdersStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// Update order status
		order, err := s.store.GetByID(req.OrderId)
		if err != nil {
			log.Printf("Failed to get order %s: %v", req.OrderId, err)
			continue
		}

		newStatus := models.OrderStatus(req.Status)
		if !order.IsValidStatusTransition(newStatus) {
			log.Printf("Invalid status transition for order %s: %s -> %s", req.OrderId, order.Status.String(), newStatus.String())
			continue
		}

		order.Status = newStatus
		order.UpdatedAt = time.Now()

		if err := s.store.Update(order); err != nil {
			log.Printf("Failed to update order %s: %v", req.OrderId, err)
			continue
		}

		// Send updated order back
		if err := stream.Send(order.ToProto()); err != nil {
			return err
		}

		s.notifyWatchers(order.ID, order)
	}
}

// Helper methods for watchers
func (s *OrderService) addWatcher(orderID string, ch chan *models.Order) {
	s.watcherMu.Lock()
	defer s.watcherMu.Unlock()
	s.watchers[orderID] = append(s.watchers[orderID], ch)
}

func (s *OrderService) removeWatcher(orderID string, ch chan *models.Order) {
	s.watcherMu.Lock()
	defer s.watcherMu.Unlock()

	channels := s.watchers[orderID]
	for i, c := range channels {
		if c == ch {
			s.watchers[orderID] = append(channels[:i], channels[i+1:]...)
			break
		}
	}

	if len(s.watchers[orderID]) == 0 {
		delete(s.watchers, orderID)
	}

	close(ch)
}

func (s *OrderService) notifyWatchers(orderID string, order *models.Order) {
	s.watcherMu.RLock()
	defer s.watcherMu.RUnlock()

	for _, ch := range s.watchers[orderID] {
		select {
		case ch <- order:
		default:
			// Channel is full, skip
		}
	}
}
