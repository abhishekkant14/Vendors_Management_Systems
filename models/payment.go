package models

import (
	"time"
)

type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	VVendorID     uint      `json:"vendor_id"`
	WorkID        uint      `json:"work_id"`
	Amount        float64   `json:"amount"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
	TransactionID string    `json:"transaction_id"`
	Remarks       string    `json:"remarks"`
	CreatedAt     time.Time `json:"created_at"`
}
