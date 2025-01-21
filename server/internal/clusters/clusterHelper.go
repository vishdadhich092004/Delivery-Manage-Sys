package clusters

import (
	"math"
	"sort"
	"time"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"github.com/vishdadhich092004/delivery-management-system/utils"
)

type AgentCapacity struct {
	Agent          domain.Agent
	RemainingTime  int     // in minutes
	RemainingDist  float64 // in km
	AssignedOrders []domain.OrderAssignment
}

const (
	MAX_WORKING_MINUTES = 600 // 10 hours
	MAX_DISTANCE_KM     = 100
	MINUTES_PER_KM      = 5
	CLUSTER_RADIUS_KM   = 5 // Maximum radius for a cluster
	MIN_CLUSTER_SIZE    = 2 // Minimum orders in a cluster
)

type OrderCluster struct {
	Orders    []domain.Order
	Centroid  Point
	TotalTime int     // Total estimated delivery time for cluster
	Distance  float64 // Total distance to cover cluster
}

type Point struct {
	Latitude  float64
	Longitude float64
}

func ClusterOrders(orders []domain.Order, warehouse domain.Warehouse) []OrderCluster {
	var clusters []OrderCluster
	unassignedOrders := make([]domain.Order, len(orders))
	copy(unassignedOrders, orders)

	for len(unassignedOrders) > 0 {
		// Take the first unassigned order as a potential cluster center
		centerOrder := unassignedOrders[0]
		cluster := OrderCluster{
			Orders: []domain.Order{centerOrder},
			Centroid: Point{
				Latitude:  centerOrder.Latitude,
				Longitude: centerOrder.Longitude,
			},
		}

		// Find nearby orders within CLUSTER_RADIUS_KM
		remainingOrders := []domain.Order{}
		for i := 1; i < len(unassignedOrders); i++ {
			order := unassignedOrders[i]
			distance := utils.HaversineDistance(
				centerOrder.Latitude,
				centerOrder.Longitude,
				order.Latitude,
				order.Longitude,
			)

			if distance <= CLUSTER_RADIUS_KM {
				cluster.Orders = append(cluster.Orders, order)
			} else {
				remainingOrders = append(remainingOrders, order)
			}
		}

		// Only keep clusters with minimum size
		if len(cluster.Orders) >= MIN_CLUSTER_SIZE {
			// Update cluster centroid and metrics
			updateClusterMetrics(&cluster, warehouse)
			clusters = append(clusters, cluster)
		} else {
			// Add small cluster orders back to remaining
			remainingOrders = append(remainingOrders, cluster.Orders...)
		}

		unassignedOrders = remainingOrders
	}

	// Sort clusters by size (descending) and then by total distance (ascending)
	sort.Slice(clusters, func(i, j int) bool {
		if len(clusters[i].Orders) == len(clusters[j].Orders) {
			return clusters[i].Distance < clusters[j].Distance
		}
		return len(clusters[i].Orders) > len(clusters[j].Orders)
	})

	return clusters
}

func updateClusterMetrics(cluster *OrderCluster, warehouse domain.Warehouse) {
	// Calculate centroid
	var sumLat, sumLng float64
	for _, order := range cluster.Orders {
		sumLat += order.Latitude
		sumLng += order.Longitude
	}
	cluster.Centroid = Point{
		Latitude:  sumLat / float64(len(cluster.Orders)),
		Longitude: sumLng / float64(len(cluster.Orders)),
	}

	// Calculate total distance and time
	// Start from warehouse to first order
	prevPoint := Point{Latitude: warehouse.Latitude, Longitude: warehouse.Longitude}
	totalDistance := 0.0

	// Sort orders by distance from previous point
	orders := make([]domain.Order, len(cluster.Orders))
	copy(orders, cluster.Orders)

	for i := 0; i < len(orders); i++ {
		minDist := math.MaxFloat64
		minIdx := i

		// Find nearest unvisited order
		for j := i; j < len(orders); j++ {
			dist := utils.HaversineDistance(
				prevPoint.Latitude,
				prevPoint.Longitude,
				orders[j].Latitude,
				orders[j].Longitude,
			)
			if dist < minDist {
				minDist = dist
				minIdx = j
			}
		}

		// Swap to put nearest order next
		orders[i], orders[minIdx] = orders[minIdx], orders[i]
		totalDistance += minDist
		prevPoint = Point{Latitude: orders[i].Latitude, Longitude: orders[i].Longitude}
	}

	// Add return distance to warehouse
	totalDistance += utils.HaversineDistance(
		prevPoint.Latitude,
		prevPoint.Longitude,
		warehouse.Latitude,
		warehouse.Longitude,
	)

	cluster.Distance = totalDistance
	cluster.TotalTime = int(totalDistance * MINUTES_PER_KM)
}

func AssignClustersToAgents(
	clusters []OrderCluster,
	agentCapacities []AgentCapacity,
	warehouse domain.Warehouse,
) []domain.OrderAssignment {
	var allAssignments []domain.OrderAssignment

	// For each cluster, find the best agent
	for _, cluster := range clusters {
		bestAgent := findBestAgentForCluster(cluster, agentCapacities)
		if bestAgent == nil {
			continue
		}

		// Create assignments for all orders in the cluster
		for _, order := range cluster.Orders {
			assignment := domain.OrderAssignment{
				OrderID:           order.ID,
				AgentID:           bestAgent.Agent.ID,
				AssignedAt:        time.Now(),
				EstimatedDistance: cluster.Distance / float64(len(cluster.Orders)),
				EstimatedDuration: cluster.TotalTime / len(cluster.Orders),
				Status:            "ASSIGNED",
			}

			bestAgent.AssignedOrders = append(bestAgent.AssignedOrders, assignment)
			allAssignments = append(allAssignments, assignment)
		}

		// Update agent capacity
		bestAgent.RemainingTime -= cluster.TotalTime
		bestAgent.RemainingDist -= cluster.Distance
	}

	return allAssignments
}

func findBestAgentForCluster(cluster OrderCluster, agentCapacities []AgentCapacity) *AgentCapacity {
	var bestAgent *AgentCapacity
	var bestScore float64 = -1

	for i := range agentCapacities {
		agent := &agentCapacities[i]

		// Check if agent has capacity for the entire cluster
		if agent.RemainingTime < cluster.TotalTime || agent.RemainingDist < cluster.Distance {
			continue
		}

		// Score based on remaining capacity and current workload
		workloadScore := float64(agent.RemainingTime) / MAX_WORKING_MINUTES
		orderCountScore := 1.0 / (float64(len(agent.AssignedOrders)) + 1)
		score := workloadScore*0.6 + orderCountScore*0.4

		if bestScore == -1 || score > bestScore {
			bestScore = score
			bestAgent = agent
		}
	}

	return bestAgent
}
