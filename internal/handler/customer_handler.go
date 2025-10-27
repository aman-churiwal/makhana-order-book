package handler

import (
	"aman/makhana/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ICustomerService interface {
	GetAllCustomers() ([]*models.Customer, error)
}

type CustomerHandler struct {
	customerService ICustomerService
}

func NewCustomerHandler(customerService ICustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	customers, err := h.customerService.GetAllCustomers()
	if err != nil {
		log.Printf("Error retrieving customers: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if customers == nil {
		customers = []*models.Customer{}
	}

	c.JSON(http.StatusOK, customers)
}
