package config

import (
	"database/sql"
	"log"
)

func NewDB() {
	connStr := ""
	db, err := sql.Open("postrges", connStr)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("DataBase Connected")
}
