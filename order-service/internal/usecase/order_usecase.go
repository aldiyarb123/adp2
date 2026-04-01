package usecase

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"order-service/internal/domain"
	"order-service/internal/repository"

	"github.com/google/uuid"
)

type OrderUsecase struct {
	Repo repository.OrderRepository
}

func (u *OrderUsecase) CreateOrder(customerID, item string, amount int64) (*domain.Order, error) {

	// initialize new order entity
	orderID := uuid.New().String()

	order := &domain.Order{
		ID:         orderID,
		CustomerID: customerID,
		ItemName:   item,
		Amount:     amount,
		Status:     "Pending",
	}

	err := u.Repo.Create(order)
	if err != nil {
		return nil, err
	}

	// HTTP client с timeout (ВАЖНО)
	// create HTTP client with timeout
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	body := map[string]interface{}{
		"order_id": order.ID,
		"amount":   amount,
	}

	jsonData, _ := json.Marshal(body)

	resp, err := client.Post("http://localhost:8081/payments", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		order.Status = "Failed"
		u.Repo.Update(order)
		return order, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		order.Status = "Paid"
	} else {
		order.Status = "Failed"
	}

	u.Repo.Update(order)
	return order, nil
}
