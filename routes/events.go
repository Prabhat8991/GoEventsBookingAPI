package routes

import (
	"api/api-request/utils"
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
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Not Authorized"})
		return
	}

	error := utils.VerifyToken(token)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Not Authorized"})
		return
	}

	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event) // Gin will parse incoming json to the event object
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	event.ID = 1
	event.UserID = 1001
	error = event.Save()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Bad request"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}

func updateEvent(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Could not parse ID")
		return
	}

	_, err = models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, "Could not find the event")
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindBodyWithJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}

	updatedEvent.UpdateEventById(id)
}

func deleteEvent(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Could not parse ID")
		return
	}

	models.DeleteEventById(id)

}
