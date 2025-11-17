package routes

import (
	"homelab/event-booker/models"
	"homelab/event-booker/utils"
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

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
