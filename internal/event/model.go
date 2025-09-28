package event

import "time"

type Event struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"not null"`
	Location string    `gorm:"not null"`
	Date     time.Time `gorm:"not null"`
	Capacity int       `gorm:"not null"`
}
