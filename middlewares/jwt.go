package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JWTProtected() func(*fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(secret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}
