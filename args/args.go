package args

import (
	"flag"
	"fmt"
	"os"

	"github.com/xdave/keyid/interfaces"
)

type Args struct {
	Mode      interfaces.Mode
	StartWith string
	Playlist  string
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
	flag.BoolVar(&a.Debug, "debug", false, "Enable debug logging")

	flag.Parse()
	a.Validate()
}

func (a *Args) Validate() {
	if a.Mode == interfaces.ModeGenerate && a.StartWith == "" {
		fmt.Fprintln(os.Stderr, "Error: must provide --startWith in 'generate' mode")
		flag.Usage()
		os.Exit(1)
	}
}
