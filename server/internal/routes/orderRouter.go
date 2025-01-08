package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/controllers"
	"gorm.io/gorm"
)

func OrderRouter(router *gin.RouterGroup, db *gorm.DB) {

	orderController := controllers.NewOrderController(db)
	router.GET("/", orderController.OrderTestController)
	router.GET("", orderController.GetOrders)
	router.GET("/:id", orderController.GetOrder)
	router.POST("", orderController.CreateOrder)
	router.PUT("/:id", orderController.UpdateOrder)
	router.POST("/:id/complete", orderController.CompleteOrder)
}
