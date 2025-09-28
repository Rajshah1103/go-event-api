package booking

import (
	"github.com/Rajshah1103/event-booking-api/internal/event"
	"github.com/Rajshah1103/event-booking-api/internal/user"
)

type Booking struct {
	ID      uint        `gorm:"primaryKey"`
	UserID  uint        `gorm:"not null"`
	EventID uint        `gorm:"not null"`
	User    user.User   `gorm:"foreignKey:UserID"`  // optional, for Preload
	Event   event.Event `gorm:"foreignKey:EventID"` // optional, for Preload
}
