package helper

import (
	"github.com/gin-gonic/gin"
)

// Response struct for consistent API responses
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse formats and sends a consistent API response
func SendResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
