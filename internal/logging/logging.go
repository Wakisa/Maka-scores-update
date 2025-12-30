package logging

import (
	"log/slog"
	"os"
)

// Logger interface for Maka
type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type slogLogger struct {
	logger *slog.Logger
}

// Debug implements Logger
func (l *slogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info implements Logger
func (l *slogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Error implements Logger
func (l *slogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

// New creates a new logger writing to stdout
func New() Logger {
	return &slogLogger{
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug, // default to debug level
		})),
	}
}
