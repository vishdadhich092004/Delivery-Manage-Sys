package routes

import "github.com/gin-gonic/gin"

func V1Router(router *gin.RouterGroup) {
	order := router.Group("/orders")
	OrderRouter(order)
	warehouse := router.Group("/warehouses")
	WarehouseRouter(warehouse)
	agent := router.Group("/agents")
	AgentRouter(agent)
}
