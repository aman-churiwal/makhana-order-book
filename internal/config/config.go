package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DbUrl   string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // default port
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, errors.New("database URL (DB_URL) is not set in environment variables")
	}

	return &Config{
		AppPort: port,
		DbUrl:   dbUrl,
	}, nil
}
