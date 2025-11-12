package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/api/events", getEvents)
	server.GET("/api/events/:id", getEventById)
	server.POST("/api/events", createEvent)
	server.PUT("/api/events/:id", updateEvent)
	server.DELETE("/api/events/:id", deleteEvent)
	server.POST("/api/signup")
}
