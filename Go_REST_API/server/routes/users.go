package routes

import (
	"github.com/gin-gonic/gin"
	"example.com/testserver/models"
	"net/http"
	"example.com/testserver/utils"
)

func signup(context *gin.Context)	{
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user","err": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {	
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token", "err": err.Error()})
		return
	}


	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
	



