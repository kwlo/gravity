package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger interface
type Logger interface {
	Infof(template string, args ...interface{})
	Sync() error
}

// NewLogger creates new logger using zap library
func NewLogger() Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{
		"stdout",
		"main.log",
	}
	logger, _ := cfg.Build()

	return logger.Sugar()
}
