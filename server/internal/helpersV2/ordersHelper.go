package helpersv2

import (
	"sort"
	"time"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"github.com/vishdadhich092004/delivery-management-system/utils"
)

type OrderScore struct {
	Order domain.Order
	Score float64
}

func ScoreOrders(orders []domain.Order, warehouse domain.Warehouse) []OrderScore {
	scoredOrders := make([]OrderScore, 0, len(orders))

	for _, order := range orders {
		// Calculate distance
		distance := utils.HaversineDistance(
			warehouse.Latitude,
			warehouse.Longitude,
			order.Latitude,
			order.Longitude,
		)

		// Calculate waiting time in hours for an order!
		waitingTime := time.Since(order.CreatedAt).Hours()

		// basically giving distance a factor of .6 and waitign time a factor of 0.4
		score := distance*0.6 - waitingTime*0.4

		scoredOrders = append(scoredOrders, OrderScore{
			Order: order,
			Score: score,
		})
	}

	// Sort by score (ascending)
	sort.Slice(scoredOrders, func(i, j int) bool {
		return scoredOrders[i].Score < scoredOrders[j].Score
	})

	return scoredOrders
}
