package routes

import (
	"homelab/event-booker/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/api/events", getEvents)
	server.GET("/api/events/:id", getEventById)

	authenticated := server.Group("/api")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/api/signup", signup)
	server.POST("/api/login", login)
}
