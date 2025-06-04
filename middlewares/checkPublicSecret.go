package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func CheckPublicSecret(c *fiber.Ctx) error {
	secretFromHeader := c.Get("X-Public-Key") // you can also use query params if needed
	expectedSecret := os.Getenv("PUBLIC_SECRET_KEY")

	if secretFromHeader != expectedSecret {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: invalid public secret key",
		})
	}

	return c.Next()
}
