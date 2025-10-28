package service

import (
	"aman/makhana/internal/models"
	"errors"
	"log"
)

type CreateCustomerRequest struct {
	Name    string
	Contact string
	Address string
}

type ICustomerRepository interface {
	GetAllCustomers() ([]*models.Customer, error)
	CreateCustomer(customer *models.Customer) error
	GetCustomerByID(id int64) (*models.Customer, error)
}

type CustomerService struct {
	customerRepository ICustomerRepository
}

func NewCustomerService(customerRepository ICustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (s *CustomerService) GetAllCustomers() ([]*models.Customer, error) {
	customers, err := s.customerRepository.GetAllCustomers()
	if err != nil {
		log.Printf("Service Error: %v", err)
		return nil, errors.New("could not retrieve customers")
	}

	return customers, nil
}

func (s *CustomerService) CreateCustomer(request CreateCustomerRequest) (*models.Customer, error) {
	if request.Name == "" {
		return nil, errors.New("name is required")
	}
	if request.Contact == "" {
		return nil, errors.New("contact is required")
	}

	customer := &models.Customer{
		Name:    request.Name,
		Contact: request.Contact,
		Address: request.Address,
	}

	err := s.customerRepository.CreateCustomer(customer)
	if err != nil {
		log.Printf("Service Error: %v", err)
		return nil, errors.New("error creating customer")
	}

	return customer, nil
}

func (s *CustomerService) GetCustomerByID(id int64) (*models.Customer, error) {
	customer, err := s.customerRepository.GetCustomerByID(id)
	if err != nil {
		log.Printf("Service Error: %v", err)
		return nil, errors.New("could not retrieve customer")
	}

	if customer == nil {
		return nil, errors.New("customer not found")
	}

	return customer, nil
}
