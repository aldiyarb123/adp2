package repository

import (
	"database/sql"
	"payment-service/internal/domain"
)

type PostgresPaymentRepository struct {
	DB *sql.DB
}

func (r *PostgresPaymentRepository) Create(p *domain.Payment) error {
	_, err := r.DB.Exec(
		"INSERT INTO payments (id, order_id, transaction_id, amount, status) VALUES ($1,$2,$3,$4,$5)",
		p.ID, p.OrderID, p.TransactionID, p.Amount, p.Status,
	)
	return err
}

func (r *PostgresPaymentRepository) GetByOrderID(orderID string) (*domain.Payment, error) {
	row := r.DB.QueryRow("SELECT id, order_id, transaction_id, amount, status FROM payments WHERE order_id=$1", orderID)

	var p domain.Payment
	err := row.Scan(&p.ID, &p.OrderID, &p.TransactionID, &p.Amount, &p.Status)
	return &p, err
}
