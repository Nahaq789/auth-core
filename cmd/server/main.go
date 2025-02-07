package main

import (
	"fmt"

	"github.com/auth-core/cmd/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := conf.SetEnv()
	fmt.Println(conf.Addr)
	fmt.Printf("addr: %s", conf.Port)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
