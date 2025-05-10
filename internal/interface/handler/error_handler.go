package handler

import (
	"errors"
	"github.com/kimxuanhong/go-server/core"
	"net/http"

	"github.com/gin-gonic/gin"
	domainErrors "github.com/kimxuanhong/go-example/internal/domain/errors"
)

// HandleError xử lý lỗi một cách thống nhất và trả về response phù hợp
func HandleError(c core.Context, err error) {
	switch {
	case errors.Is(err, domainErrors.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, domainErrors.ErrValidation):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, domainErrors.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case errors.Is(err, domainErrors.ErrForbidden):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	case errors.Is(err, domainErrors.ErrInternal):
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
