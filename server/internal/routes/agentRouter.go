package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/controllers"
	"gorm.io/gorm"
)

func AgentRouter(router *gin.RouterGroup, db *gorm.DB) {

	agentController := controllers.NewAgentController(db)

	router.GET("/", agentController.AgentTestController)
	router.GET("", agentController.GetAgents)
	router.GET("/:id", agentController.GetAgent)
	router.POST("", agentController.CreateAgent)
	router.PUT("/:id", agentController.UpdateAgent)
	router.POST("/:id/checkin", agentController.AgentCheckin)
	router.POST("/:id/checkout", agentController.AgentCheckout)
	router.GET("/:id/stats", agentController.GetAgentStats)
	router.GET("/:id/assignments", agentController.GetAgentAssignments)
}
