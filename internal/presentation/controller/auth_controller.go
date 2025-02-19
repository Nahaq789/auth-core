package controller

import (
	"context"
	"net/http"

	"github.com/auth-core/internal/application"
	"github.com/auth-core/internal/application/dto"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service application.UserService
}

func NewAuthController(s application.UserService) *AuthController {
	return &AuthController{service: s}
}

func (a *AuthController) Signup(c *gin.Context) {
	var user dto.UserDto
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ch := make(chan error, 1)
	go func(ch chan error, ctx context.Context) {
		ch <- a.service.CreateUser(ctx, &user)
	}(ch, context.Background())

	err := <-ch
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user created",
	})
}
