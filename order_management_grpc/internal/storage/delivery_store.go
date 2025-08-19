package storage

import "order_management_grpc/internal/models"

type DeliveryStore interface {
	Create(delivery *models.Delivery) error
	Get(id string) (*models.Delivery, error)
	Update(delivery *models.Delivery) error
	Delete(id string) error
	List() ([]*models.Delivery, error)
}
