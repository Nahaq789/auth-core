package logger

import (
	"log/slog"
	"os"
	"time"
)

type LoggerConfig struct {
	BaseLogLevel slog.Level
	Logger       slog.Logger
}

type Option func(*LoggerConfig)

func NewLoggerConfig(opts ...Option) *LoggerConfig {
	lc := &LoggerConfig{}

	for _, opt := range opts {
		opt(lc)
	}

	return lc
}

func withBaseLogLevel(level slog.Level) Option {
	return func(c *LoggerConfig) {
		c.BaseLogLevel = level
	}
}

func withSlog() Option {
	return func(c *LoggerConfig) {
		log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		c.Logger = *log
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

func DetermineLogLevel(status int) slog.Level {
	if status == 200 {
		return slog.LevelInfo
	}
	return slog.LevelError
}

func ConvertLatency(latency time.Duration) string {
	return latency.String()
}

func InitLogger(level string) *LoggerConfig {
	lv := ConvertLevel(level)
	lc := NewLoggerConfig(
		withBaseLogLevel(lv),
		withSlog(),
	)
	return lc
}
