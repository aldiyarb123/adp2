package usecase

import (
	"payment-service/internal/domain"
	"payment-service/internal/repository"

	"github.com/google/uuid"
)

type PaymentUsecase struct {
	Repo repository.PaymentRepository
}

func (u *PaymentUsecase) ProcessPayment(orderID string, amount int64) (*domain.Payment, error) {

	status := "Authorized"

	if amount > 100000 {
		status = "Declined"
	}

	payment := &domain.Payment{
		ID:            uuid.New().String(),
		OrderID:       orderID,
		TransactionID: uuid.New().String(),
		Amount:        amount,
		Status:        status,
	}

	err := u.Repo.Create(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
