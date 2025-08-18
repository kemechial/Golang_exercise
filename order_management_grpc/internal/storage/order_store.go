package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/order-management/internal/models"
)

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrOrderExists   = errors.New("order already exists")
)

// OrderStore interface defines the contract for order storage
type OrderStore interface {
	Create(order *models.Order) error
	GetByID(id string) (*models.Order, error)
	Update(order *models.Order) error
	Delete(id string) error
	List(customerID string, offset, limit int) ([]*models.Order, int, error)
	GetByCustomerID(customerID string) ([]*models.Order, error)
}

// InMemoryOrderStore implements OrderStore using in-memory storage
type InMemoryOrderStore struct {
	orders map[string]*models.Order
	mutex  sync.RWMutex
}

// NewInMemoryOrderStore creates a new in-memory order store
func NewInMemoryOrderStore() *InMemoryOrderStore {
	return &InMemoryOrderStore{
		orders: make(map[string]*models.Order),
	}
}

// Create stores a new order
func (s *InMemoryOrderStore) Create(order *models.Order) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.orders[order.ID]; exists {
		return ErrOrderExists
	}

	// Make a copy to avoid external modifications
	orderCopy := *order
	orderCopy.CreatedAt = time.Now()
	orderCopy.UpdatedAt = orderCopy.CreatedAt

	s.orders[order.ID] = &orderCopy
	return nil
}

// GetByID retrieves an order by its ID
func (s *InMemoryOrderStore) GetByID(id string) (*models.Order, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	order, exists := s.orders[id]
	if !exists {
		return nil, ErrOrderNotFound
	}

	// Return a copy to avoid external modifications
	orderCopy := *order
	return &orderCopy, nil
}

// Update modifies an existing order
func (s *InMemoryOrderStore) Update(order *models.Order) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.orders[order.ID]; !exists {
		return ErrOrderNotFound
	}

	// Make a copy and update timestamp
	orderCopy := *order
	orderCopy.UpdatedAt = time.Now()

	s.orders[order.ID] = &orderCopy
	return nil
}

// Delete removes an order by its ID
func (s *InMemoryOrderStore) Delete(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.orders[id]; !exists {
		return ErrOrderNotFound
	}

	delete(s.orders, id)
	return nil
}

// List returns a paginated list of orders, optionally filtered by customer ID
func (s *InMemoryOrderStore) List(customerID string, offset, limit int) ([]*models.Order, int, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var allOrders []*models.Order

	// Filter by customer ID if provided
	for _, order := range s.orders {
		if customerID == "" || order.CustomerID == customerID {
			orderCopy := *order
			allOrders = append(allOrders, &orderCopy)
		}
	}

	total := len(allOrders)

	// Apply pagination
	start := offset
	if start > total {
		start = total
	}

	end := start + limit
	if end > total {
		end = total
	}

	if start >= end {
		return []*models.Order{}, total, nil
	}

	return allOrders[start:end], total, nil
}

// GetByCustomerID returns all orders for a specific customer
func (s *InMemoryOrderStore) GetByCustomerID(customerID string) ([]*models.Order, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var customerOrders []*models.Order

	for _, order := range s.orders {
		if order.CustomerID == customerID {
			orderCopy := *order
			customerOrders = append(customerOrders, &orderCopy)
		}
	}

	return customerOrders, nil
}

// Count returns the total number of orders
func (s *InMemoryOrderStore) Count() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.orders)
}

// Clear removes all orders (useful for testing)
func (s *InMemoryOrderStore) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.orders = make(map[string]*models.Order)
}
