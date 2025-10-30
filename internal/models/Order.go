package models

import "time"

type OrderStatus string

const (
	OrderStatusPending          OrderStatus = "PENDING"
	OrderStatusDelivered        OrderStatus = "DELIVERED"
	OrderStatusCancelled        OrderStatus = "CANCELLED"
	OrderStatusPaymentCompleted OrderStatus = "PAYMENT_COMPLETED"
)

type Order struct {
	ID          int         `json:"id"`
	CustomerID  int         `json:"customer_id"`
	OrderDate   time.Time   `json:"order_date"`
	Status      OrderStatus `json:"status"`
	TotalAmount float64     `json:"total_amount"`
}

type OrderDetail struct {
	ID            int     `json:"id"`
	OrderID       int     `json:"order_id"`
	MakhanaTypeID int     `json:"makhana_type_id"`
	Quantity      float64 `json:"quantity"`
	Rate          float64 `json:"rate"`
	LineItemTotal float64 `json:"line_item_total"`
}
