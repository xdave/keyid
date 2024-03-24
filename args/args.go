package args

import (
	"flag"

	"github.com/xdave/keyid/interfaces"
)

type Args struct {
	Mode     interfaces.Mode
	Playlist string
	Debug    bool
}

func NewArgs() *Args {
	args := &Args{}

	args.Parse()

	return args
}

func (a *Args) Parse() {
	flag.StringVar(&a.Mode, "mode", "suggest", "One of 'suggest' or 'generate'")
	flag.StringVar(&a.Playlist, "playlist", "", "Name of Rekordbox Playlist to use (uses whole collection by default)")
	flag.BoolVar(&a.Debug, "debug", false, "Enable debug logging")

	flag.Parse()
	a.Validate()
}

func (a *Args) Validate() {
	// TODO
}
