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
	var auth dto.AuthDto
	if err := c.ShouldBind(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx := context.Background()
	err := a.cognitoService.SignUp(ctx, &auth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user created",
	})
	return
}

func (a *AuthController) VerifyCode(c *gin.Context) {
	var code *dto.VerifyCodeDto
	if err := c.ShouldBind(&code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	ctx := context.Background()
	err := a.cognitoService.VerifyCode(ctx, code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "verify code success",
	})
}
