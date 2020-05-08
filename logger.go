package logger2

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// global variable declaration
var (
	Client Logger
	STA    StackTraceArray
)

// Logger contains necessary methods of zap
type Logger interface {
	Info(key string, fields ...zapcore.Field)
	Error(key string, fields ...zapcore.Field)
	Debug(key string, fields ...zapcore.Field)
	Core() zapcore.Core
	Sync() error
	WithOptions(opts ...zap.Option) *zap.Logger
}

// NewProduction returns new Production level logger
func NewProduction() *zap.Logger {
	return new()
}

// NewLocal returns a debug level logger
func NewLocal() *zap.Logger {
	return new(zapcore.DebugLevel)
}

// new defines a logger that complies with rdp v2 payload convention
func new(level ...zapcore.LevelEnabler) *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "summary",
		LevelKey:       "level",
		NameKey:        "logger",
		StacktraceKey:  "stackTraceArray",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		TimeKey:        "sortDate",
	}
	var zcLevel zapcore.LevelEnabler
	if level == nil {
		zcLevel = zapcore.InfoLevel
	} else {
		zcLevel = level[0]
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, zcLevel)

	return zap.New(core)
}

// InitStackTrace intializes STA
func InitStackTrace() {
	STA = StackTraceArray{}
}

// AppendStackTrace ...
func appendStackTrace(stackType string, detail interface{}, caller string) {
	STA = append(STA, &StackTrace{
		Type:   stackType,
		Detail: detail,
		Caller: []string{caller},
	})
}

// Error appends error to stack trace array
func Error(detail interface{}, caller string) {
	appendStackTrace("error", detail, caller)
}

// Info appends error to stack trace array
func Info(detail interface{}, caller string) {
	appendStackTrace("info", detail, caller)
}
