package repository

import "payment-service/internal/domain"

type PaymentRepository interface {
	Create(p *domain.Payment) error
	GetByOrderID(orderID string) (*domain.Payment, error)
}
