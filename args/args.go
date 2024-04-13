package args

import (
	"flag"

	"github.com/xdave/keyid/interfaces"
)

type Args struct {
	Mode      interfaces.Mode
	StartWith string
	Playlist  string
	Random    bool
	Debug     bool
}

func NewArgs() *Args {
	args := &Args{}

	args.Parse()

	return args
}

func (a *Args) Parse() {
	flag.StringVar(&a.Mode, "mode", "suggest", "One of 'suggest' or 'generate'")
	flag.StringVar(&a.StartWith, "startWith", "", "Some part of the Track Title to start with in 'generate' mode (otherwise starts with first track in provided 'playlist')")
	flag.StringVar(&a.Playlist, "playlist", "", "Name of Rekordbox Playlist to use (uses whole collection by default)")
	flag.BoolVar(&a.Random, "random", false, "Randomize playlist before 'generate'")
	flag.BoolVar(&a.Debug, "debug", false, "Enable debug logging")

	flag.Parse()
	a.Validate()
}

func (a *Args) Validate() {
	// TODO
}
