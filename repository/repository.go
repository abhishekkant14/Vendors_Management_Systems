package repository

import (
	"Vendor_Management_System/config"
	"Vendor_Management_System/models"
)

func CreateVendor(vendor *models.Vendor) error {

	return config.DB.Create(vendor).Error

}

func GetAllVendor(vendors *[]models.Vendor) error {
	return config.DB.Find(vendors).Error

}

func GetVendorByID(vendor *models.Vendor, id uint) error {
	return config.DB.First(vendor, id).Error

}

func UpdateVendor(vendor *models.Vendor) error {

	return config.DB.Save(vendor).Error
}

func DeleteVendor(vendor *models.Vendor) error {

	return config.DB.Delete(vendor).Error
}
