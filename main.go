package main

import (
	"api/db"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", handleEvent)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func handleEvent(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event) // Gin will parse incoming json to the event object
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	event.ID = 1
	event.UserID = 1001
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}
