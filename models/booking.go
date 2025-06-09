package models

import "time"

type Booking struct {
	ID        uint      `gorm:"primaryKey"`
	CourtID   uint      `gorm:"not null"`
	Court     Court     `gorm:"foreignKey:CourtID"`
	UserID    uint      `gorm:"not null"`
	StartTime time.Time `gorm:"not null"` // Date + Time
	EndTime   time.Time `gorm:"not null"`
	Status    string    `gorm:"default:'booked'"` // booked/cancelled/etc
	CreatedAt time.Time
	UpdatedAt time.Time
}
