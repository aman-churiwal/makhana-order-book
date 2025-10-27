package server

import (
	"aman/makhana/internal/router"
	"database/sql"
	"net/http"
)

func CreateServer(port string, db *sql.DB) *http.Server {

	r := router.CreateRouter(db)
	return &http.Server{
		Addr:    port,
		Handler: r,
	}
}
