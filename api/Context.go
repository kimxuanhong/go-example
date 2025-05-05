package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-http/server"
)

func RegisterHandler(s server.Server, method, path string, handler func(ctx *gin.Context) (interface{}, error)) {
	s.RegisterRoute(method, path, func(c *gin.Context) {
		res, err := handler(c)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, res)
	})
}
