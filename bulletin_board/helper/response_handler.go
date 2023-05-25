package helper

import (
	"gin_test/bulletin_board/data/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseHandler(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := response.Response{
		Code:    statusCode,
		Status:  http.StatusText(statusCode),
		Message: message,
		Data:    data,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, response)
}
