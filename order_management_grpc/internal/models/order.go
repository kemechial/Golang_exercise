package models

import (
	"time"

	pb "order_management_grpc/proto/orderpb/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Order represents the internal order model
type Order struct {
	ID           string      `json:"id"`
	CustomerID   string      `json:"customer_id"`
	CustomerName string      `json:"customer_name"`
	Items        []OrderItem `json:"items"`
	Status       OrderStatus `json:"status"`
	TotalAmount  float64     `json:"total_amount"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int32   `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

// OrderStatus represents the order status
type OrderStatus int32

const (
	StatusPending   OrderStatus = 0
	StatusConfirmed OrderStatus = 1
	StatusPreparing OrderStatus = 2
	StatusShipped   OrderStatus = 3
	StatusDelivered OrderStatus = 4
	StatusCancelled OrderStatus = 5
)

// String returns the string representation of OrderStatus
func (s OrderStatus) String() string {
	switch s {
	case StatusPending:
		return "PENDING"
	case StatusConfirmed:
		return "CONFIRMED"
	case StatusPreparing:
		return "PREPARING"
	case StatusShipped:
		return "SHIPPED"
	case StatusDelivered:
		return "DELIVERED"
	case StatusCancelled:
		return "CANCELLED"
	default:
		return "UNKNOWN"
	}
}

// ToProto converts internal Order model to protobuf Order
func (o *Order) ToProto() *pb.Order {
	pbItems := make([]*pb.OrderItem, len(o.Items))
	for i, item := range o.Items {
		pbItems[i] = &pb.OrderItem{
			ProductId:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
			TotalPrice:  item.TotalPrice,
		}
	}

	return &pb.Order{
		Id:           o.ID,
		CustomerId:   o.CustomerID,
		CustomerName: o.CustomerName,
		Items:        pbItems,
		Status:       pb.OrderStatus(o.Status),
		TotalAmount:  o.TotalAmount,
		CreatedAt:    timestamppb.New(o.CreatedAt),
		UpdatedAt:    timestamppb.New(o.UpdatedAt),
	}
}

// FromProto converts protobuf Order to internal Order model
func FromProtoOrder(pbOrder *pb.Order) *Order {
	items := make([]OrderItem, len(pbOrder.Items))
	for i, pbItem := range pbOrder.Items {
		items[i] = OrderItem{
			ProductID:   pbItem.ProductId,
			ProductName: pbItem.ProductName,
			Quantity:    pbItem.Quantity,
			UnitPrice:   pbItem.UnitPrice,
			TotalPrice:  pbItem.TotalPrice,
		}
	}

	return &Order{
		ID:           pbOrder.Id,
		CustomerID:   pbOrder.CustomerId,
		CustomerName: pbOrder.CustomerName,
		Items:        items,
		Status:       OrderStatus(pbOrder.Status),
		TotalAmount:  pbOrder.TotalAmount,
		CreatedAt:    pbOrder.CreatedAt.AsTime(),
		UpdatedAt:    pbOrder.UpdatedAt.AsTime(),
	}
}

// CalculateTotal calculates the total amount of the order
func (o *Order) CalculateTotal() {
	total := 0.0
	for i := range o.Items {
		o.Items[i].TotalPrice = float64(o.Items[i].Quantity) * o.Items[i].UnitPrice
		total += o.Items[i].TotalPrice
	}
	o.TotalAmount = total
}

// IsValidStatusTransition checks if the status transition is valid
func (o *Order) IsValidStatusTransition(newStatus OrderStatus) bool {
	// Define valid transitions
	validTransitions := map[OrderStatus][]OrderStatus{
		StatusPending:   {StatusConfirmed, StatusCancelled},
		StatusConfirmed: {StatusPreparing, StatusCancelled},
		StatusPreparing: {StatusShipped, StatusCancelled},
		StatusShipped:   {StatusDelivered},
		StatusDelivered: {},
		StatusCancelled: {},
	}

	allowedStatuses, exists := validTransitions[o.Status]
	if !exists {
		return false
	}

	for _, status := range allowedStatuses {
		if status == newStatus {
			return true
		}
	}
	return false
}
