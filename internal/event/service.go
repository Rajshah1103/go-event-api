package event

import "github.com/Rajshah1103/event-booking-api/internal/db"

func CreateEvent(event *Event) error {
	return db.DB.Create(event).Error
}

func GetAllEvents() ([]Event, error) {
	var events []Event
	err := db.DB.Find(&events).Error
	return events, err
}
