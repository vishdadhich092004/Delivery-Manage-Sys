package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vishdadhich092004/delivery-management-system/internal/config"
	"github.com/vishdadhich092004/delivery-management-system/internal/routes"
)

func main() {

	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatalf("Error setting up config : %v", err)
	}
	db, err := config.NewDB(cfg.POSTGRES_URI)
	if err != nil {
		log.Fatalf("Error Initialising the DataBase %v", err)
	}
	config.Migrate(db)
	router := routes.SetupRoutes(db)

	// to seed data
	// seed.SeedDB()
	log.Printf("Server is running on %s", cfg.PORT)
	router.Run(":" + cfg.PORT)

}
