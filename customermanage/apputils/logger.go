package apputils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

type Level int8

const (
	DEBUG Level = iota + 1
	INFO
	WARNING
	ERROR
)

func initLog(level zapcore.Level) {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	Logger, _ = cfg.Build()

}

func SetLogLevel(level Level) {

	switch level {
	case DEBUG:
		initLog(zapcore.DebugLevel)
		return
	case INFO:
		initLog(zapcore.InfoLevel)
		return
	case WARNING:
		initLog(zapcore.WarnLevel)
		return
	case ERROR:
		initLog(zapcore.ErrorLevel)
	default:
		initLog(zapcore.ErrorLevel)
		return
	}
}
