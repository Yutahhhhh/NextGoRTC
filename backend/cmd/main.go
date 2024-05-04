package main

import (
	"NextGoRTC/internal/infrastructures/gorm"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	gorm.Connect()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}