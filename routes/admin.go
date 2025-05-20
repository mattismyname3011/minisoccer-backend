package routes

import (
	"minisoccer-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterAdminRoutes(app fiber.Router) {
	admin := app.Group("/admin", middlewares.JWTProtected())

	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome Admin"})
	})
}
