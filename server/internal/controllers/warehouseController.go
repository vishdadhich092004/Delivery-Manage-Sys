package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"github.com/vishdadhich092004/delivery-management-system/internal/services"
	"github.com/vishdadhich092004/delivery-management-system/pkg/response"
	"gorm.io/gorm"
)

type WarehouseController struct {
	db *gorm.DB
}

func NewWarehouseController(db *gorm.DB) *WarehouseController {
	return &WarehouseController{db: db}
}

func (wc *WarehouseController) WarehouseTestController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hey from Warehoeezay",
	}))
}

// Fetch all warehouses
func (wc *WarehouseController) GetWarehouses(ctx *gin.Context) {
	var warehouses []domain.Warehouse
	if err := wc.db.Find(&warehouses).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to fetch warehouses",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", warehouses))
}

// Add a new warehouse
func (wc *WarehouseController) CreateWarehouse(ctx *gin.Context) {
	var warehouse domain.Warehouse
	if err := ctx.ShouldBindJSON(&warehouse); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse(400, "error", gin.H{
			"error": "Invalid warehouse data",
		}))
		return
	}
	if err := wc.db.Create(&warehouse).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to create warehouse",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", warehouse))
}

// Fetch a single warehouse by ID
func (wc *WarehouseController) GetWarehouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var warehouse domain.Warehouse
	if err := wc.db.First(&warehouse, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Warehouse not found",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", warehouse))
}

// Update warehouse details
func (wc *WarehouseController) UpdateWarehouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var warehouse domain.Warehouse
	if err := wc.db.First(&warehouse, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Warehouse not found",
		}))
		return
	}
	if err := ctx.ShouldBindJSON(&warehouse); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse(400, "error", gin.H{
			"error": "Invalid warehouse data",
		}))
		return
	}
	if err := wc.db.Save(&warehouse).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to update warehouse",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", warehouse))
}

// Fetch all agents for a specific warehouse
func (wc *WarehouseController) GetWarehouseAgents(ctx *gin.Context) {
	id := ctx.Param("id")
	var agents []domain.Agent
	if err := wc.db.Where("warehouse_id = ?", id).Find(&agents).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to fetch agents for warehouse",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", agents))
}

// Fetch all orders for a specific warehouse
func (wc *WarehouseController) GetWarehouseOrders(ctx *gin.Context) {
	id := ctx.Param("id")
	var orders []domain.Order
	if err := wc.db.Where("warehouse_id = ?", id).Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to fetch orders for warehouse",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", orders))
}

// Allocate orders for a warehouse
func (wc *WarehouseController) AllocateOrders(ctx *gin.Context) {
	warehouseID := ctx.Param("id")

	// Fetch the warehouse
	var warehouse domain.Warehouse
	if err := wc.db.First(&warehouse, "id = ?", warehouseID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Warehouse not found",
		}))
		return
	}

	// Fetch agents for the warehouse
	var agents []domain.Agent
	if err := wc.db.Where("warehouse_id = ?", warehouseID).Find(&agents).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to fetch agents",
		}))
		return
	}

	// Fetch pending orders for the warehouse
	var orders []domain.Order
	if err := wc.db.Where("warehouse_id = ? AND status = ?", warehouseID, "PENDING").Order("created_at").Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to fetch orders",
		}))
		return
	}

	if len(orders) == 0 {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse(400, "error", gin.H{
			"error": "No pending orders available",
		}))
		return
	}

	// Allocate orders using the service logic
	allocations := services.AllocateOrders(agents, orders, warehouse)

	// Save the allocations to the database
	for _, allocation := range allocations {
		if err := wc.db.Create(&allocation).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
				"error": "Failed to save order assignments",
			}))
			return
		}
	}

	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"allocations": allocations,
	}))
}
