package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/response"
)

func AgentTestController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse(200, "success", gin.H{
		"message": "Hye from Agent",
	}))
}
