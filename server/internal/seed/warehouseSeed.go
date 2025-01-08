package seed

import (
	"fmt"
	"time"

	"github.com/vishdadhich092004/delivery-management-system/internal/seed/data"
	"gorm.io/gorm"
)

func SeedWarehouse(db *gorm.DB) {
	// Seed warehouses
	fmt.Println("Seeding warehouses...")
	for _, w := range data.Warehouses {
		w.CreatedAt = time.Now()
		if err := db.Create(&w).Error; err != nil {
			fmt.Printf("Error creating warehouse: %v\n", err)
		}
	}
}
