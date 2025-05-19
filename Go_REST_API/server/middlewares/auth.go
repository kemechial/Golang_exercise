package	middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"example.com/testserver/utils" // Adjust the import path to your utils package
)



func Authenticate(context *gin.Context) {
	// Check if the request has a valid token
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
		context.Abort()
		return
	}

	userId, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Set the user ID in the context
	context.Set("userId", userId)
	context.Next()
	
	
}