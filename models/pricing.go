package models

import (
	"time"
)

type Pricing struct {
	ID        uint      `gorm:"primaryKey"`
	CourtID   uint      `gorm:"not null;index"`     // Foreign key reference to Court
	Court     Court     `gorm:"foreignKey:CourtID"` // Belongs to a Court
	StartTime time.Time `gorm:"type:time;not null"` // e.g., "08:00"
	EndTime   time.Time `gorm:"type:time;not null"` // e.g., "10:00"
	Price     float64   `gorm:"not null"`           // e.g., 150000.00
	CreatedAt time.Time
	UpdatedAt time.Time
}
