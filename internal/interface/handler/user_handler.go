package handler

import (
	"github.com/kimxuanhong/go-server/core"
	"net/http"

	"github.com/kimxuanhong/go-example/internal/facade"
	"github.com/kimxuanhong/go-example/internal/interface/dto"
)

type UserHandler struct {
	userFacade *facade.UserFacade
}

func NewUserHandler(userFacade *facade.UserFacade) *UserHandler {
	return &UserHandler{userFacade}
}

func (h *UserHandler) GetUser(c core.Context) {
	userName := c.Param("user")
	user, err := h.userFacade.GetUser(c.Context(), userName)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c core.Context) {
	var req dto.UserRequest
	if !BindAndValidate(c, &req) {
		return
	}

	createdUser, err := h.userFacade.CreateUser(c.Context(), &req)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusCreated, createdUser)
}

func (h *UserHandler) UpdateUser(c core.Context) {
	userName := c.Param("user")
	var req dto.UserRequest
	if !BindAndValidate(c, &req) {
		return
	}

	updatedUser, err := h.userFacade.UpdateUser(c.Context(), userName, &req)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusOK, updatedUser)
}
