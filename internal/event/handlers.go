package event

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateEventHandler(c *gin.Context) {
	var input struct {
		Name     string    `json:"name" binding:"required"`
		Location string    `json:"location" binding:"required"`
		Date     time.Time `json:"date" binding:"required"`
		Capacity int       `json:"capacity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := Event{
		Name:     input.Name,
		Location: input.Location,
		Date:     input.Date,
		Capacity: input.Capacity,
	}
	if err := CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create event"})
		return
	}
	c.JSON(http.StatusCreated, event)

}

func GetAllEventsHandler(c *gin.Context) {
	events, err := GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}
