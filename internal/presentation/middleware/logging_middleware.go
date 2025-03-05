package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	requestIdHeaderKey = "x-request-id"
	timeformat         = "2006/1/2 15:04:05.000 JTS"
)

type LoggerConfig struct {
	BaseLogLevel slog.Level
}

type Option func(*LoggerConfig)

func NewLoggerConfig(opts ...Option) *LoggerConfig {
	lc := &LoggerConfig{}

	for _, opt := range opts {
		opt(lc)
	}

	return lc
}

func WithBaseLogLevel(level slog.Level) Option {
	return func(c *LoggerConfig) {
		c.BaseLogLevel = level
	}
}

func ConvertLevel(level string) slog.Level {
	if level == "INFO" {
		return slog.LevelInfo
	}
	if level == "DEBUG" {
		return slog.LevelDebug
	}
	if level == "WARN" {
		return slog.LevelWarn
	}
	if level == "ERROR" {
		return slog.LevelError
	}
	return slog.LevelInfo
}

func LoggingMiddleware(logger *slog.Logger, config *LoggerConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			logger.Error("Failed to load location for Asia/Tokyo", "error", err)
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

		loggerWithRequestID := logger.With("RequestID", requestID)

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

		logLevel := determineLogLevel(status)

		responseAttr := []slog.Attr{
			slog.String("Time", endStr),
			slog.String("Latency", convertLatency(latency)),
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

func determineLogLevel(status int) slog.Level {
	if status == 200 {
		return slog.LevelInfo
	}
	return slog.LevelError
}

func convertLatency(latency time.Duration) string {
	return latency.String()
}
