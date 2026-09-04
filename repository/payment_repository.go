package repository

import (
	"vendor_Management_system/config"
	"vendor_Management_system/models"
)

func CreatePayment(payment *models.Payment) error {
	return config.DB.Create(payment).Error
}

func GetAllPayments(payments *[]models.Payment) error {
	return config.DB.Find(payments).Error
}

func GetPaymentByID(payment *models.Payment, id uint) error {
	return config.DB.First(payment, id).Error
}

func GetPaymentsByVendorID(payments *[]models.Payment, vendorID uint) error {
	return config.DB.Where("vendor_id = ?", vendorID).Find(payments).Error
}

func DeletePayment(payment *models.Payment) error {
	return config.DB.Delete(payment).Error
}
