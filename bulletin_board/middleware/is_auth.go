package middlewares

import (
	"fmt"
	interfaces "gin_test/bulletin_board/dao/user"
	"gin_test/bulletin_board/helper"

	"gin_test/bulletin_board/utils"

	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func IsAuth(usersInterface interfaces.UsersInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// authorizeHeader := ctx.GetHeader("Authorization")
		// token := strings.TrimPrefix(authorizeHeader, "Bearer ")
		cookie, err := ctx.Request.Cookie("token")
		if err != nil || cookie.Value == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "You are not logged in."})
			return
		}
		token := cookie.Value
		errenv := godotenv.Load(".env")
		helper.ErrorPanic(errenv)
		tokenSecret := os.Getenv("TOKEN_SECRET")
		sub, err := utils.ValidateToken(token, tokenSecret)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": err.Error()})
			return
		}

		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		helper.ErrorPanic(err_id)
		result, err := usersInterface.FindById(id)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "status": "Unauthorized", "message": "User not found."})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Set("Id", id)
		ctx.Next()
	}
}
