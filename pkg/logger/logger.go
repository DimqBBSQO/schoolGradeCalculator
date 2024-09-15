package logger

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

var (
	logger    *zap.SugaredLogger
	startTime time.Time
)

func ZapLoggerInit() {
	startTime = time.Now()

	zapLogger, _ := zap.NewProduction()
	logger = zapLogger.Sugar()
}

func Sync() {
	if err := logger.Sync(); err != nil {
		return
	}
}

func Info(message string) {
	logger.Info(message)
}

func Infof(message string, args ...any) {
	logger.Infof(message, args...)
}

func Debug(message string) {
	logger.Debug(message)
}

func Debugf(message string, args ...any) {
	logger.Debugf(message, args...)
}

func Warn(message string) {
	logger.Warn(message)
}

func Warnf(message string, args ...any) {
	logger.Warnf(message, args...)
}

func Error(message string) {
	logger.Error(message)
}

func Errorf(message string, args ...any) {
	logger.Errorf(message, args...)
}

func Fatal(message string) {
	logger.Fatal(message)
}

func Fatalf(message string, args ...any) {
	logger.Fatalf(message, args...)
}

func ServiceError(custom, err error, data ...any) {
	Error(fmt.Sprintf("%v, err - %v, info - %+v", custom, err, data))
}

func Uptime() time.Duration {
	return time.Since(startTime)
}
