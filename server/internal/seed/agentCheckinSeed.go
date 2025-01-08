package seed

import (
	"fmt"
	"time"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

func SeedAgentCheckin(db *gorm.DB) {
	// Seed agent check-ins
	fmt.Println("Seeding agent check-ins...")
	var agents []domain.Agent
	db.Find(&agents)

	for _, agent := range agents {
		checkin := domain.AgentCheckin{
			AgentID:     agent.ID,
			WarehouseID: agent.WarehouseID,
			CheckinTime: time.Now().Truncate(24 * time.Hour).Add(8*time.Hour + time.Duration(rand.Int63n(60))*time.Minute),
			Status:      "CHECKED_IN",
			CreatedAt:   time.Now(),
		}
		if err := db.Create(&checkin).Error; err != nil {
			fmt.Printf("Error creating check-in: %v\n", err)
		}
	}
}
