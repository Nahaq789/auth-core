package logger

import (
	"log/slog"
	"time"
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
		WithBaseLogLevel(lv),
	)
	return lc
}
