package interfaces

type Client interface {
	LoadPlaylist(name string) Collection
	GetTrackByTitle(pattern string, from Collection) Item
	GetNowPlaying(collection Collection) Item
	GetCompatibleTracks(track Item, from Collection) Collection
	Suggest(collection Collection)
	Generate(collection Collection)
	Run()
	Close()
}
