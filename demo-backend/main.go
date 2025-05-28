package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/wolewd/database-docker-testing/config"
	"github.com/wolewd/database-docker-testing/db"
	"github.com/wolewd/database-docker-testing/handlers"
)

func main() {
	config.LoadEnv()
	db.InitDB()

	err := gin.Default()

	// Routes
	err.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	err.GET("/users", handlers.GetUsers(db.DB_A))
	err.POST("/users", handlers.CreateUser(db.DB_A))

	err.GET("/products", handlers.GetProducts(db.DB_B))
	err.POST("/products", handlers.CreateProduct(db.DB_B))

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	err.Run(":" + port)
}
