package service

import (
	"aman/makhana/internal/models"
	"errors"
	"log"
)

type ICustomerRepository interface {
	GetAllCustomers() ([]*models.Customer, error)
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
