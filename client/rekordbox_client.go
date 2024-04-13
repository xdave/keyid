package client

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/xdave/keyid/args"
	"github.com/xdave/keyid/interfaces"
	"github.com/xdave/keyid/models"

	"github.com/dvcrn/go-rekordbox/rekordbox"
	"github.com/mattn/go-nulltype"
	"go.uber.org/fx"
)

type RekordboxClient struct {
	client          *rekordbox.Client
	optionsResolver *RekordboxOptionsResolver
	history         *RekordboxHistory
	args            *args.Args
	shutdowner      fx.Shutdowner
}

type RekordboxClientParams struct {
	fx.In
	fx.Lifecycle
	fx.Shutdowner
	OptionsResolver *RekordboxOptionsResolver
	History         *RekordboxHistory
	Args            *args.Args
}

type RekordboxClientResult struct {
	fx.Out
	Client interfaces.Client
}

func NewRekordboxClient(params RekordboxClientParams) RekordboxClientResult {
	optionsFilePath := params.OptionsResolver.Resolve()

	client, err := rekordbox.NewClient(optionsFilePath)
	if err != nil {
		panic(err)
	}

	rbClient := &RekordboxClient{
		client:          client,
		optionsResolver: params.OptionsResolver,
		history:         params.History,
		args:            params.Args,
		shutdowner:      params.Shutdowner,
	}

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			rbClient.Close()
			return nil
		},
	})

	return RekordboxClientResult{
		Client: rbClient,
	}
}

func (c *RekordboxClient) LoadPlaylist(name string) interfaces.Collection {
	tracks := []interfaces.Item{}

	if c.args.Playlist == "" {
		allContent, _ := c.client.AllDjmdContent(context.Background())
		for _, content := range allContent {
			tracks = append(tracks, NewTrackFromContent(c.client, content))
		}
	} else {
		playlists, _ := c.client.DjmdPlaylistByName(context.Background(), nulltype.NullStringOf(name))
		if len(playlists) == 0 {
			err := fmt.Sprintf("Error: cannot find a playlist with name '%s'", name)
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		playlist := playlists[0]
		playlistSongs, _ := c.client.DjmdSongPlaylistByPlaylistID(context.Background(), playlist.ID)
		sort.Slice(playlistSongs, func(i, j int) bool {
			return playlistSongs[i].TrackNo.Int64Value() < playlistSongs[j].TrackNo.Int64Value()
		})
		for _, song := range playlistSongs {
			content, _ := c.client.DjmdContentByID(context.Background(), song.ContentID)
			tracks = append(tracks, NewTrackFromContent(c.client, content))
		}
	}

	return models.NewInMemoryCollection(tracks...)
}

func (c *RekordboxClient) GetTrackByTitle(pattern string, from interfaces.Collection) interfaces.Item {
	for _, track := range from.Items() {
		if strings.Contains(strings.ToLower(track.GetTitle()), strings.ToLower(pattern)) {
			return track
		}
	}

	err := fmt.Sprintf("Error: cannot find a track with '%s' in the name", c.args.StartWith)
	fmt.Fprintln(os.Stderr, err)
	return nil
}

func (c *RekordboxClient) GetNowPlaying(collection interfaces.Collection) interfaces.Item {
	if c.args.StartWith != "" {
		return c.GetTrackByTitle(c.args.StartWith, collection)
	}

	songHistories, _ := c.client.RecentDjmdSongHistory(context.Background(), 1)
	if len(songHistories) == 0 {
		return nil
	}
	item := songHistories[0]
	content, err := c.client.DjmdContentByID(context.Background(), item.ContentID)

	if err != nil {
		panic(err)
	}

	track := NewTrackFromContent(c.client, content)
	c.history.Add(track)
	return track
}

func (c *RekordboxClient) GetCompatibleTracks(track interfaces.Item, from interfaces.Collection) interfaces.Collection {
	compat := models.NewInMemoryCollection()

	from.ForEach(func(item interfaces.Item) {
		if !c.history.Contains(item) {
			if !item.Equals(track) && track.IsCompatible(item.AsBpm(track.GetBPM())) {
				compat.Add(item)
			}
		}
	})

	return compat
}

func (c *RekordboxClient) Run() {
	collection := c.LoadPlaylist(c.args.Playlist)

	if collection == nil {
		c.shutdowner.Shutdown(fx.ExitCode(1))
		return
	}

	if c.args.Mode == interfaces.ModeSuggest {
		c.Suggest(collection)
	} else if c.args.Mode == interfaces.ModeGenerate {
		c.Generate(collection)
	}
}

func (c *RekordboxClient) Suggest(collection interfaces.Collection) {
	track := c.GetNowPlaying(collection)

	if track == nil {
		c.shutdowner.Shutdown(fx.ExitCode(1))
		return
	}

	fmt.Println(track)
	compat := c.GetCompatibleTracks(track, collection)
	compat.ForEach(func(item interfaces.Item) {
		fmt.Println(" ->", item)
	})
	c.shutdowner.Shutdown(fx.ExitCode(0))
}

func (c *RekordboxClient) Generate(collection interfaces.Collection) {
	crate := models.NewInMemoryCollection(collection.Items()...)

	if c.args.Random {
		crate.RandomShuffle()
	}

	startWith := c.GetTrackByTitle(c.args.StartWith, crate)

	if startWith == nil {
		c.shutdowner.Shutdown(fx.ExitCode(1))
		return
	}

	playlist := models.NewInMemoryCollection(startWith)

	var lastTrack interfaces.Item

	for {
		lastTrack = playlist.Last()

		compatible := c.GetCompatibleTracks(lastTrack, crate)
		if compatible.IsEmpty() || playlist.Len() == crate.Len() {
			break
		}
		hasCompatibleTrack := false
		for _, track := range compatible.Items() {
			if playlist.Contains(track) {
				continue
			}
			hasCompatibleTrack = true
			playlist.Add(track)
			break

		}
		if !hasCompatibleTrack {
			break
		}
	}

	playlist.ForEach(func(track interfaces.Item) {
		fmt.Println(track)
	})

	c.shutdowner.Shutdown(fx.ExitCode(0))
}

func (c *RekordboxClient) Close() {
	c.client.Close()
}
