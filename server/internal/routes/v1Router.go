package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func V1Router(router *gin.RouterGroup, db *gorm.DB) {
	order := router.Group("/orders")
	OrderRouter(order, db)
	warehouse := router.Group("/warehouses")
	WarehouseRouter(warehouse, db)
	agent := router.Group("/agents")
	AgentRouter(agent, db)
}
