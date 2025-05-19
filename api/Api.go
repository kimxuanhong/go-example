package api

import (
	"github.com/kimxuanhong/go-server/core"

	"github.com/kimxuanhong/go-example/internal/delivery/http"
)

func UserRoutes(route *http.UserHandler) []core.RouteConfig {
	return []core.RouteConfig{
		{
			Path:       "users/:user",
			Method:     core.MethodGet,
			Handler:    route.GetUser,
			Middleware: []core.Handler{},
		},
		{
			Path:       "users",
			Method:     core.MethodPost,
			Handler:    route.CreateUser,
			Middleware: []core.Handler{},
		},
		{
			Path:       "users/:user",
			Method:     core.MethodPut,
			Handler:    route.UpdateUser,
			Middleware: []core.Handler{},
		},
	}
}
