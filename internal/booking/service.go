package booking

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Rajshah1103/event-booking-api/internal/db"
	"github.com/Rajshah1103/event-booking-api/internal/kafka"
)

func CreateBooking(b *Booking) error {
	// Create booking record
	if err := db.DB.Create(b).Error; err != nil {
		return err
	}

	// ✅ Reload the booking with related User & Event
	if err := db.DB.Preload("User").Preload("Event").First(b, b.ID).Error; err != nil {
		log.Println("Failed to load full booking:", err)
		return err
	}

	// Marshal booking to JSON
	bookingJSON, err := json.Marshal(b)
	if err != nil {
		log.Println("Failed to marshal booking:", err)
		return err
	}

	// Send to Kafka
	if err := kafka.SendMessage("booking-created", fmt.Sprint(b.ID), bookingJSON); err != nil {
		log.Println("Failed to send booking to Kafka:", err)
		return err
	}

	return nil
}

func GetBookingByUserId(userId uint) ([]Booking, error) {
	var bookings []Booking
	err := db.DB.Preload("Event").Where("user_id = ?", userId).Find(&bookings).Error
	return bookings, err
}
