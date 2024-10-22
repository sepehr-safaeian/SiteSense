package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CustomTimeEncoder formats the time in a custom way
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// CustomCallerEncoder adds file and line number details
func CustomCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%s:%d", caller.File, caller.Line))
}

// GetLogLevelEmoji returns an emoji based on the log level
func GetLogLevelEmoji(level zapcore.Level) string {
	switch level {
	case zapcore.DebugLevel:
		return "🐛"
	case zapcore.InfoLevel:
		return "ℹ️"
	case zapcore.WarnLevel:
		return "⚠️"
	case zapcore.ErrorLevel:
		return "❌"
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		return "💥"
	default:
		return "❔"
	}
}

// CustomEncoderConfig defines the custom log format
func CustomEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeCaller:   CustomCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
}

// NewCustomLogger creates a new instance of a custom logger
func NewCustomLogger() *zap.Logger {
	config := zap.Config{
		Encoding:         "console",
		EncoderConfig:    CustomEncoderConfig(),
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Create a custom core to add emoji to the log messages
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config.EncoderConfig),
		zapcore.AddSync(zapcore.Lock(os.Stdout)),
		config.Level,
	)

	return zap.New(core, zap.AddCaller(), zap.Development(), zap.AddStacktrace(zap.ErrorLevel)).WithOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(c, zapcore.NewCore(zapcore.NewConsoleEncoder(config.EncoderConfig), zapcore.AddSync(os.Stderr), zap.ErrorLevel))
	}))
}
