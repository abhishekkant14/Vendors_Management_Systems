package controllers

import (
	"net/http"
	"strconv"

	"vendor_Management_system/models"
	"vendor_Management_system/repository"

	"github.com/gin-gonic/gin"
)

func CreateVendor(c *gin.Context) {

	var vendor models.Vendor

	// Read JSON request body
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Save vendor
	if err := repository.CreateVendor(&vendor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create vendor",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Vendor created successfully",
		"vendor":  vendor,
	})
}

func GetVendors(c *gin.Context) {

	var vendors []models.Vendor

	if err := repository.GetAllVendors(&vendors); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch vendors",
		})
		return
	}

	c.JSON(http.StatusOK, vendors)
}

func GetVendorByID(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid vendor ID",
		})
		return
	}

	var vendor models.Vendor

	if err := repository.GetVendorByID(&vendor, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Vendor not found",
		})
		return
	}

	c.JSON(http.StatusOK, vendor)
}

func UpdateVendor(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid vendor ID",
		})
		return
	}

	var vendor models.Vendor

	if err := repository.GetVendorByID(&vendor, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Vendor not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	vendor.ID = uint(id)

	if err := repository.UpdateVendor(&vendor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update vendor",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vendor updated successfully",
		"vendor":  vendor,
	})
}

func DeleteVendor(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid vendor ID",
		})
		return
	}

	var vendor models.Vendor

	if err := repository.GetVendorByID(&vendor, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Vendor not found",
		})
		return
	}

	if err := repository.DeleteVendor(&vendor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete vendor",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vendor deleted successfully",
	})
}
