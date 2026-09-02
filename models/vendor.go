package models

import "time"

type Vendor struct {
	ID        uint      `gorm:"primaryKey" json:"ID"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	Phone     string    `json:"Phone"`
	Address   string    `json:"Address"`
	Company   string    `json:"Company"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
