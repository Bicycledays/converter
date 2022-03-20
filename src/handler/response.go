package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type response struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

func newResultResponse(c *gin.Context, result interface{}) {
	log.Println("response", result)
	c.JSON(http.StatusOK, response{
		Success: true,
		Result:  result,
	})
}
