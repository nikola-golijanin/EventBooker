package routes

import (
	"homelab/event-booker/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event ID."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(404, gin.H{"error": "Could not fetch event or event not found."})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(500, gin.H{"error": "Could not register for event."})
		return
	}
	context.JSON(200, gin.H{"message": "Successfully registered for event."})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event ID."})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(500, gin.H{"error": "Could not cancel registration for event."})
		return
	}
	context.JSON(200, gin.H{"message": "Successfully canceled registration for event."})
}
