package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	internalLogger *logrus.Logger
}

func New() Logger {
	internalLogger := logrus.New()
	internalLogger.SetFormatter(&logrus.JSONFormatter{})
	return Logger{
		internalLogger: internalLogger,
	}
}

func (l *Logger) Debug(ctx context.Context, message string, fields map[string]any) {
	l.internalLogger.WithContext(ctx).WithFields(fields).Debug(message)
}

func (l *Logger) Info(ctx context.Context, message string, fields map[string]any) {
	l.internalLogger.WithContext(ctx).WithFields(fields).Info(message)
}

func (l *Logger) Warn(ctx context.Context, message string, fields map[string]any) {
	l.internalLogger.WithContext(ctx).WithFields(fields).Warn(message)
}

func (l *Logger) Error(ctx context.Context, message string, fields map[string]any) {
	l.internalLogger.WithContext(ctx).WithFields(fields).Error(message)
}
