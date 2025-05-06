package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindAndValidate binds and validates the request body
func BindAndValidate(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

// SendResponse sends a JSON response with the given status code and data
func SendResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
