package controller

import (
	"context"
	"net/http"

	"github.com/auth-core/internal/application"
	"github.com/auth-core/internal/application/dto"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService    application.UserService
	cognitoService application.CognitoService
}

func NewAuthController(s application.UserService, c application.CognitoService) *AuthController {
	return &AuthController{userService: s, cognitoService: c}
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
		ch <- a.userService.CreateUser(ctx, &user)
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
