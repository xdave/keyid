# keyid

Uses the Camelot system to either:

- suggest the next track (based on your current Master deck)
  - Rekordbox takes 60 seconds (by default) to add the Master deck track to your history
- generate a new playlist from the given playlist name (or your whole collection) with the selections in a compatible key and tempo.

This assumes you have analyzed your collection using the Alphanumeric key notation.

# Build instructions for macOS (only tested on Sonoma)

- Install [rekordbox](https://rekordbox.com/en/)
- Set the key display format to [Alphanumeric](https://support.pioneerdj.com/hc/en-us/articles/8943219092761-Can-I-change-the-display-format-for-keys)
- Analyze your tracks (I recommend using [Mixed In Key](https://mixedinkey.com/integration/rekordbox-integration/))
- Install the [command-line tools](https://mac.install.guide/commandlinetools/4)
- Install [homebrew](https://brew.sh/)
- Install openssl: `brew install openssl@3`
- Install [go 1.22.x](https://go.dev/doc/install)
- clone this repo: `git clone https://github.com/xdave/keyid`
- switch to it's directory: `cd keyid`
- build the app: `go build .`
- run the app: `./keyid --help` with the `--help` to get usage instructions

# Build instructions for Windows

- [TODO]
