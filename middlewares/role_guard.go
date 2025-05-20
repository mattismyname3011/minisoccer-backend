package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RoleGuard(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		if user == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unauthorized"})
		}

		claims, ok := user.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		role, ok := claims["role"].(string)
		if !ok || role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Access denied. Insufficient role",
			})
		}

		return c.Next()
	}
}
