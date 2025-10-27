package main

import (
	"aman/makhana/internal/config"
	"aman/makhana/internal/database"
	"aman/makhana/internal/server"
	"log"
)

func main() {
	log.Println("Application Started")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	db, err := database.StartConnection(cfg.DbUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Close database connection whenever the main function exits
	defer db.Close()

	newServer := server.CreateServer(cfg.AppPort, db)
	log.Printf("Starting server on port %s", cfg.AppPort)
	err = newServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
