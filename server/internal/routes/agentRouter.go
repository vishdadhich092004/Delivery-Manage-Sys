package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/controllers"
)

func AgentRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.AgentTestController)
}
