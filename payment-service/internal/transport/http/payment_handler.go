package http

import (
	"payment-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase *usecase.PaymentUsecase
}

func (h *Handler) CreatePayment(c *gin.Context) {
	var req struct {
		OrderID string `json:"order_id"`
		Amount  int64  `json:"amount"`
	}
	// parse incoming JSON request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if req.Amount <= 0 {
		c.JSON(400, gin.H{"error": "invalid amount"})
		return
	}

	payment, err := h.Usecase.ProcessPayment(req.OrderID, req.Amount)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if payment.Status == "Declined" {
		c.JSON(400, payment)
		return
	}

	c.JSON(200, payment)

}
