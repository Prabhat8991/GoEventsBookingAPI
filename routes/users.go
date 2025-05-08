package routes

import (
	"api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		log.Println("createUser err", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	error := user.CreateUser()

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User created"})
}
