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

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messgae": "Bad request"})
		return
	}

	error := user.ValidateCredentials()

	if error != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Credentials do not match"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Login Successful"})

}
