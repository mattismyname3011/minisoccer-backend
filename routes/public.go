// Login godoc
// @Summary Logs in a user
// @Description Authenticates a user and returns JWT token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body map[string]string true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]

package routes

import (
	"minisoccer-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(app fiber.Router) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
