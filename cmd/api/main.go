package main

import (
	"context"
	"log"
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
	ctx := context.Background()
	a, err := conf.LoadAppSetting()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	logger := logger.InitLogger(a.Server.Level)

	awsClient, err := conf.InitClient(ctx, &a.Aws)
	if err != nil {
		log.Fatalf("aws client error : %v", err)
	}

	cr := di.Initialize(&logger.Logger, &awsClient.Dynamodb, &awsClient.Cognito, &a.Aws)
	err = Routing(ctx, r, *cr, *logger)
	if err != nil {
		log.Fatal("routing error: %", err)
	}

	s := &http.Server{
		Addr:    ":" + a.Server.Port,
		Handler: r,
	}

	logger.Logger.Info("Server is running on port", "port", a.Server.Port)
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	<-ctx.Done()
}

func Routing(
	ctx context.Context,
	r gin.IRouter,
	cr di.ControllerSet,
	logger logger.LoggerConfig) error {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.LoggingMiddleware(&logger))
	{
		v1.POST("/auth/signup", cr.AuthController.Signup)
		v1.POST("/auth/confirm", cr.AuthController.ConfirmSignUp)
		v1.POST("/auth/signin", cr.AuthController.InitiateAuth)
		v1.POST("/auth/challenge", cr.AuthController.AuthChallenge)
	}

	return nil
}
