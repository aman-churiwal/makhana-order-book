package router

import (
	"aman/makhana/internal/handler"
	"aman/makhana/internal/repository"
	"aman/makhana/internal/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func CreateRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		customerRoutes := v1.Group("/customers")
		{
			customerRepository := repository.NewCustomerRepository(db)
			customerService := service.NewCustomerService(customerRepository)
			customerHandler := handler.NewCustomerHandler(customerService)

			// Endpoints
			customerRoutes.GET("", customerHandler.GetAllCustomers)
			customerRoutes.POST("/create", customerHandler.CreateCustomer)
		}

	}

	return router
}
