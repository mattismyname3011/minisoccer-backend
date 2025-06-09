package seeder

import (
	"fmt"
	"time"

	"minisoccer-backend/config"
	"minisoccer-backend/models"
	"minisoccer-backend/utils"
)

func SeedUsers() {
	adminPass, _ := utils.HashPassword("admin123")
	userPass, _ := utils.HashPassword("user123")

	users := []models.User{
		{Email: "admin@example.com", PasswordHash: adminPass, Role: "admin"},
		{Email: "user@example.com", PasswordHash: userPass, Role: "user"},
	}

	for _, u := range users {
		if err := config.DB.Create(&u).Error; err != nil {
			fmt.Println("Failed to create user:", u.Email, "-", err)
		} else {
			fmt.Println("User created:", u.Email)
		}
	}
}

func SeedPricing() {
	db := config.DB

	// Ensure court exists first (assume one court for now)
	court := models.Court{
		Name:     "Mini Field A",
		Location: "Downtown Arena",
	}
	if err := db.FirstOrCreate(&court, models.Court{Name: court.Name}).Error; err != nil {
		fmt.Println("❌ Failed to create/find court:", err)
		return
	}

	// Time layout for parsing only hour and minute
	layout := "15:04"

	// Define pricing per time slot
	rawPricings := []struct {
		Start string
		End   string
		Price float64
	}{
		{"06:00", "12:00", 100000},
		{"12:00", "18:00", 150000},
		{"18:00", "23:00", 200000},
	}

	for _, raw := range rawPricings {
		start, err := time.Parse(layout, raw.Start)
		if err != nil {
			fmt.Println("❌ Failed to parse start time:", raw.Start)
			continue
		}
		end, err := time.Parse(layout, raw.End)
		if err != nil {
			fmt.Println("❌ Failed to parse end time:", raw.End)
			continue
		}

		pricing := models.Pricing{
			CourtID:   court.ID,
			StartTime: start,
			EndTime:   end,
			Price:     raw.Price,
		}

		if err := db.Create(&pricing).Error; err != nil {
			fmt.Printf("❌ Failed to seed pricing (%s-%s): %v\n", raw.Start, raw.End, err)
		} else {
			fmt.Printf("✅ Seeded pricing: %s - %s = Rp%.0f\n", raw.Start, raw.End, raw.Price)
		}
	}
}

func SeedBookings() {
	db := config.DB

	// Example: Booking 2 slots for today
	userID := uint(1)
	courtID := uint(1)
	date := time.Now().Truncate(24 * time.Hour)

	bookings := []models.Booking{
		{
			UserID:    userID,
			CourtID:   courtID,
			StartTime: date.Add(6 * time.Hour), // 06:00
			EndTime:   date.Add(9 * time.Hour), // 09:00
			Status:    "pending",
		},
		{
			UserID:    userID,
			CourtID:   courtID,
			StartTime: date.Add(15 * time.Hour), // 15:00
			EndTime:   date.Add(18 * time.Hour), // 18:00
			Status:    "booked",
		},
		{
			UserID:    userID,
			CourtID:   courtID,
			StartTime: date.Add(15 * time.Hour), // 15:00
			EndTime:   date.Add(18 * time.Hour), // 18:00
			Status:    "cancelled",
		},
	}

	for _, b := range bookings {
		if err := db.Create(&b).Error; err != nil {
			fmt.Println("❌ Failed to create booking:", err)
		} else {
			fmt.Printf("✅ Booking seeded: %s to %s\n", b.StartTime.Format("15:04"), b.EndTime.Format("15:04"))
		}
	}
	// for _, u := range users {
	// 	if err := config.DB.Create(&u).Error; err != nil {
	// 		fmt.Println("Failed to create user:", u.Email, "-", err)
	// 	} else {
	// 		fmt.Println("User created:", u.Email)
	// 	}
	// }
}
