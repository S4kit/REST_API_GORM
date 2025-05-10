package Models

import (
	"time"
)

type Product struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"product_name"`
	SerialNumber string `json:"serial_number"`
}
