package controllers

import (
	"time"

	"os"

	"minisoccer-backend/config"
	"minisoccer-backend/models"
	"minisoccer-backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	hash, err := utils.HashPassword(body.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	user := models.User{
		Email:        body.Email,
		PasswordHash: hash,
		Role:         "user", // default role
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "User already exists?")
	}

	return c.JSON(fiber.Map{"message": "User registered"})
}

func Login(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	var user models.User
	if err := config.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	if !utils.CheckPasswordHash(body.Password, user.PasswordHash) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Generate JWT
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // expires in 72 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not sign token")
	}

	return c.JSON(fiber.Map{"token": signedToken})
}
