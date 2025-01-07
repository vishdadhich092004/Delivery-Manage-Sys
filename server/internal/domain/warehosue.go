package domain

import "time"

type Warehouse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
