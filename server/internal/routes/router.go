package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishdadhich092004/delivery-management-system/internal/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	router.Use(middleware.CORSMiddleware())
	// Test route
	router.GET("/hey", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hey from Backend, Its all Good here:~",
		})
	})

	// API routes
	v1 := router.Group("/api/v1")
	V1Router(v1, db)

	return router
}
