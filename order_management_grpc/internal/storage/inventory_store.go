package storage

import "order_management_grpc/internal/models"

type InventoryStore interface {
	Create(item *models.InventoryItem) error
	Get(id string) (*models.InventoryItem, error)
	Update(item *models.InventoryItem) error
	Delete(id string) error
	List() ([]*models.InventoryItem, error)
}
