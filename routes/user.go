package routes

import (
	"minisoccer-backend/controllers"
	"minisoccer-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app fiber.Router) {
	user := app.Group("/user", middlewares.JWTProtected(), middlewares.RoleGuard("user"), middlewares.CheckBlacklist)

	user.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome User"})
	})

	user.Get("/logout", controllers.Logout) // Logout endpoint for users)
}
