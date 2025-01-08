package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/controllers"
	"gorm.io/gorm"
)

func WarehouseRouter(router *gin.RouterGroup, db *gorm.DB) {

	warehouseController := controllers.NewWarehouseController(db)

	router.GET("/test", warehouseController.WarehouseTestController)
	router.GET("/", warehouseController.GetWarehouses)
	router.GET("/:id", warehouseController.GetWarehouse)
	router.POST("", warehouseController.CreateWarehouse)
	router.PUT("/:id", warehouseController.UpdateWarehouse)
	router.GET("/:id/agents", warehouseController.GetWarehouseAgents)
	router.GET("/:id/orders", warehouseController.GetWarehouseOrders)
	router.POST("/:id/orders/assign", warehouseController.AllocateOrders)
}
