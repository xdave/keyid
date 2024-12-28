package interfaces

type Item interface {
	GetID() string
	GetBPM() float64
	GetScale() Scale
	GetArtist() string
	GetTitle() string
	GetEnergy() int
	Equals(other Item) bool
	String() string
	BpmMatchesTarget(targetBpm float64) bool
	IsCompatible(other Item) bool
	AsBpm(targetBpm float64) Item
	GetPath() string
	GetDateAdded() string
	GetTags() []string
}
