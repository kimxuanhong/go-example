package handler

import (
	"fmt"
	"github.com/kimxuanhong/go-server/core"
	"net/http"

	"github.com/kimxuanhong/go-example/internal/facade"
	"github.com/kimxuanhong/go-example/internal/interface/dto"
)

// UserHandler
// @BaseUrl /users
type UserHandler struct {
	userFacade *facade.UserFacade
}

func NewUserHandler(userFacade *facade.UserFacade) *UserHandler {
	return &UserHandler{userFacade}
}

// GetUser
// @Api GET /:user
func (h *UserHandler) GetUser(c core.Context) {
	userName := c.Param("user")
	id := c.Query("id")
	fmt.Println("id= " + id)
	user, err := h.userFacade.GetUser(c.Context(), userName)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusOK, user)
}

// CreateUser
// @Api POST /
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

// UpdateUser
// @Api POST /:user
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
