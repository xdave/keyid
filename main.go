package main

import (
	"github.com/xdave/keyid/app"
	"github.com/xdave/keyid/args"
	"github.com/xdave/keyid/interfaces"
	"github.com/xdave/keyid/logger"
	"github.com/xdave/keyid/printer"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		logger.GetLogger(),
		args.Module,
		printer.Module,
		app.Module,
		fx.Invoke(interfaces.Client.Run),
	).Run()
}
