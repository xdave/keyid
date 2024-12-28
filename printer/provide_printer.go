package printer

import (
	"github.com/xdave/keyid/args"
	"github.com/xdave/keyid/interfaces"
)

func ProvidePrinter(args *args.Args) interfaces.Printer {
	if args.Mode == interfaces.ModeGenerate && args.M3U {
		return NewM3uPrinter()
	}
	return NewCliPrinter()
}
