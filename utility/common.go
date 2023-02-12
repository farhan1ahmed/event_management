package utility

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GenerateResponse(ctx *gin.Context, statusCode int, message string, isError bool, responseBody interface{}){
	if responseBody == nil{
		status := "success"
		if isError == true{
			status = "error"
		}
		message = strings.ToLower(message)
		responseBody = gin.H{
			"status": status,
			"message": message,
		}
	}
	ctx.JSON(statusCode, responseBody)
}