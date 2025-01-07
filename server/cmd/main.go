package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hey from Backend",
		})
	})

	r.Run(":" + os.Getenv("PORT"))

}
