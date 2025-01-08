package helpers

import (
	"sort"

	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"github.com/vishdadhich092004/delivery-management-system/utils"
)

func SortOrdersByDistance(orders []domain.Order, warehouse domain.Warehouse) []domain.Order {

	type orderWithDistance struct {
		order    domain.Order
		distance float64
	}
	var ordersWithDistances []orderWithDistance

	for _, order := range orders {
		distance := utils.HaversineDistance(warehouse.Latitude, warehouse.Longitude, order.Latitude, order.Longitude)
		ordersWithDistances = append(ordersWithDistances, orderWithDistance{order: order, distance: distance})
	}
	sort.Slice(ordersWithDistances, func(i, j int) bool {
		return ordersWithDistances[i].distance < ordersWithDistances[j].distance
	})

	sortedOrders := make([]domain.Order, len(ordersWithDistances))
	for i, od := range ordersWithDistances {
		sortedOrders[i] = od.order
	}

	return sortedOrders
}
