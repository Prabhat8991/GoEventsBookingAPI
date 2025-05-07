package routes

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Err while fetch": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func handleEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Could not parse ID")
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	context.JSON(http.StatusOK, event)

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
	error := event.Save()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Bad request"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}
