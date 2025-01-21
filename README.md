# Delivery Management System

A robust backend system for managing delivery operations, agent assignments, and warehouse logistics. This system provides efficient order allocation and delivery tracking capabilities.

## Table of Contents
- [Installation](#installation)
- [Models](#models)
- [Features](#features)
- [System Architecture](#system-architecture)
- [Core Functionality](#core-functionality)

## Installation

### Prerequisites
- Go 1.16 or higher
- PostgreSQL database
- Git

### Quick Start
1. Clone the repository:
```bash
git clone https://github.com/vishdadhich092004/Delivery-Manage-Sys.git
```

2. Navigate to the project directory:
```bash
cd Delivery-Manage-Sys
```

3. Navigate to the server directory:
```bash
cd server
```

4. Install dependencies:
```bash
go mod tidy
```

5. Set up environment variables:
```bash
    PORT  = 8080 or any choice
    POSTGRES_URI = YOUR_POSTGRES
```

5. Run the application:
```bash
go run cmd/main.go
```

## Models

### Agent Model
```go
type Agent struct {
    ID          uint      `json:"id"`
    WarehouseID uint      `json:"warehouse_id"`
    Name        string    `json:"name"`
    Phone       string    `json:"phone"`
    VehicleType string    `json:"vehicle_type"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
}
```

### Agent Checkin Model
```go
type AgentCheckin struct {
    ID           uint       `json:"id"`
    AgentID      uint       `json:"agent_id"`
    WarehouseID  uint       `json:"warehouse_id"`
    CheckinTime  time.Time  `json:"checkin_time"`
    CheckoutTime *time.Time `json:"checkout_time"`
    Status       string     `json:"status"`
    CreatedAt    time.Time  `json:"created_at"`
}
```

### Agent Daily Stats Model
```go
type AgentDailyStats struct {
    ID            uint      `json:"id"`
    AgentID       uint      `json:"agent_id"`
    Date          time.Time `json:"date"`
    TotalOrders   int       `json:"total_orders"`
    TotalDistance float64   `json:"total_distance"`
    TotalDuration int       `json:"total_duration"` // in minutes
    Earnings      float64   `json:"earnings"`
}
```

### Order Model
```go
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
```

### Order Assignment Model
```go
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
```

### Warehouse Model
```go
type Warehouse struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Address   string    `json:"address"`
    Latitude  float64   `json:"latitude"`
    Longitude float64   `json:"longitude"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}
```

## Features
- Warehouse management and geocoding
- Delivery agent tracking and status management
- Order allocation with distance and duration optimization
- Agent performance analytics and daily statistics
- Real-time order status tracking
- Agent check-in/check-out system

## System Architecture
The system is built using Go and follows domain-driven design principles. Key components include:
- Domain models for core business entities
- Allocation algorithms for optimizing delivery routes
- Geospatial calculations for distance estimation
- Time-based scheduling and tracking

## Core Functionality

### Order Allocation Algorithm (v1)
The system implements an intelligent order allocation algorithm:

```go
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

        totalDistance += distance
        totalDuration += duration
    }

    return allocatedOrders
}
```

Key features of the algorithm:
- Respects agent work hour limits (10 hours max)
- Considers maximum travel distance (100 km)
- Calculates optimal routes using Haversine distance
- Estimates delivery duration based on distance (5 mins per km)

## Core Functionality

### Order Allocation Algorithm (v2)
The system implements an advanced order allocation algorithm that optimizes delivery assignments using multiple factors:

```go

func AllocateOrdersv2(agents []domain.Agent, orders []domain.Order, warehouse domain.Warehouse) []domain.OrderAssignment {
	// Initialize agent capacities
	agentCapacities := helpersv2.InitializeAgentCapacities(agents)

	// Score and sort orders based on distance and basically waiting time
	orderScores := helpersv2.ScoreOrders(orders, warehouse)

	// main fxn allocating orders to agents
	return helpersv2.AssignOrdersToAgents(agentCapacities, orderScores, warehouse)
}

```

```go
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

```

```go
// this workload score helps to identify the best agent
		workloadScore := float64(agent.RemainingTime) / MAX_WORKING_MINUTES
		orderCountScore := 1.0 / (float64(len(agent.AssignedOrders)) + 1)

		// similar to order, we gave workload a factor of 0.6 and orderCount a factor f .4
		score := workloadScore*0.6 + orderCountScore*0.4


```

Key Components:
1. **Agent Capacity Management**
   - Tracks remaining work hours (10-hour daily limit)
   - Monitors distance capacity (100 km limit)
   - Maintains list of assigned orders per agent

2. **Order Prioritization**
   - Scores orders based on multiple factors:
     - Distance from warehouse (60% weight)
     - Order age/waiting time (40% weight)
   - Lower scores get higher priority

3. **Intelligent Agent Selection**
   - Evaluates agents based on:
     - Current workload capacity
     - Number of existing assignments (load balancing)
     - Vehicle type suitability

Key Features:
- Dynamic workload balancing across available agents
- Prioritizes older orders while considering distance efficiency
- Real-time capacity tracking and adjustment
- Configurable parameters for distance and time limits
  - Maximum working time: 10 hours (600 minutes)
  - Maximum travel distance: 100 km
  - Travel time estimation: 5 minutes per kilometer

Algorithm Flow:
1. Initialize agent capacities with fresh daily limits
2. Score and sort pending orders by priority
3. For each order:
   - Find the most suitable agent based on capacity and workload
   - Calculate estimated distance and duration
   - Update agent's remaining capacity
   - Create order assignment

Improvements over v1:
- More sophisticated order prioritization
- Better load balancing across agents
- Consideration of agent's current workload
- More efficient resource utilization
- Scalable and maintainable code structure



## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
