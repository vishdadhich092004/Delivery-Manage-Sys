package config

import (
	"log"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.Agent{},
		&domain.AgentCheckin{},
		&domain.AgentDailyStats{},
		&domain.Order{},
		&domain.OrderAssignment{},
		&domain.Warehouse{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database : %v", err)
	}
	log.Println("database tables created successfully")
}
