package routes

import (
	"github.com/gin-gonic/gin"
	"example.com/testserver/middlewares"
)


func RegisterRoutes(server *gin.Engine) {
	
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/shutdown", shutdownServer)	
	
}