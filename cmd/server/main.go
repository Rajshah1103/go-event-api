package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// health check api
	r.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"message": "Event booking api is running",
		})
	})

	// start the server
	r.Run(":8080") // listens on port 80
}