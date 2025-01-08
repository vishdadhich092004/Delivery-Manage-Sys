package seed

import (
	"fmt"
	"time"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"gorm.io/gorm"
)

func SeedAgent(db *gorm.DB) {
	// Seed agents
	fmt.Println("Seeding agents...")
	for warehouseID := uint(1); warehouseID <= 10; warehouseID++ {
		for agentCount := 1; agentCount <= 5; agentCount++ {
			agent := domain.Agent{
				WarehouseID: warehouseID,
				Name:        fmt.Sprintf("Agent %d-%d", warehouseID, agentCount),
				Phone:       fmt.Sprintf("98765%02d%02d", warehouseID, agentCount),
				VehicleType: map[bool]string{true: "BIKE", false: "SCOOTER"}[agentCount%2 == 0],
				Status:      "ACTIVE",
				CreatedAt:   time.Now(),
			}
			if err := db.Create(&agent).Error; err != nil {
				fmt.Printf("Error creating agent: %v\n", err)
			}
		}
	}
}

func SeedAgentDailyStats(db *gorm.DB) {

	fmt.Println("Initializing agent daily stats...")
	var agents []domain.Agent
	db.Find(&agents)
	for _, agent := range agents {
		stats := domain.AgentDailyStats{
			AgentID:       agent.ID,
			Date:          time.Now().Truncate(24 * time.Hour),
			TotalOrders:   0,
			TotalDistance: 0,
			TotalDuration: 0,
			Earnings:      0,
		}
		if err := db.Create(&stats).Error; err != nil {
			fmt.Printf("Error creating daily stats: %v\n", err)
		}
	}
}
