package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/auth-core/cmd/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	a, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s := &http.Server{
		Addr:    a.Port,
		Handler: r,
	}

	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		<-ctx.Done()

		c, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := s.Shutdown(c); err != nil {
			log.Printf("Server shutdown error: %v", err)
			return
		}
	}(ctx)

	s.ListenAndServe()
	wg.Wait()
}
