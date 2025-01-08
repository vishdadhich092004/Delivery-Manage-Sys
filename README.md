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

### Order Allocation Algorithm
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


## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
