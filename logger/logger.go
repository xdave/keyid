package logger

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/xdave/keyid/args"
)

func GetLogger() fx.Option {
	return fx.WithLogger(func(args *args.Args) fxevent.Logger {
		if args.Debug {
			return &fxevent.ConsoleLogger{W: os.Stderr}
		}
		return fxevent.NopLogger
	})
}
