package middlewares

import (
	interfaces "gin_test/bulletin_board/dao/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdmin(usersInterface interfaces.UsersInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the user ID from the context
		id, exists := ctx.Get("Id")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "You are not logged in."})
			return
		}

		// Retrieve the user information
		result, err := usersInterface.FindById(id.(int))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "User not found."})
			return
		}

		// Check if the user is an admin
		if result.Type != "1" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": 403, "status": "Forbidden", "message": "Only admins can access this route."})
			return
		}

		ctx.Next()
	}
}
