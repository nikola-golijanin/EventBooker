package routes

import (
	"homelab/event-booker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
	}

	err = user.Save()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user_id": user.ID})
}
