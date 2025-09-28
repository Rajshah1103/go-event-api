package booking

import (
	"net/http"

	"github.com/Rajshah1103/event-booking-api/internal/user"
	"github.com/gin-gonic/gin"
)

func CreateBookingHandler(c *gin.Context) {
	claims := c.MustGet("claims").(user.Claims) // JWT middleware

	var input struct {
		EventID uint `json:"event_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booking := Booking{
		UserID:  claims.ID,
		EventID: input.EventID,
	}

	if err := CreateBooking(&booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

func GetBookingsHandler(c *gin.Context) {
	claims := c.MustGet("claims").(user.Claims) // JWT middleware
	bookings, err := GetBookingByUserId(claims.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}
