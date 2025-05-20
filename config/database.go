package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"minisoccer-backend/models"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL is not set in .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if os.Getenv("AUTO_MIGRATE") == "true" {
		// Auto migrate models
		err = db.AutoMigrate(
			&models.User{},
			// &models.Court{},
			// &models.TimeSlot{},
			// &models.Booking{},
			// &models.Addon{},
			// &models.Pricing{},
		)
		if err != nil {
			log.Fatalf("Auto migration failed: %v", err)
		}
	}

	DB = db
	fmt.Println("Database connection established.")
	return db
}
