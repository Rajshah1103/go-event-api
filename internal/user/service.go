package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var users = []User{}
var nextID = 1

// Register a user 
func Register(username, password string) (User , error) {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	// create user
	u := User {
		ID : nextID,
		Username: username,
		Password: string(hashedPassword),
	}
	nextID++
	users = append(users, u)

	return u, nil
} 