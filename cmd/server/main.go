package main

import (
	"github.com/dmytrosydor/PriceTracker/config"
	// "github.com/dmytrosydor/PriceTracker/pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDataBase()
	// config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.PriceHistory{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run(":8080")
}
