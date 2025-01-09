package helpersv2

import (
	"time"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"github.com/vishdadhich092004/delivery-management-system/utils"
)

const (
	MAX_WORKING_MINUTES = 600 // 10 hours
	MAX_DISTANCE_KM     = 100
	MINUTES_PER_KM      = 5
)

type AgentCapacity struct {
	Agent          domain.Agent
	RemainingTime  int     // in minutes
	RemainingDist  float64 // in km
	AssignedOrders []domain.OrderAssignment
}

func AssignOrdersToAgents(
	agentCapacities []AgentCapacity,
	orderScores []OrderScore,
	warehouse domain.Warehouse,
) []domain.OrderAssignment {
	allAssignments := make([]domain.OrderAssignment, 0)

	// For each order we find best agent
	for _, orderScore := range orderScores {
		order := orderScore.Order
		bestAgent := findBestAgentForOrder(order, agentCapacities, warehouse)

		if bestAgent == nil {
			continue // No agent can handle this order
		}

		// Calculate metrics
		distance := utils.HaversineDistance(
			warehouse.Latitude,
			warehouse.Longitude,
			order.Latitude,
			order.Longitude,
		)
		duration := int(distance * MINUTES_PER_KM)

		// Create assignment
		assignment := domain.OrderAssignment{
			OrderID:           order.ID,
			AgentID:           bestAgent.Agent.ID,
			AssignedAt:        time.Now(),
			EstimatedDistance: distance,
			EstimatedDuration: duration,
			Status:            "ASSIGNED",
		}

		// Update agent capacity
		bestAgent.RemainingTime -= duration
		bestAgent.RemainingDist -= distance
		bestAgent.AssignedOrders = append(bestAgent.AssignedOrders, assignment)

		allAssignments = append(allAssignments, assignment)
	}

	return allAssignments
}

func InitializeAgentCapacities(agents []domain.Agent) []AgentCapacity {
	capacities := make([]AgentCapacity, 0, len(agents))

	for _, agent := range agents {
		// Only include active agents
		if agent.Status == "ACTIVE" {
			capacity := AgentCapacity{
				Agent:          agent,
				RemainingTime:  MAX_WORKING_MINUTES,
				RemainingDist:  MAX_DISTANCE_KM,
				AssignedOrders: make([]domain.OrderAssignment, 0),
			}
			capacities = append(capacities, capacity)
		}
	}

	return capacities
}

func findBestAgentForOrder(
	order domain.Order,
	agentCapacities []AgentCapacity,
	warehouse domain.Warehouse,
) *AgentCapacity {
	var bestAgent *AgentCapacity
	var bestScore float64 = -1

	orderDistance := utils.HaversineDistance(
		warehouse.Latitude,
		warehouse.Longitude,
		order.Latitude,
		order.Longitude,
	)
	orderDuration := int(orderDistance * MINUTES_PER_KM)

	for i := range agentCapacities {
		agent := &agentCapacities[i]

		// Check if agent has capacity
		if agent.RemainingTime < orderDuration || agent.RemainingDist < orderDistance {
			continue
		}

		// this workload score helps to identify the best agent
		workloadScore := float64(agent.RemainingTime) / MAX_WORKING_MINUTES
		orderCountScore := 1.0 / (float64(len(agent.AssignedOrders)) + 1)

		// similar to order, we gave workload a factor of 0.6 and orderCount a factor f .4
		score := workloadScore*0.6 + orderCountScore*0.4

		// if nothing is allocated
		if bestScore == -1 || score > bestScore {
			bestScore = score
			bestAgent = agent
		}
	}

	return bestAgent
}
