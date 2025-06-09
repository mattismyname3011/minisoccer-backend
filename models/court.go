package models

import (
	"time"
)

type Court struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	Location  string    `gorm:"not null"`
	Category  string    `gorm:"type:varchar(50)"` // e.g., synthetic, grass, indoor
	IsActive  bool      `gorm:"default:true"`
	Pricings  []Pricing `gorm:"foreignKey:CourtID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
