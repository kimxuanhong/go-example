package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-example/internal/interface/handler"
	"github.com/kimxuanhong/go-http/server"
	"net/http"
)

func UserRoutes(route *handler.UserHandler) []server.RouteConfig {
	return []server.RouteConfig{
		{
			Path:       "users/:user",
			Method:     http.MethodGet,
			HandleFunc: route.GetUser,
			Middleware: []gin.HandlerFunc{},
		},
	}
}
