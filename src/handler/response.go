package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type response struct {
	Success bool `json:"success"`
}

func newResponse(c *gin.Context) {
	log.Println("Success: true")
	c.JSON(http.StatusOK, response{
		Success: true,
	})
}
