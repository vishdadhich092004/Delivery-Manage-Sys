package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"github.com/vishdadhich092004/delivery-management-system/pkg/response"
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
	var agents []domain.Agent
	if err := ac.db.Find(&agents).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to fetch Agents",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", agents))
}
func (ac *AgentController) GetAgent(ctx *gin.Context) {
	id := ctx.Param("id")
	var agent domain.Agent
	if err := ac.db.Find(&agent, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Agent Not Found",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", agent))
}

// Create a new agent
func (ac *AgentController) CreateAgent(ctx *gin.Context) {
	var agent domain.Agent
	if err := ctx.ShouldBindJSON(&agent); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse(400, "error", gin.H{
			"error": "Invalid input",
		}))
		return
	}
	if err := ac.db.Create(&agent).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to create agent",
		}))
		return
	}
	ctx.JSON(http.StatusCreated, response.ApiResponse(201, "success", agent))
}

// Update an agent by ID
func (ac *AgentController) UpdateAgent(ctx *gin.Context) {
	id := ctx.Param("id")
	var agent domain.Agent
	if err := ac.db.First(&agent, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Agent not found",
		}))
		return
	}
	var updatedData domain.Agent
	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse(400, "error", gin.H{
			"error": "Invalid input",
		}))
		return
	}
	if err := ac.db.Model(&agent).Updates(updatedData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to update agent",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", agent))
}

// Agent check-in
func (ac *AgentController) AgentCheckin(ctx *gin.Context) {
	id := ctx.Param("id")
	var agent domain.Agent
	if err := ac.db.First(&agent, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Agent not found",
		}))
		return
	}

	if err := ac.db.Model(&agent).Update("status", "checked-in").Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to check in",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Agent checked in successfully",
	}))
}

// Agent check-out
func (ac *AgentController) AgentCheckout(ctx *gin.Context) {
	id := ctx.Param("id")
	var agent domain.Agent
	if err := ac.db.First(&agent, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Agent not found",
		}))
		return
	}
	if err := ac.db.Model(&agent).Update("status", "checked-out").Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to check out",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Agent checked out successfully",
	}))
}

func (ac *AgentController) GetAgentStats(ctx *gin.Context) {
	id := ctx.Param("id")
	var stats domain.AgentDailyStats
	// Replace this with the actual logic to fetch stats
	if err := ac.db.First(&stats, "agent_id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Stats not found",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", stats))
}

func (ac *AgentController) GetAgentAssignments(ctx *gin.Context) {
	id := ctx.Param("id")
	var assignments []domain.OrderAssignment
	if err := ac.db.Where("agent_id = ?", id).Find(&assignments).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to fetch assignments",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", assignments))
}
