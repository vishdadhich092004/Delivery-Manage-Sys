package models

import "time"

type Order struct {
	ID              uint      `json:"id"`
	WarehouseID     uint      `json:"warehouse_id"`
	CustomerName    string    `json:"customer_name"`
	DeliveryAddress string    `json:"delivery_address"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	Status          string    `json:"status"`
	ScheduledDate   time.Time `json:"scheduled_date"`
	CreatedAt       time.Time `json:"created_at"`
}
type OrderAssignment struct {
	ID                uint       `json:"id"`
	OrderID           uint       `json:"order_id"`
	AgentID           uint       `json:"agent_id"`
	AssignedAt        time.Time  `json:"assigned_at"`
	EstimatedDistance float64    `json:"estimated_distance"`
	EstimatedDuration int        `json:"estimated_duration"`
	Status            string     `json:"status"`
	CompletedAt       *time.Time `json:"completed_at"`
}
