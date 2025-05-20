package routes

import (
	"minisoccer-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

// AdminDashboard godoc
// @Summary Admin dashboard
// @Description Returns admin-only dashboard data
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Success 200 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Router /admin/dashboard [get]

func RegisterAdminRoutes(app fiber.Router) {
	admin := app.Group("/admin", middlewares.JWTProtected(), middlewares.RoleGuard("admin"))

	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome Admin"})
	})
}
