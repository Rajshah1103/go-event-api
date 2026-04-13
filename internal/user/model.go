package user

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role string `gorm:"type:varchar(20);default:user"`  // default user role is user
	gorm.Model
}
