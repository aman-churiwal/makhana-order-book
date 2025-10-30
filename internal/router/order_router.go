package router

import (
	"aman/makhana/internal/handler"
	"aman/makhana/internal/repository"
	"aman/makhana/internal/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRouter(r *gin.RouterGroup, db *sql.DB) {
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	orderRoutes := r.Group("/orders")
	{
		orderRoutes.GET("/", orderHandler.GetAllOrders)
	}
}
