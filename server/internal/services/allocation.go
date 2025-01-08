package services

import (
	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"github.com/vishdadhich092004/delivery-management-system/internal/helpers"
)

func AllocateOrders(agents []domain.Agent, orders []domain.Order, warehouse domain.Warehouse) []domain.OrderAssignment {

	// filter  out the active agents
	activeAgents := helpers.FilterAgents(agents, warehouse.ID)

	//sort the orders on the basis of distnace
	sortedOrders := helpers.SortOrdersByDistance(orders, warehouse)

	// allocating

	assignments := []domain.OrderAssignment{}

	for _, agent := range activeAgents {
		agentAssignments := helpers.AllocateToAgent(agent, sortedOrders, warehouse)

		assignments = append(assignments, agentAssignments...)
	}
	return assignments
}
