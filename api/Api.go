package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-example/internal/interface/handler"
	"github.com/kimxuanhong/go-http/server"
)

func UserRoutes(route *handler.UserHandler) []server.RouteConfig {
	return []server.RouteConfig{
		{
			Path:       "users/:user",
			Method:     http.MethodGet,
			HandleFunc: route.GetUser,
			Middleware: []gin.HandlerFunc{},
		},
		{
			Path:       "users",
			Method:     http.MethodPost,
			HandleFunc: route.CreateUser,
			Middleware: []gin.HandlerFunc{},
		},
		{
			Path:       "users/:user",
			Method:     http.MethodPut,
			HandleFunc: route.UpdateUser,
			Middleware: []gin.HandlerFunc{},
		},
	}
}
