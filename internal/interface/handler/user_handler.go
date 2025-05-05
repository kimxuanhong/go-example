package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-example/internal/usecase"
	"net/http"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("user")
	user, err := h.uc.GetUser(c, idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
