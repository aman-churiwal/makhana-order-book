package repository

import (
	"aman/makhana/internal/models"
	"database/sql"
	"errors"
	"fmt"
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
		log.Printf("Error creating user: %v", err)
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

func (r *CustomerRepository) GetCustomerByID(id int64) (*models.Customer, error) {
	query := `SELECT id, name, contact, address FROM customers WHERE id = $1`

	customer := &models.Customer{}
	err := r.db.QueryRow(query, id).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Contact,
		&customer.Address,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No customer found with ID: %d", id)
			return nil, fmt.Errorf(`no customer found with ID: %d`, id)
		}

		log.Printf("Error querying customer with ID: %d: %v", id, err)
		return nil, err
	}

	return customer, nil
}
