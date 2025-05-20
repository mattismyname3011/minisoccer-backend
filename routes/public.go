package routes

import (
	"minisoccer-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(app fiber.Router) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
