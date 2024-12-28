package client

import (
	"context"
	"fmt"
	"math"
	"runtime/debug"

	"github.com/xdave/keyid/interfaces"
	"github.com/xdave/keyid/models"
	"github.com/xdave/keyid/util"

	"github.com/dvcrn/go-rekordbox/rekordbox"
)

type Track struct {
	ID        string
	BPM       float64
	Scale     interfaces.Scale
	Artist    string
	Title     string
	Energy    int
	Path      string
	DateAdded string
	Tags      []string
}

func NewTrackFromContent(client *rekordbox.Client, content *rekordbox.DjmdContent) interfaces.Item {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("Could not get item", string(debug.Stack()))
		}
	}()
	bpm := float64(content.BPM.Int64Value()) / 100.0
	key, _ := client.DjmdKeyByID(context.Background(), content.KeyID)
	var scaleName string
	if key == nil {
		fmt.Println("Track has no scale", content.FileNameL)
		scaleName = "0A"
	} else {
		scaleName = key.ScaleName.String()
	}
	camelotKey := models.NewKey(scaleName)
	artist, _ := client.DjmdArtistByID(context.Background(), content.ArtistID)
	myTags, _ := client.DjmdSongMyTagByContentID(context.Background(), content.ID)

	var artistName string
	if artist == nil {
		artistName = "<none>"
	} else {
		artistName = artist.Name.String()
	}

	tags := []string{}

	for _, t := range myTags {
		tag, _ := client.DjmdMyTagByID(context.Background(), t.MyTagID)
		if tag != nil {
			tags = append(tags, tag.Name.String())
		}
	}

	return &Track{
		ID:        content.ID.String(),
		BPM:       bpm,
		Scale:     camelotKey,
		Artist:    artistName,
		Title:     content.Title.String(),
		Energy:    util.ParseEnergy(content.Commnt.String()),
		Path:      content.FolderPath.String(),
		DateAdded: content.DateCreated.String(),
		Tags:      tags,
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
	// return math.Abs(percent) <= 7
	// return math.Abs(percent) <= 2.5
	return math.Abs(percent) <= 1.8
	// return math.Abs(percent) <= 0.9
}

func (track *Track) IsCompatible(other interfaces.Item) bool {
	return track.BpmMatchesTarget(other.GetBPM()) && other.GetScale().IsCompatible(track.GetScale())
}

func (track *Track) AsBpm(targetBpm float64) interfaces.Item {
	diff := track.BPM - targetBpm
	percent := diff / targetBpm * 100.0

	newTrack := &Track{
		ID:        track.ID,
		BPM:       track.BPM,
		Scale:     track.Scale,
		Artist:    track.Artist,
		Title:     track.Title,
		Energy:    track.Energy,
		DateAdded: track.DateAdded,
		Tags:      track.Tags,
	}

	if math.Abs(percent) > 5.0 && math.Abs(percent) < 6.5 {
		newTrack.BPM = targetBpm
		if targetBpm > track.BPM {
			newTrack.Scale = newTrack.Scale.ChangeIndex(7)
		} else {
			newTrack.Scale = newTrack.Scale.ChangeIndex(-7)
		}
	}

	return newTrack
}

func (t *Track) GetPath() string {
	return t.Path
}

func (t *Track) GetDateAdded() string {
	return t.DateAdded
}

func (t *Track) GetTags() []string {
	return t.Tags
}
