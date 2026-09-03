package main

import (
	"fmt"
	"log"

	"vendor_Management_system/config"
	"vendor_Management_system/models"
	"vendor_Management_system/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Connecting to MySQL..")

	// Database connection
	config.ConnectDB()

	// Database migration
	err := config.DB.AutoMigrate(&models.Vendor{})

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Vendor table migrated successfully")

	// Create Gin router
	router := gin.Default()

	// Register vendor routes
	routes.VendorRoutes(router)

	// Home API
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Vendor Management System API",
		})
	})

	// Start server
	fmt.Println("Server running on port 9091")

	err = router.Run(":9091")

	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
