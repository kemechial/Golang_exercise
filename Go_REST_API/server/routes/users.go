package routes

import (
	"github.com/gin-gonic/gin"
	"example.com/testserver/models"
	"net/http"
)

func signup(context *gin.Context)	{
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully", "user": user})
}
