package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/domain"
	"github.com/vishdadhich092004/delivery-management-system/pkg/response"
	"gorm.io/gorm"
)

type OrderController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{db: db}
}

// Test Controller
func (oc *OrderController) OrderTestController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "hey from Orders",
	}))
}

// Fetch all orders
func (oc *OrderController) GetOrders(ctx *gin.Context) {
	var orders []domain.Order
	if err := oc.db.Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to fetch orders",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", orders))
}

// Fetch a specific order by ID
func (oc *OrderController) GetOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var order domain.Order
	if err := oc.db.First(&order, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Order not found",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", order))
}

// Create a new order
func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var order domain.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse(400, "error", gin.H{
			"error": "Invalid input",
		}))
		return
	}
	if err := oc.db.Create(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to create order",
		}))
		return
	}
	ctx.JSON(http.StatusCreated, response.ApiResponse(201, "success", order))
}

// Update an order by ID
func (oc *OrderController) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var order domain.Order
	if err := oc.db.First(&order, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Order not found",
		}))
		return
	}
	var updatedData domain.Order
	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse(400, "error", gin.H{
			"error": "Invalid input",
		}))
		return
	}
	if err := oc.db.Model(&order).Updates(updatedData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to update order",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", order))
}

// Mark an order as completed
func (oc *OrderController) CompleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var order domain.Order
	if err := oc.db.First(&order, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, response.ApiResponse(404, "error", gin.H{
			"error": "Order not found",
		}))
		return
	}
	order.Status = "completed"
	if err := oc.db.Save(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse(500, "error", gin.H{
			"error": "Failed to complete order",
		}))
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Order marked as completed",
	}))
}
