package repository

import (
	"aman/makhana/internal/models"
	"database/sql"
	"log"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetAllOrders() ([]*models.Order, error) {
	query := `SELECT id, customer_id, order_date, status, total_amount FROM orders`

	rows, err := r.db.Query(query)

	if err != nil {
		log.Printf("Error querying all orders: %v", err)
		return nil, err
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order := &models.Order{}
		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.OrderDate,
			&order.Status,
			&order.TotalAmount,
		)

		if err != nil {
			log.Printf("Error scanning orders row in GetAllOrders: %v", err)
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row iteration error in GetAllOrders: %v", err)
		return nil, err
	}

	return orders, nil
}
