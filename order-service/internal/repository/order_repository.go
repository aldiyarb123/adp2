package repository

import "order-service/internal/domain"

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id string) (*domain.Order, error)
	Update(order *domain.Order) error
}
