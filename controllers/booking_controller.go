package controllers

import (
	"minisoccer-backend/config"
	"minisoccer-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GET /api/pricing
func GetPricing(c *fiber.Ctx) error {
	var pricings []models.Pricing
	db := config.DB

	if err := db.Preload("Court").Find(&pricings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to load pricing"})
	}

	return c.JSON(pricings)
}

// GET /api/availability?date=YYYY-MM-DD
func GetAvailability(c *fiber.Ctx) error {
	dateStr := c.Query("date")
	if dateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "date query parameter is required"})
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid date format"})
	}

	db := config.DB

	var pricings []models.Pricing
	if err := db.Find(&pricings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to load pricings"})
	}

	var bookings []models.Booking
	startOfDay := date
	endOfDay := date.AddDate(0, 0, 1)

	if err := db.Where("start_time >= ? AND start_time < ?", startOfDay, endOfDay).Find(&bookings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to load bookings"})
	}

	type Availability struct {
		ID        uint    `json:"id"`
		StartTime string  `json:"start_time"`
		EndTime   string  `json:"end_time"`
		Price     float64 `json:"price"`
		Available bool    `json:"available"`
	}

	availabilities := make([]Availability, 0, len(pricings))

	for _, p := range pricings {
		start, _ := time.Parse("15:04", p.StartTime.Format("15:04"))
		end, _ := time.Parse("15:04", p.EndTime.Format("15:04"))

		start = time.Date(date.Year(), date.Month(), date.Day(), start.Hour(), start.Minute(), 0, 0, time.UTC)
		end = time.Date(date.Year(), date.Month(), date.Day(), end.Hour(), end.Minute(), 0, 0, time.UTC)

		available := true
		for _, b := range bookings {
			if b.StartTime.Before(end) && b.EndTime.After(start) {
				available = false
				break
			}
		}

		availabilities = append(availabilities, Availability{
			ID:        p.ID,
			StartTime: p.StartTime.Format("2006-01-02T15:04:05Z07:00"),
			EndTime:   p.EndTime.Format("2006-01-02T15:04:05Z07:00"),
			Price:     p.Price,
			Available: available,
		})
	}

	return c.JSON(availabilities)
}

// POST /api/bookings
func CreateBooking(c *fiber.Ctx) error {
	type BookingInput struct {
		UserID    uint   `json:"user_id"`
		CourtID   uint   `json:"court_id"`
		StartTime string `json:"start_time"` // ISO8601 format expected
		EndTime   string `json:"end_time"`
	}

	var input BookingInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	startTime, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid start_time format"})
	}

	endTime, err := time.Parse(time.RFC3339, input.EndTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid end_time format"})
	}

	if !endTime.After(startTime) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "end_time must be after start_time"})
	}

	db := config.DB

	var count int64
	db.Model(&models.Booking{}).
		Where("court_id = ? AND start_time < ? AND end_time > ?", input.CourtID, endTime, startTime).
		Count(&count)

	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "court is already booked for the selected time range"})
	}

	booking := models.Booking{
		UserID:    input.UserID,
		CourtID:   input.CourtID,
		StartTime: startTime,
		EndTime:   endTime,
		Status:    "booked",
	}

	if err := db.Create(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create booking"})
	}

	return c.Status(fiber.StatusCreated).JSON(booking)
}
