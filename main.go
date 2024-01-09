package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default Gin router
	router := gin.Default()

	// Define a route handler
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin server!",
		})
	})

	// Run the server on port 8080
	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}
