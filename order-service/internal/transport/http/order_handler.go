package http

import (
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase *usecase.OrderUsecase
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req struct {
		CustomerID string `json:"customer_id"`
		ItemName   string `json:"item_name"`
		Amount     int64  `json:"amount"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := h.Usecase.CreateOrder(req.CustomerID, req.ItemName, req.Amount)
	if err != nil {
		c.JSON(503, gin.H{"error": "Payment service unavailable"})
		return
	}

	c.JSON(200, order)
}
