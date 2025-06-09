// @title MiniSoccer Backend API
// @version 1.0
// @description Backend for field booking, authentication, and admin control.
// @host localhost:3011
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"log"
	"os"

	_ "minisoccer-backend/docs"
	"minisoccer-backend/seeder"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	"minisoccer-backend/config"
	"minisoccer-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":3011"
	}

	db := config.InitDatabase()
	if db == nil {
		log.Fatal("Failed to initialize database")
	}

	if os.Getenv("SEED") == "true" {
		seeder.SeedUsers()
		seeder.SeedPricing()
		seeder.SeedBookings()
		log.Println("Database seeded with initial data")

	}

	app := fiber.New()

	// ✅ Add root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🟢 Mini Soccer API is running")
	})
	// Register API routes
	api := app.Group("/api")
	routes.RegisterPublicRoutes(api)
	routes.RegisterAdminRoutes(api)
	routes.RegisterUserRoutes(api)
	routes.RegisterBookingRoutes(api)

	// After your route registrations
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Printf("🚀 Server running on http://localhost%s\n", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
