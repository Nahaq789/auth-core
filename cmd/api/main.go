package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/auth-core/cmd/conf"
	"github.com/auth-core/cmd/di"
	"github.com/auth-core/pkg/db"
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

func Routing(r gin.IRouter, settng *conf.AppSetting) error {
	v1 := r.Group("/api/v1")
	client, err := db.NewDynamoDbClient(context.Background(), settng.Aws.Region)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	cr := di.Initialize(client)
	{
		v1.POST("/signup", cr.AuthController.Signup)
	}

	return nil
}
