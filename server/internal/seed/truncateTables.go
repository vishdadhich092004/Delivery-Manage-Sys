package seed

import (
	"fmt"

	"gorm.io/gorm"
)

func TruncateTables(db *gorm.DB) {
	// Disable foreign key checks to allow truncation of tables with relationships
	db.Exec("SET session_replication_role = 'replica';")

	// Truncate all tables
	tables := []string{"agent_daily_stats", "agent_checkins", "orders", "agents", "warehouses", "order_assignments"} // Add more table names as needed
	for _, table := range tables {
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", table))
	}

	// Re-enable foreign key checks
	db.Exec("SET session_replication_role = 'origin';")

	fmt.Println("All tables truncated successfully.")
}
