package printer

import (
	"fmt"

	"github.com/xdave/keyid/interfaces"
)

type M3uPrinter struct {
}

func NewM3uPrinter() interfaces.Printer {
	return &M3uPrinter{}
}

func (c *M3uPrinter) PrintHeader() {
	fmt.Println("#EXTM3U")
	fmt.Println("")
}

func (c *M3uPrinter) Print(track interfaces.Item) {
	fmt.Println("#EXTINF:-1,", track.GetArtist(), "-", track.GetTitle())
	fmt.Println(track.GetPath())
}
