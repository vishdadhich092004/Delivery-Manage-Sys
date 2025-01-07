package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vishdadhich092004/delivery-management-system/config"
)

func main() {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hey from Backend",
		})
	})

	db, err := config.NewDB(os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatalf("Error Initialising the DataBase : %w", err)
	}
	defer db.Close()
	r.Run(":" + os.Getenv("PORT"))

}
