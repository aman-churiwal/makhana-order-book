package service

import (
	"aman/makhana/internal/models"
	"errors"
	"log"
)

type IOrderRepository interface {
	GetAllOrders() ([]*models.Order, error)
}

type OrderService struct {
	orderRepository IOrderRepository
}

func NewOrderService(orderRepository IOrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (s *OrderService) GetAllOrders() ([]*models.Order, error) {
	orders, err := s.orderRepository.GetAllOrders()

	if err != nil {
		log.Printf("Error in GetAllOrders service: %v", err)
		return nil, errors.New("could not retrieve orders")
	}

	return orders, nil
}
