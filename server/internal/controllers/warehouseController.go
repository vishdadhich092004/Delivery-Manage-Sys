package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/response"
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

func (wc *WarehouseController) GetWarehouses(ctx *gin.Context) {

}
func (wc *WarehouseController) CreateWarehouse(ctx *gin.Context) {

}
func (wc *WarehouseController) GetWarehouse(ctx *gin.Context) {

}
func (wc *WarehouseController) UpdateWarehouse(ctx *gin.Context) {

}
func (wc *WarehouseController) GetWarehouseAgents(ctx *gin.Context) {

}
func (wc *WarehouseController) GetWarehouseOrders(ctx *gin.Context) {

}
