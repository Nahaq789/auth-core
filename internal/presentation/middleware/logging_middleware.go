package middleware

import (
	"github.com/gin-gonic/gin"
)

var (
	requestIdHeaderKey = "x-request-id"
	timeformat         = "2006/1/2 15:04:05.000 JTS"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}
