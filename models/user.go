// models/user.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"default:user"` // user or admin
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// type User struct {
// 	ID           uint           `gorm:"primaryKey"`
// 	Name         string         `gorm:"not null"`
// 	Email        string         `gorm:"uniqueIndex;not null"`
// 	PasswordHash string         `gorm:"not null"`
// 	PhoneNumber  string         `gorm:"not null"`
// 	Role         string         `gorm:"default:user"` // user or admin
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// 	DeletedAt    gorm.DeletedAt `gorm:"index"`
// }
