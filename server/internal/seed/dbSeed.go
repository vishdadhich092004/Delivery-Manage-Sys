package seed

import (
	"fmt"
	"os"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SeedDB() {

	dsn := os.Getenv("POSTGRES_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	// seed warehouse
	SeedWarehouse(db)

	// seed agent
	SeedAgent(db)

	// seed agentCheckin
	SeedAgentCheckin(db)

	// Seed orders
	SeedOrder(db)

	// seed agent  daily stats
	SeedAgentDailyStats(db)

	fmt.Println("Seeding completed!")

	// Print counts
	var warehouseCount, agentCount, orderCount int64
	db.Model(&domain.Warehouse{}).Count(&warehouseCount)
	db.Model(&domain.Agent{}).Count(&agentCount)
	db.Model(&domain.Order{}).Count(&orderCount)

	fmt.Printf("\nFinal counts:\n")
	fmt.Printf("Warehouses: %d\n", warehouseCount)
	fmt.Printf("Agents: %d\n", agentCount)
	fmt.Printf("Orders: %d\n", orderCount)

}
