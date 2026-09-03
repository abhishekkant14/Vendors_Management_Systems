package routes

import (
	"vendor_Management_system/controllers"

	"github.com/gin-gonic/gin"
)

func VendorRoutes(router *gin.Engine) {

	//create vendor
	router.POST("/vendors", controllers.CreateVendor)

	//get all vendors
	router.GET("/vendors", controllers.GetVendors)

	//get vendor by id
	router.GET("/vendors/:id", controllers.GetVendorByID)

	//update vendor
	router.PUT("/vendors/:id", controllers.UpdateVendor)

	//delete vendor
	router.DELETE("/vendors/:id", controllers.DeleteVendor)

}
