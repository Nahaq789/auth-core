package main

import (
	"log"
	"net/http"

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
	s.ListenAndServe()
}
