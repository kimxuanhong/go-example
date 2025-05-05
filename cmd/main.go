package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-example/api"
	"github.com/kimxuanhong/go-example/di"
	"github.com/kimxuanhong/go-example/internal/infrastructure/repository"
	"log"
)

func main() {
	app, err := di.InitApp()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	app.Server.RegisterRoute("GET", "/", Pong())
	app.Server.Routes(api.UserRoutes(app.UserHandler))

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
	return &repository.UserModel{
		UserName: "hong",
	}, nil
}
