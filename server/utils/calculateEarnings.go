package utils

import "github.com/vishdadhich092004/delivery-management-system/pkg/constants"

func CalculateEarnings(orderCount int) float64 {
	basePay := constants.MinGuarantee
	if orderCount > constants.Tier1OrderCount {
		return basePay + float64(orderCount)*constants.Tier1PayPerOrder
	} else if orderCount > constants.Tier2OrderCount {
		return basePay + float64(orderCount)*constants.Tier2PayPerOrder
	}
	return basePay
}
