AP2 Assignment 1 – Order & Payment Microservices

📌 Overview

This project implements a microservices-based system using Clean Architecture in Go.

The system consists of two independent services:
	•	Order Service
	•	Payment Service

These services communicate with each other using REST API.

⸻

🏗 Architecture

Each service follows Clean Architecture principles:
	•	Domain Layer – business entities (Order, Payment)
	•	Usecase Layer – business logic
	•	Repository Layer – database operations
	•	Transport Layer (HTTP) – REST API handlers

⸻

🔗 Microservices Principles
	•	Each service has its own database
	•	No shared code between services
	•	Communication via REST (HTTP)
	•	Timeout implemented for resilience (2 seconds)

⸻

⚙️ Technologies Used
	•	Go (Golang)
	•	Gin (HTTP framework)
	•	PostgreSQL
	•	Docker

⸻

🚀 How to Run

1. Start databases
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=orders postgres

docker run -d -p 5433:5432 -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=payments postgres



2. Run services
cd payment-service
go run cmd/payment-service/main.go

cd order-service
go run cmd/order-service/main.go




📡 API Endpoints

Order Service
	•	POST /orders – create order
	•	GET /orders/{id} – get order
	•	PATCH /orders/{id}/cancel – cancel order

Payment Service
	•	POST /payments – process payment
	•	GET /payments/{order_id} – get payment status

⸻

⚠️ Business Rules
	•	Amount is stored as int64 (no float)
	•	Orders must have amount > 0
	•	If amount > 100000 → payment is Declined
	•	Paid orders cannot be cancelled

⸻

❗ Failure Handling
	•	HTTP client timeout = 2 seconds
	•	If Payment Service is unavailable:
	•	Order Service returns 503
	•	Order status becomes Failed

⸻

📊 System Flow
	1.	Order is created with status Pending
	2.	Order Service calls Payment Service
	3.	Payment is processed:
	•	Authorized → Order becomes Paid
	•	Declined → Order becomes Failed

⸻

📁 Project Structure
order-service/
payment-service/


Each service contains:
cmd/
internal/domain/
internal/usecase/
internal/repository/
internal/transport/http/


🎯 Conclusion

This project demonstrates:
	•	Clean Architecture
	•	Microservices design
	•	REST communication
	•	Fault tolerance using timeouts



    ---

## 🚀 ЧТО ДАЛЬШЕ

Теперь сделай:

```bash
git add README.md
git commit -m "add professional README documentation"
git push