package service

import (
	"context"
	inventorypb "order_management_grpc/proto/inventorypb"
	"order_management_grpc/internal/storage"
)

type InventoryService struct {
	inventorypb.UnimplementedInventoryServiceServer
	store storage.InventoryStore
}

func NewInventoryService(store storage.InventoryStore) *InventoryService {
	return &InventoryService{
		store: store,
	}
}

// CreateInventoryItem implements the CreateInventoryItem RPC
func (s *InventoryService) CreateInventoryItem(
	ctx context.Context,
	req *inventorypb.CreateInventoryItemRequest,
) (*inventorypb.CreateInventoryItemResponse, error) {
	// TODO: Add storage logic
	return &inventorypb.CreateInventoryItemResponse{Item: req.GetItem()}, nil
}

// GetInventoryItem implements the GetInventoryItem RPC
func (s *InventoryService) GetInventoryItem(
	ctx context.Context,
	req *inventorypb.GetInventoryItemRequest,
) (*inventorypb.GetInventoryItemResponse, error) {
	// TODO: Add storage logic
	return &inventorypb.GetInventoryItemResponse{Item: &inventorypb.InventoryItem{Id: req.GetId()}}, nil
}

// UpdateInventoryItem implements the UpdateInventoryItem RPC
func (s *InventoryService) UpdateInventoryItem(
	ctx context.Context,
	req *inventorypb.UpdateInventoryItemRequest,
) (*inventorypb.UpdateInventoryItemResponse, error) {
	// TODO: Add storage logic
	return &inventorypb.UpdateInventoryItemResponse{Item: req.GetItem()}, nil
}

// DeleteInventoryItem implements the DeleteInventoryItem RPC
func (s *InventoryService) DeleteInventoryItem(
	ctx context.Context,
	req *inventorypb.DeleteInventoryItemRequest,
) (*inventorypb.DeleteInventoryItemResponse, error) {
	// TODO: Add storage logic
	return &inventorypb.DeleteInventoryItemResponse{Success: true}, nil
}

// ListInventoryItems implements the ListInventoryItems RPC
func (s *InventoryService) ListInventoryItems(
	ctx context.Context,
	req *inventorypb.ListInventoryItemsRequest,
) (*inventorypb.ListInventoryItemsResponse, error) {
	// TODO: Add storage logic
	return &inventorypb.ListInventoryItemsResponse{Items: []*inventorypb.InventoryItem{}}, nil
}
