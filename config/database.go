package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Replace with your database driver
)

func InitDatabase() *sql.DB {
	// Replace with your database connection details
	connStr := "user=username dbname=dbname password=password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection is not alive: %v", err)
	}

	log.Println("Database connection established")
	return db
}
