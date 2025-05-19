package http

import (
	"github.com/kimxuanhong/go-server/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindAndValidate binds and validates the request body
func BindAndValidate(c core.Context, obj interface{}) bool {
	if err := c.Bind(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

// SendResponse sends a JSON response with the given status code and data
func SendResponse(c core.Context, status int, data interface{}) {
	c.JSON(status, data)
}
