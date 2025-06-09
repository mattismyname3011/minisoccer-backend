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
	"minisoccer-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

// func RegisterPublicRoutes(app fiber.Router) {
// 	app.Post("/register", controllers.Register)
// 	app.Post("/login", controllers.Login)
// 	app.Get("/users", controllers.GetUsers)
// }

func RegisterPublicRoutes(app fiber.Router) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	public := app.Group("/public", middlewares.CheckPublicSecret)

	public.Get("/info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is protected public info",
		})
	})

	
}
