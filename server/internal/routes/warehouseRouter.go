package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/controllers"
)

func WarehouseRouter(router *gin.RouterGroup) {
	err := router.GET("/", controllers.WarehouseTestController)
	if err != nil {
		return
	}
}
