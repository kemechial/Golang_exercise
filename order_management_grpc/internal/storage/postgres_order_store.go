package storage

import (
	"context"
	"time"

	"order_management_grpc/internal/models"

	"github.com/jackc/pgx/v5"
)

type PostgresOrderStore struct {
	db *pgx.Conn
}

func NewPostgresOrderStore(conn *pgx.Conn) *PostgresOrderStore {
	return &PostgresOrderStore{db: conn}
}

func (s *PostgresOrderStore) Create(order *models.Order) error {
	// Example: Insert order into DB (simplified, no items)
	_, err := s.db.Exec(context.Background(),
		`INSERT INTO orders (id, customer_id, customer_name, status, total_amount, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		order.ID, order.CustomerID, order.CustomerName, order.Status, order.TotalAmount, order.CreatedAt, order.UpdatedAt)
	return err
}

func (s *PostgresOrderStore) GetByID(id string) (*models.Order, error) {
	row := s.db.QueryRow(context.Background(),
		`SELECT id, customer_id, customer_name, status, total_amount, created_at, updated_at FROM orders WHERE id=$1`, id)
	var order models.Order
	if err := row.Scan(&order.ID, &order.CustomerID, &order.CustomerName, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt); err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *PostgresOrderStore) Update(order *models.Order) error {
	_, err := s.db.Exec(context.Background(),
		`UPDATE orders SET status=$1, total_amount=$2, updated_at=$3 WHERE id=$4`,
		order.Status, order.TotalAmount, time.Now(), order.ID)
	return err
}

func (s *PostgresOrderStore) Delete(id string) error {
	_, err := s.db.Exec(context.Background(), `DELETE FROM orders WHERE id=$1`, id)
	return err
}

func (s *PostgresOrderStore) List(customerID string, offset, limit int) ([]*models.Order, int, error) {
	var rows pgx.Rows
	var err error
	if customerID == "" {
		rows, err = s.db.Query(context.Background(), `SELECT id, customer_id, customer_name, status, total_amount, created_at, updated_at FROM orders OFFSET $1 LIMIT $2`, offset, limit)
	} else {
		rows, err = s.db.Query(context.Background(), `SELECT id, customer_id, customer_name, status, total_amount, created_at, updated_at FROM orders WHERE customer_id=$1 OFFSET $2 LIMIT $3`, customerID, offset, limit)
	}
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var orders []*models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.CustomerName, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, 0, err
		}
		orders = append(orders, &order)
	}
	return orders, len(orders), nil
}

func (s *PostgresOrderStore) GetByCustomerID(customerID string) ([]*models.Order, error) {
	rows, err := s.db.Query(context.Background(), `SELECT id, customer_id, customer_name, status, total_amount, created_at, updated_at FROM orders WHERE customer_id=$1`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []*models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.CustomerName, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
