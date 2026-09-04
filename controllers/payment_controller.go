package controllers

import (
	"net/http"
	"strconv"

	"vendor_Management_system/models"
	"vendor_Management_system/repository"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {

	var payment models.Payment

	// Read JSON request body
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Validate payment amount
	if payment.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Payment amount must be greater than zero",
		})
		return
	}

	// Save payment into database
	if err := repository.CreatePayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create payment",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Payment created successfully",
		"payment": payment,
	})
}

func GetPayments(c *gin.Context) {

	var payments []models.Payment

	// Get all payments
	if err := repository.GetAllPayments(&payments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch payments",
		})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func GetPaymentByID(c *gin.Context) {

	// Get payment ID from URL
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid payment ID",
		})
		return
	}

	var payment models.Payment

	// Find payment by ID
	if err := repository.GetPaymentByID(
		&payment,
		uint(id),
	); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Payment not found",
		})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func GetPaymentsByVendorID(c *gin.Context) {

	// Get vendor ID from URL
	vendorID, err := strconv.ParseUint(
		c.Param("vendor_id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid vendor ID",
		})
		return
	}

	var payments []models.Payment

	// Get payments for specific vendor
	if err := repository.GetPaymentsByVendorID(
		&payments,
		uint(vendorID),
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch vendor payments",
		})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func DeletePayment(c *gin.Context) {

	// Get payment ID from URL
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid payment ID",
		})
		return
	}

	var payment models.Payment

	// Find payment before deleting
	if err := repository.GetPaymentByID(
		&payment,
		uint(id),
	); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Payment not found",
		})
		return
	}

	// Delete payment
	if err := repository.DeletePayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete payment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment deleted successfully",
	})
}
