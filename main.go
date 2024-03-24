package main

// #cgo CFLAGS: -I/opt/homebrew/include
// #cgo LDFLAGS: -L/opt/homebrew/lib
import "C"

import (
	"github.com/xdave/keyid/app"
	"github.com/xdave/keyid/args"
	"github.com/xdave/keyid/interfaces"
	"github.com/xdave/keyid/logger"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		logger.GetLogger(),
		args.Module,
		app.Module,
		fx.Invoke(func(c interfaces.Client) {
			c.Run()
			defer c.Close()
		}),
	).Run()
}
