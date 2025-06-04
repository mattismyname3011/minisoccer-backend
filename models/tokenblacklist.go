package models

import "time"

type TokenBlacklist struct {
	Token     string    `gorm:"primaryKey;not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
