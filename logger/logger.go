package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxLogger struct{}

func LoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	if l, ok := ctx.Value(ctxLogger{}).(*zap.SugaredLogger); ok {
		return l
	}
	return initLogger()
}

func ContextWithLogger(ctx context.Context, l *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

func initLogger() *zap.SugaredLogger {
	opts := make([]zap.Option, 0)

	opts = append(opts, zap.AddCaller())
	cfg := zap.NewDevelopmentConfig()
	level := zap.DebugLevel

	cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(
				cfg.EncoderConfig,
			),
			os.Stderr,
			level,
		),
		opts...,
	).Sugar()
}
