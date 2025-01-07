package config

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database : %w", err)
	}

	// testing the db connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %w", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %w", err)
	}
	log.Println("DataBase Connected")
	return db, nil
}
