package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-example/api"
	"github.com/kimxuanhong/go-example/di"
	"github.com/kimxuanhong/go-example/entity"
	"log"
)

func main() {
	app, err := di.InitApp()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	app.Server.RegisterRoute("GET", "/", Pong())
	api.RegisterHandler(app.Server, "GET", "/test", Test)

	app.Server.RegisterRoutes(func(rg *gin.RouterGroup) {
		rg.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
		rg.GET("/pong", func(c *gin.Context) { c.JSON(200, gin.H{"message": "ping"}) })
		rg.GET("/user", func(c *gin.Context) {
			var users []entity.User
			_ = app.Postgres.Select(c, &users, "user_name = ?", "123")
			c.JSON(200, users)
		})
	})

	// Chạy server
	if err := app.Server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func Pong() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	}
}

func Test(c *gin.Context) (interface{}, error) {
	return &entity.User{
		UserName: "hong",
	}, nil
}
