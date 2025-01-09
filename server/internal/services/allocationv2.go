package services

import (
	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	helpersv2 "github.com/vishdadhich092004/delivery-management-system/internal/helpersV2"
)

func AllocateOrdersv2(agents []domain.Agent, orders []domain.Order, warehouse domain.Warehouse) []domain.OrderAssignment {
	// Initialize agent capacities
	agentCapacities := helpersv2.InitializeAgentCapacities(agents)

	// Score and sort orders based on distance and basically waiting time
	orderScores := helpersv2.ScoreOrders(orders, warehouse)

	// main fxn allocating orders to agents
	return helpersv2.AssignOrdersToAgents(agentCapacities, orderScores, warehouse)
}
