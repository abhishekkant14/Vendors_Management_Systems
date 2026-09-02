package main

import (
	"fmt"
	"log"
	"vendor_Management_system/config"
	"vendor_Management_system/models"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Connecting to MySQl..")

	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.Vendor{})

	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("Vendor table migrated successfully")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Vendor Management System API",
		})
	})

	router.Run(":9091")
}
