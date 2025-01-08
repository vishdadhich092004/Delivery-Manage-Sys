package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/response"
	"gorm.io/gorm"
)

type AgentController struct {
	db *gorm.DB
}

func NewAgentController(db *gorm.DB) *AgentController {
	return &AgentController{db: db}
}

func (ac *AgentController) AgentTestController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}
func (ac *AgentController) GetAgents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}
func (ac *AgentController) GetAgent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}
func (ac *AgentController) CreateAgent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}

func (ac *AgentController) UpdateAgent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}

func (ac *AgentController) AgentCheckin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}

func (ac *AgentController) AgentCheckout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}

func (ac *AgentController) GetAgentStats(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}

func (ac *AgentController) GetAgentAssignments(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}
