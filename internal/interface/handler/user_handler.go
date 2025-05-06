package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	domainErrors "github.com/kimxuanhong/go-example/internal/domain/errors"
	"github.com/kimxuanhong/go-example/internal/interface/dto"
	"github.com/kimxuanhong/go-example/internal/usecase"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userName := c.Param("user")
	user, err := h.uc.GetUser(c, userName)
	if err != nil {
		switch {
		case errors.Is(err, domainErrors.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, domainErrors.ErrValidation):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	response := dto.ToUserResponse(user)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := dto.ToUserDomain(&req)
	createdUser, err := h.uc.CreateUser(c, user)
	if err != nil {
		switch {
		case errors.Is(err, domainErrors.ErrValidation):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errors.Is(err, domainErrors.ErrInternal):
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	response := dto.ToUserResponse(createdUser)
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userName := c.Param("user")
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := dto.ToUserDomain(&req)
	user.UserName = userName // Ensure username from path is used

	updatedUser, err := h.uc.UpdateUser(c, user)
	if err != nil {
		switch {
		case errors.Is(err, domainErrors.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, domainErrors.ErrValidation):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errors.Is(err, domainErrors.ErrInternal):
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	response := dto.ToUserResponse(updatedUser)
	c.JSON(http.StatusOK, response)
}
