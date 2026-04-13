package main

import (
	"net/http"

	"github.com/Rajshah1103/event-booking-api/internal/booking"
	"github.com/Rajshah1103/event-booking-api/internal/db"
	"github.com/Rajshah1103/event-booking-api/internal/event"
	"github.com/Rajshah1103/event-booking-api/internal/kafka"
	"github.com/Rajshah1103/event-booking-api/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load env file
	godotenv.Load(".env")
	// init DB
	db.InitDB()

	// auto migrate models
	db.DB.AutoMigrate(&user.User{}, &event.Event{}, &booking.Booking{})
	// initialize gin
	r := gin.Default()

	// health check api
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Event booking api is running",
		})
	})
	r.POST("/register", user.RegisterHandler)
	r.POST("/login", user.LoginHandler)
	r.GET("/protected", user.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "You have accessed a protected route"})
	})

	// Events api
	r.POST("/events", user.JWTAuthMiddleware(), event.CreateEventHandler)
	r.GET("/events", user.JWTAuthMiddleware(), event.GetAllEventsHandler)

	//Booking api
	r.POST("/bookings", user.JWTAuthMiddleware(), booking.CreateBookingHandler)
	r.GET("/bookings", user.JWTAuthMiddleware(), booking.GetBookingsHandler)

	go kafka.StartBookingConsumer("localhost:9092", "booking-created", "booking-group-1")

	// start the server
	r.Run(":8080") // listens on port 80
}
