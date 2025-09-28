package user

import (
	"github.com/Rajshah1103/event-booking-api/internal/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register a user
func Register(username, password string) (User, error) {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	// create user
	u := User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := db.DB.Create(&u).Error; err != nil {
		return User{}, err
	}
	return u, nil
}

func AuthenticateUser(username, password string) (User, error) {
	var u User
	if err := db.DB.Where("username = ?", username).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return User{}, err
		}
		return User{}, err
	}
	return u, nil
}
