package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", handleEvents)
	server.GET("/events/:id", handleEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", createUser)
	server.POST("/login", login)
	server.Run(":8080")
}
