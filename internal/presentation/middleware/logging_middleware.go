package middleware

import (
	"log/slog"
	"time"

	"github.com/auth-core/pkg/logger"
	"github.com/gin-gonic/gin"
)

var (
	requestIdHeaderKey = "x-request-id"
	timeformat         = "2006/1/2 15:04:05.000 JTS"
)

func LoggingMiddleware(config *logger.LoggerConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			config.Logger.Error("Failed to load location for Asia/Tokyo", "error", err)
		}
		start := time.Now().In(jst)
		startStr := start.Format(timeformat)
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		method := ctx.Request.Method
		clientIP := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		requestID := ctx.GetHeader(requestIdHeaderKey)

		params := map[string]string{}
		for _, p := range ctx.Params {
			params[p.Key] = p.Value
		}

		loggerWithRequestID := config.Logger.With("RequestID", requestID)

		requestAttr := []slog.Attr{
			slog.String("Time", startStr),
			slog.String("Method", method),
			slog.String("Path", path),
			slog.String("Query", query),
			slog.Any("Params", params),
			slog.String("ClientIP", clientIP),
			slog.String("User Agent", userAgent),
		}

		loggerWithRequestID.LogAttrs(
			ctx.Request.Context(),
			config.BaseLogLevel,
			"Request Log",
			requestAttr...,
		)

		ctx.Next()

		end := time.Now().In(jst)
		endStr := end.Format(timeformat)
		latency := end.Sub(start)
		status := ctx.Writer.Status()

		logLevel := logger.DetermineLogLevel(status)

		responseAttr := []slog.Attr{
			slog.String("Time", endStr),
			slog.String("Latency", logger.ConvertLatency(latency)),
			slog.Int("Status", status),
		}

		loggerWithRequestID.LogAttrs(
			ctx.Request.Context(),
			logLevel,
			"Response Log",
			responseAttr...,
		)
	}
}
