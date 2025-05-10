package api

import (
	"github.com/kimxuanhong/go-server/core"
	"net/http"

	"github.com/kimxuanhong/go-example/internal/interface/handler"
)

func UserRoutes(route *handler.UserHandler) []core.RouteConfig {
	return []core.RouteConfig{
		{
			Path:       "users/:user",
			Method:     http.MethodGet,
			Handler:    route.GetUser,
			Middleware: []core.Handler{},
		},
		{
			Path:       "users",
			Method:     http.MethodPost,
			Handler:    route.CreateUser,
			Middleware: []core.Handler{},
		},
		{
			Path:       "users/:user",
			Method:     http.MethodPut,
			Handler:    route.UpdateUser,
			Middleware: []core.Handler{},
		},
	}
}
