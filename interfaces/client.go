package interfaces

type Client interface {
	LoadPlaylist(name string) Collection
	GetNowPlaying() Item
	GetCompatibleTracks(track Item, from Collection) Collection
	Suggest(collection Collection)
	Generate(collection Collection)
	Run()
	Close()
}
