package seed

import (
	"fmt"
	"time"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

func SeedOrder(db *gorm.DB) {
	fmt.Println("Seeding orders...")
	var warehouses []domain.Warehouse
	db.Find(&warehouses)
	for _, warehouse := range warehouses {
		// Create 1200 orders per warehouse (60 orders * 20 agents)
		for i := 1; i <= 25; i++ {
			// Random coordinates within ~5km of warehouse
			latitude := warehouse.Latitude + (rand.Float64()*0.1 - 0.05)
			longitude := warehouse.Longitude + (rand.Float64()*0.1 - 0.05)

			order := domain.Order{
				WarehouseID:     warehouse.ID,
				CustomerName:    fmt.Sprintf("Customer %d-%d", warehouse.ID, i),
				DeliveryAddress: fmt.Sprintf("Address %d-%d", warehouse.ID, i),
				Latitude:        latitude,
				Longitude:       longitude,
				Status:          "PENDING",
				ScheduledDate:   time.Now().Truncate(24 * time.Hour),
				CreatedAt:       time.Now().Add(-time.Duration(rand.Int63n(24)) * time.Hour),
			}
			if err := db.Create(&order).Error; err != nil {
				fmt.Printf("Error creating order: %v\n", err)
			}
		}
	}
}
