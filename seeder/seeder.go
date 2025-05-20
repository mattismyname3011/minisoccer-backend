package seeder

import (
	"fmt"

	"minisoccer-backend/config"
	"minisoccer-backend/models"
	"minisoccer-backend/utils"
)

func SeedUsers() {
	adminPass, _ := utils.HashPassword("admin123")
	userPass, _ := utils.HashPassword("user123")

	users := []models.User{
		{Email: "admin@example.com", PasswordHash: adminPass, Role: "admin"},
		{Email: "user@example.com", PasswordHash: userPass, Role: "user"},
	}

	for _, u := range users {
		if err := config.DB.Create(&u).Error; err != nil {
			fmt.Println("Failed to create user:", u.Email, "-", err)
		} else {
			fmt.Println("User created:", u.Email)
		}
	}
}
