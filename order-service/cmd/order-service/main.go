package main

import (
	"database/sql"
	"log"

	"order-service/internal/repository"
	"order-service/internal/transport/http"
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/orders?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := &repository.PostgresOrderRepository{DB: db}
	uc := &usecase.OrderUsecase{Repo: repo}
	handler := &http.Handler{Usecase: uc}

	r := gin.Default()

	r.POST("/orders", handler.CreateOrder)

	r.Run(":8080")
}
