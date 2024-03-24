package logger

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/xdave/keyid/args"
)

func GetLogger() fx.Option {
	if !args.CurrentArgs.Debug {
		return fx.NopLogger
	}
	return fx.WithLogger(func() fxevent.Logger {
		return &fxevent.ConsoleLogger{W: os.Stderr}
	})
}
