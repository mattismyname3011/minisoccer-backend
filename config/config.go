package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	AppName  string
	Port     string
	Database string
}

// LoadConfig loads configuration from environment variables or .env file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		AppName:  getEnv("APP_NAME", "MiniSoccerBackend"),
		Port:     getEnv("PORT", "8080"),
		Database: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/minisoccer"),
	}
}

// getEnv retrieves environment variables or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
