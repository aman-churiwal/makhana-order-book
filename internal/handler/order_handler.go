package handler

import (
	"aman/makhana/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IOrderService interface {
	GetAllOrders() ([]*models.Order, error)
}

type OrderHandler struct {
	orderService IOrderService
}

func NewOrderHandler(orderService IOrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		log.Printf("Error retrieving orders: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if orders == nil {
		orders = []*models.Order{}
	}

	c.JSON(http.StatusOK, orders)
}
