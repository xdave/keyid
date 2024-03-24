package client

import (
	"context"
	"fmt"
	"math"

	"github.com/xdave/keyid/interfaces"
	"github.com/xdave/keyid/models"
	"github.com/xdave/keyid/util"

	"github.com/dvcrn/go-rekordbox/rekordbox"
)

type Track struct {
	ID     string
	BPM    float64
	Scale  interfaces.Scale
	Artist string
	Title  string
	Energy int
}

func NewTrackFromContent(client *rekordbox.Client, content *rekordbox.DjmdContent) interfaces.Item {
	bpm := float64(content.BPM.Int64Value()) / 100.0
	key, _ := client.DjmdKeyByID(context.Background(), content.KeyID)
	camelotKey := models.NewKey(key.ScaleName.String())
	artist, _ := client.DjmdArtistByID(context.Background(), content.ArtistID)

	return &Track{
		ID:     content.ID.String(),
		BPM:    bpm,
		Scale:  camelotKey,
		Artist: artist.Name.String(),
		Title:  content.Title.String(),
		Energy: util.ParseEnergy(content.Commnt.String()),
	}
}

func (t *Track) GetID() string {
	return t.ID
}

func (t *Track) Equals(other interfaces.Item) bool {
	return other.
		GetID() == t.GetID()
}

func (t *Track) GetBPM() float64 {
	return t.BPM
}

func (t *Track) GetScale() interfaces.Scale {
	return t.Scale
}

func (t *Track) GetArtist() string {
	return t.Artist
}

func (t *Track) GetTitle() string {
	return t.Title
}

func (t *Track) GetEnergy() int {
	return t.Energy
}

func (track *Track) String() string {
	return fmt.Sprintf(
		"%d\t%s\t%d\t%s - %s",
		int64(track.BPM),
		track.Scale.String(),
		track.Energy,
		track.Artist,
		track.Title,
	)
}

func (track *Track) BpmMatchesTarget(targetBpm float64) bool {
	diff := track.BPM - targetBpm
	percent := diff / targetBpm * 100.0
	return math.Abs(percent) <= 3.0
}

func (track *Track) IsCompatible(other interfaces.Item) bool {
	return track.BpmMatchesTarget(other.GetBPM()) && track.Scale.IsCompatible(other.GetScale())
}

func (track *Track) AsBpm(targetBpm float64) interfaces.Item {
	diff := track.BPM - targetBpm
	percent := diff / targetBpm * 100.0

	newTrack := &Track{
		ID:     track.ID,
		BPM:    track.BPM,
		Scale:  track.Scale,
		Artist: track.Artist,
		Title:  track.Title,
		Energy: track.Energy,
	}

	if math.Abs(percent) > 4.0 && math.Abs(percent) < 7.0 {
		newTrack.BPM = targetBpm
		newTrack.Scale = newTrack.Scale.ChangeIndex(7)
	}

	return newTrack
}
