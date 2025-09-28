package booking

import "github.com/Rajshah1103/event-booking-api/internal/db"

func CreateBooking(b *Booking) error {
	return db.DB.Create(b).Error
}

func GetBookingByUserId(userId uint) ([]Booking, error) {
	var bookings []Booking
	err := db.DB.Preload("Event").Where("user_id = ?", userId).Find(&bookings).Error
	return bookings, err
}