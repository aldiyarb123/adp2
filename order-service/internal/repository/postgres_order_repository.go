package repository

import (
	"database/sql"
	"order-service/internal/domain"
)

type PostgresOrderRepository struct {
	DB *sql.DB
}

func (r *PostgresOrderRepository) Create(order *domain.Order) error {
	_, err := r.DB.Exec(
		"INSERT INTO orders (id, customer_id, item_name, amount, status) VALUES ($1,$2,$3,$4,$5)",
		order.ID, order.CustomerID, order.ItemName, order.Amount, order.Status,
	)
	return err
}

func (r *PostgresOrderRepository) GetByID(id string) (*domain.Order, error) {
	row := r.DB.QueryRow("SELECT id, customer_id, item_name, amount, status FROM orders WHERE id=$1", id)

	var o domain.Order
	err := row.Scan(&o.ID, &o.CustomerID, &o.ItemName, &o.Amount, &o.Status)
	return &o, err
}

func (r *PostgresOrderRepository) Update(order *domain.Order) error {
	_, err := r.DB.Exec("UPDATE orders SET status=$1 WHERE id=$2", order.Status, order.ID)
	return err
}
