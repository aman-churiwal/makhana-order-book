package handler

import (
	"aman/makhana/internal/models"
	"aman/makhana/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ICustomerService interface {
	GetAllCustomers() ([]*models.Customer, error)
	CreateCustomer(request *models.Customer) (*models.Customer, error)
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

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var createCustomerRequest service.CreateCustomerRequest

	if err := c.ShouldBindJSON(&createCustomerRequest); err != nil {
		log.Printf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	customer, err := h.customerService.CreateCustomer(createCustomerRequest)
	if err != nil {
		log.Printf("Error creating customer: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}
