// package main

// import (
// 	"fmt"
// 	"log"
// 	"minisoccer-backend/config"
// 	"net/http"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// Define the server address and port
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	addr := ":" + os.Getenv("PORT")
// 	if addr == ":" {
// 		addr = ":3011"
// 	}

// 	// Initialize the database
// 	db := config.InitDatabase()
// 	if db == nil {
// 		log.Fatal("Failed to initialize database")
// 	}
// 	// Define a simple handler
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Welcome to Mini Soccer Backend!")
// 	})

// 	// seeder.SeedUsers() // 👈 Call it just once, then remove or comment out

// 	// Start the server
// 	log.Printf("Server is running on http://localhost%s\n", addr)
// 	if err := http.ListenAndServe(addr, nil); err != nil {
// 		log.Fatalf("Could not start server: %s\n", err)
// 	}
// }

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()

	// Register API routes
	api := app.Group("/api")
	routes.RegisterPublicRoutes(api)
	routes.RegisterAdminRoutes(api)

	log.Printf("🚀 Server running on http://localhost%s\n", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
