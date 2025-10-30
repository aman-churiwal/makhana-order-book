package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func CreateRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	RegisterCustomerRouter(api, db)
	RegisterOrderRouter(api, db)

	return router
}
