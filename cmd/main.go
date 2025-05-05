package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-example/app"
	"log"
)

func main() {
	initApp, err := app.InitApp()
	if err != nil {
		log.Fatalf("failed to init initApp: %v", err)
	}

	server := initApp.Server
	server.RegisterRoute("GET", "/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Chạy server
	if err := initApp.Server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
