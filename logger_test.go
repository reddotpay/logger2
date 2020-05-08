package logger2_test

import (
	"fmt"
	"runtime/debug"
	"testing"
	"time"

	"github.com/reddotpay/logger2"
	"go.uber.org/zap"
)

func TestNew_Production(t *testing.T) {
	logger2.Client = logger2.NewProduction()
	logger2.Client.Info("hi")
}
func TestNew_Local(t *testing.T) {
	logger2.Client = logger2.NewLocal()
	logger2.Client.Info("hi")
	logger2.Client.Debug("hi")
}
func TestLoggerDebugLevel(t *testing.T) {
	logger2.InitStackTrace()

	logger2.Client = logger2.NewLocal().With(
		zap.Any("product", "testP"),
		zap.String("createAt", time.Now().UTC().Format(time.RFC1123)),
	)
	defer func() {
		logger2.Client.Sync()
	}()

	doSomeStuff()
	logger2.Client.Info("summary of TestLoggerDebugLevel", zap.Any("stackTraceArray", logger2.STA))
}
func TestLoggerProductionLevel(t *testing.T) {
	logger2.InitStackTrace()

	logger2.Client = logger2.NewProduction().With(
		zap.Any("product", "testP"),
		zap.String("createAt", time.Now().UTC().Format(time.RFC1123)),
	)
	defer func() {
		logger2.Client.Sync()
	}()

	doSomeStuff()
	logger2.Client.Info("summary of TestLoggerDebugLevel", zap.Any("stackTraceArray", logger2.STA))
}

func doSomeStuff() {

	fmt.Println(string(debug.Stack()))

	logger2.Client.Debug("the debug stuff is hererererererererre", zap.String("d", "debugggggg"))
	logger2.Info(
		map[string]interface{}{
			"ValA": "A",
		},
		logger2.WhereAmI())
}
