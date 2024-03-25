# keyid

Uses the Camelot system to either:

- suggest the next track (based on your current Master deck)
  - Rekordbox takes 60 seconds (by default) to add the Master deck track to your history. See [here](https://assets-global.website-files.com/624b7d52289aa4ed9d117a25/6395995aaa7b424b012bfea7_Screen%20Shot%202022-12-11%20at%2012.47.54%20AM.png) for how to change this setting
- generate a new playlist from the given playlist name (or your whole collection) with the selections in a compatible key and tempo.

This assumes you have analyzed your collection using the Alphanumeric key notation.

# Prerequisites

- Install [rekordbox](https://rekordbox.com/en/)
- Set the key display format to [Alphanumeric](https://support.pioneerdj.com/hc/en-us/articles/8943219092761-Can-I-change-the-display-format-for-keys)
- Analyze your tracks (I recommend using [Mixed In Key](https://mixedinkey.com/integration/rekordbox-integration/))

# Build instructions for macOS (only tested on Sonoma)

## Dependencies

- Install the [command-line tools](https://mac.install.guide/commandlinetools/4)
- Install [homebrew](https://brew.sh/)
- Install openssl: `brew install openssl@3`
- Install [go 1.22.x](https://go.dev/doc/install)

## Build

- clone this repo: `git clone https://github.com/xdave/keyid.git`
- switch to it's directory: `cd keyid`
- install build dependencies: `go get`
- build the app: `go build .`
- run the app: `./keyid --help` with the `--help` to get usage instructions

# Build instructions for Windows

## Dependencies

- [TODO]

## Build

- [TODO]

# Roadmap

- Multiplatform downloadable builds
- Generate .m3u playlists in `generate` mode
- Better documentation
