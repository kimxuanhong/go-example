package http

import (
	"github.com/kimxuanhong/go-example/internal/infrastructure/external"
	"github.com/kimxuanhong/go-logger/logger"
	"github.com/kimxuanhong/go-server/core"
	"net/http"

	"github.com/kimxuanhong/go-example/internal/delivery/dto"
	"github.com/kimxuanhong/go-example/internal/facade"
)

// UserHandler
// @BaseUrl /v1
type UserHandler struct {
	userFacade    *facade.UserFacade
	accountClient *external.AccountClient
}

func NewUserHandler(userFacade *facade.UserFacade, accountClient *external.AccountClient) *UserHandler {
	return &UserHandler{userFacade, accountClient}
}

// GetUser
// @Api GET /users/:user
func (h *UserHandler) GetUser(c core.Context) {
	logger.Log.Info(c.Context().Value("requestId"))
	log := logger.WithContext(c.Context())
	userName := c.Param("user")
	log.Infof("id= %s", "123124124123213213")
	user, err := h.userFacade.GetUser(c.Context(), userName)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusOK, user)
}

// CreateUser
// @Api POST /users
func (h *UserHandler) CreateUser(c core.Context) {
	var req dto.UserRequest
	if !BindAndValidate(c, &req) {
		return
	}

	//user, err := h.accountClient.GetUser(c.Context(), req.UserName, req.Email)
	//fmt.Println(user, err)

	createdUser, err := h.userFacade.CreateUser(c.Context(), &req)
	if err != nil {
		HandleError(c, err)
		return
	}

	SendResponse(c, http.StatusCreated, createdUser)
}

// UpdateUser
// @Api POST /users/:user
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
