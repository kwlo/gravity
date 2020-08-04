package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitializeLogging initialize logger
func InitializeLogging() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{
		"stdout",
		"main.log",
	}
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)

	return logger
}
