package main

import (
	"github.com/auth-core/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

func Routing(r gin.IRouter) {
	v1 := r.Group("/api/v1")
	con := controller.AuthController{}
	{
		v1.POST("/signup", con.Signup)
	}
}
