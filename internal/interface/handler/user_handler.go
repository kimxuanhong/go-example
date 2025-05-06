package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-example/internal/facade"
	"github.com/kimxuanhong/go-example/internal/interface/dto"
)

type UserHandler struct {
	userFacade *facade.UserFacade
}

func NewUserHandler(userFacade *facade.UserFacade) *UserHandler {
	return &UserHandler{userFacade}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userName := c.Param("user")
	user, err := h.userFacade.GetUser(c, userName)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserRequest
	if !BindAndValidate(c, &req) {
		return
	}

	createdUser, err := h.userFacade.CreateUser(c, &req)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusCreated, createdUser)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userName := c.Param("user")
	var req dto.UserRequest
	if !BindAndValidate(c, &req) {
		return
	}

	updatedUser, err := h.userFacade.UpdateUser(c, userName, &req)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusOK, updatedUser)
}
