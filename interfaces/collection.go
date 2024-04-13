package interfaces

type Collection interface {
	Contains(other Item) bool
	Add(item Item)
	Remove(toRemove Item) Item
	MoveTo(item Item, collection Collection)
	ForEach(fn func(i Item))
	Map(mapper func(i Item) Item) Collection
	Filter(filter func(i Item) bool) Collection
	Reduce(reducer func(i Item, acc Item) Item) Item
	Len() int
	IsEmpty() bool
	First() Item
	Last() Item
	Get(index int) Item
	IndexOf(item Item) int
	Find(predicate func(i Item) bool) Item
	FindIndex(predicate func(i Item) bool) int
	SortWith(comparator func(i, j Item) bool) Collection
	Items() []Item
	RandomShuffle()
}
