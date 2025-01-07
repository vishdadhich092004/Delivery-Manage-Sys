package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/controllers"
)

func AgentRouter(router *gin.RouterGroup) {
	err := router.GET("/", controllers.AgentTestController)
	if err != nil {
		return
	}
}
