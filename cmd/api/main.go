package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/auth-core/cmd/conf"
	"github.com/auth-core/cmd/di"
	"github.com/auth-core/internal/presentation/middleware"
	"github.com/auth-core/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	a, err := conf.LoadAppSetting()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	err = Routing(r, a)
	if err != nil {
		log.Fatal("routing error: %", err)
	}

	s := &http.Server{
		Addr:    ":" + a.Server.Port,
		Handler: r,
	}

	log.Println("Server is running on port", a.Server.Port)
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	<-ctx.Done()
}

func Routing(ctx context.Context, r gin.IRouter, setting *conf.AppSetting) error {
	v1 := r.Group("/api/v1")

	lc := logger.InitLogger(setting.Server.Level)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	v1.Use(middleware.LoggingMiddleware(logger, lc))

	cr := di.Initialize(dynamodb, cognit, &setting.Aws)
	{
		v1.POST("/signup", cr.AuthController.Signup)
	}

	return nil
}
