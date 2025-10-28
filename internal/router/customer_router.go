package router

import (
	"aman/makhana/internal/handler"
	"aman/makhana/internal/repository"
	"aman/makhana/internal/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRouter(r *gin.RouterGroup, db *sql.DB) {

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	customerRoutes := r.Group("/customers")
	{
		customerRoutes.GET("", customerHandler.GetAllCustomers)
		customerRoutes.POST("/create", customerHandler.CreateCustomer)
		customerRoutes.GET("/:id", customerHandler.GetCustomerById)
	}
}
