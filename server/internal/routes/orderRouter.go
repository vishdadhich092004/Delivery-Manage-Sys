package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/controllers"
)

func OrderRouter(router *gin.RouterGroup) {
	err := router.GET("/", controllers.OrderTestController)
	if err != nil {
		return
	}
}
