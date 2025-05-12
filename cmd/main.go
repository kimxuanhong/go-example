package main

import (
	"github.com/kimxuanhong/go-middleware/middleware"
	"github.com/kimxuanhong/go-server/core"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-example/di"
)

func main() {
	app, err := di.InitApp()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	// Đăng ký middleware
	//app.Server.RegisterMiddleware(middleware.RecoveryMiddleware())
	//app.Server.RegisterMiddleware(middleware.LogRequestMiddleware())
	//app.Server.RegisterMiddleware(middleware.LogResponseMiddleware())
	app.Server.Add("GET", "/", Pong())
	app.Server.SetHandlers(app.Handlers...)

	// Xử lý graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Chạy server
		if err := app.Server.Start(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// Đợi signal để shutdown
	<-sigChan
	log.Println("Shutting down server...")

	// In metrics trước khi shutdown
	middleware.GetMetrics().PrintMetrics()
}

func Pong() core.Handler {
	return func(c core.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	}
}
