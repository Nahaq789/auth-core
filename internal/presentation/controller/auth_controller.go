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
	cognitoService application.AuthService
}

func NewAuthController(s application.UserService, c application.AuthService) *AuthController {
	return &AuthController{userService: s, cognitoService: c}
}

func (a *AuthController) Signup(c *gin.Context) {
	var auth dto.SignUpDto
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

func (a *AuthController) ConfirmSignUp(c *gin.Context) {
	var code *dto.ConfirmSignUpDto
	if err := c.ShouldBind(&code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	ctx := context.Background()
	err := a.cognitoService.ConfirmSignUp(ctx, code)
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

func (a *AuthController) InitiateAuth(c *gin.Context) {
	var credential *dto.SignInDto
	if err := c.ShouldBind(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	ctx := context.Background()
	res, err := a.cognitoService.InitiateAuth(ctx, credential)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (a *AuthController) AuthChallenge(c *gin.Context) {
	var challenge *dto.AuthChallengeDto
	if err := c.ShouldBind(&challenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx := context.Background()
	result, err := a.cognitoService.AuthChallenge(ctx, challenge)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
