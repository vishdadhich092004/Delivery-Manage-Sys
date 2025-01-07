package models

import "time"

type Agent struct {
	ID          uint      `json:"id"`
	WarehouseID uint      `json:"warehouse_id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	VehicleType string    `json:"vehicle_type"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type AgentCheckin struct {
	ID           uint       `json:"id"`
	AgentID      uint       `json:"agent_id"`
	WarehouseID  uint       `json:"warehouse_id"`
	CheckinTime  time.Time  `json:"checkin_time"`
	CheckoutTime *time.Time `json:"checkout_time"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
}

type AgentDailyStats struct {
	ID            uint      `json:"id"`
	AgentID       uint      `json:"agent_id"`
	Date          time.Time `json:"date"`
	TotalOrders   int       `json:"total_orders"`
	TotalDistance float64   `json:"total_distance"`
	TotalDuration int       `json:"total_duration"` // in minutes
	Earnings      float64   `json:"earnings"`
}
