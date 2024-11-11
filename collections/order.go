package collections

import "time"

type Order struct {
	ID              int         `json:"orderId"`
	CustomerID      int         `json:"customerId"`
	CustomerName    string      `json:"customerName"`
	CustomerPhone   string      `json:"customerPhone"`
	ShippingAddress Address     `json:"shippingAddress"`
	PaymentMethod   int         `json:"paymentMethod"`
	TotalAmount     float64     `json:"totalAmount"`
	FinalAmount     float64     `json:"finalAmount"`
	OrderDate       time.Time   `json:"orderDate"`
	Status          string      `json:"status"`
	OrderItems      []OrderItem `json:"orderItems"`
}
