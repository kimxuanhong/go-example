package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusOK, dto.ToUserResponse(user))
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserRequest
	if !BindAndValidate(c, &req) {
		return
	}

	user := dto.ToUserDomain(&req)
	createdUser, err := h.uc.CreateUser(c, user)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusCreated, dto.ToUserResponse(createdUser))
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userName := c.Param("user")
	var req dto.UserRequest
	if !BindAndValidate(c, &req) {
		return
	}

	user := dto.ToUserDomain(&req)
	user.UserName = userName // Ensure username from path is used

	updatedUser, err := h.uc.UpdateUser(c, user)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusOK, dto.ToUserResponse(updatedUser))
}
