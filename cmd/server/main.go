// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	// Define the server address and port
// 	addr := ":8080"

// 	// Define a simple handler
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Welcome to Mini Soccer Backend!")
// 	})

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

	"github.com/joho/godotenv"
	"github.com/mattismyname3011/minisoccer-backend/config"
	"github.com/mattismyname3011/minisoccer-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDatabase()
	app := routes.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Listen(":" + port)
}
