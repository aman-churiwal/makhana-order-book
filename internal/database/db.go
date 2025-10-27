package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func StartConnection(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}
