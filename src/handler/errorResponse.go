package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func newErrorResponse(c *gin.Context, statusCode int, message, error string) {
	log.Println("error", message, error)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Success: false,
		Message: message,
		Error:   error,
	})
}
