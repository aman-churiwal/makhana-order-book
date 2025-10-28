package repository

import (
	"aman/makhana/internal/models"
	"database/sql"
	"log"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) CreateCustomer(customer *models.Customer) error {
	query := `INSERT INTO customers (name, contact, address) VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRow(
		query,
		customer.Name,
		customer.Contact,
		customer.Address,
	).Scan(&customer.ID)

	if err != nil {
		log.Printf("Error creating user: %v")
		return err
	}

	log.Printf("Customer created with ID: %d", customer.ID)
	return nil
}

func (r *CustomerRepository) GetAllCustomers() ([]*models.Customer, error) {
	query := `SELECT id, name, contact, address FROM customers`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error querying all customers: %v", err)
		return nil, err
	}
	defer rows.Close()

	var customers []*models.Customer

	for rows.Next() {
		customer := &models.Customer{}
		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Contact,
			&customer.Address,
		)

		if err != nil {
			log.Printf("Error scanning customer row in GetAllCustomers: %v", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error in GetAllCustomers: %v", err)
		return nil, err
	}

	return customers, nil
}
