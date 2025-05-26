package main

import (
	"errors"
	"github.com/google/uuid"
	"github.com/kimxuanhong/go-cron/cron"
	"github.com/kimxuanhong/go-logger/logger"
	"github.com/kimxuanhong/go-server/core"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/go-example/di"
)

type MyHandler struct{}

// LinkAccount
// Giây     Phút     Giờ     Ngày     Tháng     Thứ
// */30      *        *       *         *        *
// //@Cron cron.link-account
func (m *MyHandler) LinkAccount() {
	println("run cron cron.link-account")
}

// Notify
// Giây     Phút     Giờ     Ngày     Tháng     Thứ
// */30      *        *       *         *        *
// //@Cron cron.notify
func (m *MyHandler) Notify() {
	println("run cron cron.notify")
}

func main() {
	err := logger.Init()
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
		return
	}
	app, err := di.InitApp()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	// Đăng ký middleware
	//jwtComp := jwt.NewJwt(app.Cfg.Jwt)
	app.Server.Use(func(c core.Context) {
		c.Set("requestId", uuid.NewString())
		c.Next()
	})
	app.Server.RegisterHandlersWithTags(app.Handlers...)
	app.Server.HealthCheck()

	scheduler := cron.NewCronJob()
	defer scheduler.Stop()
	scheduler.RegisterJobWithTags(&MyHandler{})
	_ = scheduler.Start()

	// Xử lý graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Chạy server
		if err := app.Server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %v", err)
		}
		sigChan <- syscall.SIGINT
	}()

	// Đợi signal để shutdown
	<-sigChan

}

func Pong() core.Handler {
	return func(c core.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	}
}
