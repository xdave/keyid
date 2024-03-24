package client

import (
	"github.com/xdave/keyid/interfaces"
	"github.com/xdave/keyid/models"

	"go.uber.org/fx"
)

type RekordboxHistory struct {
	tracks interfaces.Collection
}

type RekordboxHistoryParams struct {
	fx.In
}

type RekordboxHistoryResult struct {
	fx.Out
	History *RekordboxHistory
}

func NewRekordboxHistory(params RekordboxHistoryParams) RekordboxHistoryResult {
	return RekordboxHistoryResult{
		History: &RekordboxHistory{
			tracks: models.NewInMemoryCollection(),
		},
	}
}

func (h *RekordboxHistory) Add(track interfaces.Item) {
	h.tracks.Add(track)
}

func (h *RekordboxHistory) Contains(track interfaces.Item) bool {
	return h.tracks.Contains(track)
}
