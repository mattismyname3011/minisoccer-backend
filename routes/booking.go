package routes

import (
	"minisoccer-backend/controllers"
	"minisoccer-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

// func RegisterBookingRoutes1(r *gin.Engine) {
// 	api := r.Group("/api")
// 	{
// 		api.GET("/pricing", controllers.GetPricing)
// 		api.GET("/availability", controllers.GetAvailability)
// 		api.POST("/bookings", controllers.CreateBooking)
// 	}
// }

func RegisterBookingRoutes(app fiber.Router) {

	public := app.Group("/public", middlewares.CheckPublicSecret)
	public.Get("/pricing", controllers.GetPricing)
	public.Get("/availability", controllers.GetAvailability)
	public.Post("/bookings", controllers.CreateBooking)
}
