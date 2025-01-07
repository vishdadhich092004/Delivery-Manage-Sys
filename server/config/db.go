package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database : %w", err)
	}

	// testing the db connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database %w", err)
	}
	log.Println("DataBase Connected")
	return db, nil
}
