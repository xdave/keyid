package printer

import (
	"fmt"

	"github.com/xdave/keyid/interfaces"
)

type CliPrinter struct {
}

func NewCliPrinter() interfaces.Printer {
	return &CliPrinter{}
}

func (c *CliPrinter) PrintHeader() {}

func (c *CliPrinter) Print(track interfaces.Item) {
	fmt.Println(track)
}
