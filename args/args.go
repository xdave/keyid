package args

import (
	"flag"

	"github.com/xdave/keyid/interfaces"
)

type Args struct {
	Mode        interfaces.Mode
	From        string
	StartWith   string
	Tags        string
	ExcludeTags string
	Playlist    string
	Random      bool
	M3U         bool
	Debug       bool
}

func NewArgs() *Args {
	args := &Args{}

	args.Parse()

	return args
}

func (a *Args) Parse() {
	flag.StringVar(&a.Mode, "mode", "suggest", "One of 'suggest' or 'generate'")
	flag.StringVar(&a.From, "from", "1970-01-01", "Only look at tracks newer than this date")
	flag.StringVar(&a.StartWith, "startWith", "", "Some part of the Track Title to start with in 'generate' mode (otherwise starts with first track in provided 'playlist')")
	flag.StringVar(&a.Tags, "tags", "", "Only include tracks that match the given tags (comma-separated)")
	flag.StringVar(&a.ExcludeTags, "excludeTags", "", "Exclude tracks that match the given tags (comma-separated)")
	flag.StringVar(&a.Playlist, "playlist", "", "Name of Rekordbox Playlist to use (uses whole collection by default)")
	flag.BoolVar(&a.Random, "random", false, "Randomize playlist before 'generate'")
	flag.BoolVar(&a.M3U, "m3u", false, "Generate an M3U playlist in 'generate' mode")
	flag.BoolVar(&a.Debug, "debug", false, "Enable debug logging")

	flag.Parse()
	a.Validate()
}

func (a *Args) Validate() {
	// TODO
}
