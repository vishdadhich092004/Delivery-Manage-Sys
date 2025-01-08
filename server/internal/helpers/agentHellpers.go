package helpers

import (
	"time"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"github.com/vishdadhich092004/delivery-management-system/utils"
)

func FilterAgents(agents []domain.Agent, warehouseID uint) []domain.Agent {
	var activeAgents []domain.Agent

	for _, agent := range agents {
		if agent.WarehouseID == warehouseID && agent.Status == "ACTIVE" {
			activeAgents = append(activeAgents, agent)
		}
	}
	return activeAgents
}

func AllocateToAgent(agent domain.Agent, orders []domain.Order, warehouse domain.Warehouse) []domain.OrderAssignment {
	allocatedOrders := []domain.OrderAssignment{}
	totalDistance := 0.0
	totalDuration := 0

	for _, order := range orders {
		distance := utils.HaversineDistance(warehouse.Latitude, warehouse.Longitude, order.Latitude, order.Longitude)
		duration := int(distance * 5) // 5 mins per km

		// 10 hour limit and 100 km distance limit
		if totalDuration+duration > 600 || totalDistance+distance > 100 {
			break
		}

		// Assign the order to the agent
		assignment := domain.OrderAssignment{
			OrderID:           order.ID,
			AgentID:           agent.ID,
			AssignedAt:        time.Now(),
			EstimatedDistance: distance,
			EstimatedDuration: duration,
			Status:            "assigned",
		}
		allocatedOrders = append(allocatedOrders, assignment)

		// Update cumulative stats for the agent
		totalDistance += distance
		totalDuration += duration
	}

	return allocatedOrders
}
