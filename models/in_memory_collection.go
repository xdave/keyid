package models

import (
	"slices"
	"sort"

	"github.com/xdave/keyid/interfaces"
)

type InMemoryCollection struct {
	items []interfaces.Item
}

func NewInMemoryCollection(items ...interfaces.Item) interfaces.Collection {
	return &InMemoryCollection{
		items: items,
	}
}

func (c *InMemoryCollection) Contains(item interfaces.Item) bool {
	return slices.ContainsFunc(c.items, func(i interfaces.Item) bool {
		return item.Equals(i)
	})
}

func (c *InMemoryCollection) Add(item interfaces.Item) {
	if !c.Contains(item) {
		c.items = append(c.items, item)
	}
}

func (c *InMemoryCollection) Remove(toRemove interfaces.Item) interfaces.Item {
	if !c.Contains(toRemove) {
		return toRemove
	}

	newItems := []interfaces.Item{}

	for _, item := range c.items {
		if item.Equals(toRemove) {
			newItems = append(newItems, item)
		}
	}
	c.items = newItems

	return toRemove
}

func (c *InMemoryCollection) MoveTo(item interfaces.Item, collection interfaces.Collection) {
	collection.Add(item)
	c.Remove(item)
}

func (c *InMemoryCollection) ForEach(fn func(i interfaces.Item)) {
	for _, item := range c.items {
		fn(item)
	}
}

func (c *InMemoryCollection) Map(mapper func(i interfaces.Item) interfaces.Item) interfaces.Collection {
	newItems := []interfaces.Item{}
	for _, item := range c.items {
		newItems = append(newItems, mapper(item))
	}
	return NewInMemoryCollection(newItems...)
}

func (c *InMemoryCollection) Filter(filter func(i interfaces.Item) bool) interfaces.Collection {
	newItems := []interfaces.Item{}
	for _, item := range c.items {
		if filter(item) {
			newItems = append(newItems, item)
		}
	}
	return NewInMemoryCollection(newItems...)
}

func (c *InMemoryCollection) Reduce(reducer func(i interfaces.Item, acc interfaces.Item) interfaces.Item) interfaces.Item {
	acc := c.items[0]
	for _, item := range c.items[1:] {
		acc = reducer(item, acc)
	}
	return acc
}

func (c *InMemoryCollection) Len() int {
	return len(c.items)
}

func (c *InMemoryCollection) IsEmpty() bool {
	return c.Len() == 0
}

func (c *InMemoryCollection) First() interfaces.Item {
	if c.IsEmpty() {
		return nil
	}
	return c.items[0]
}

func (c *InMemoryCollection) Last() interfaces.Item {
	if c.IsEmpty() {
		return nil
	}
	return c.items[len(c.items)-1]
}

func (c *InMemoryCollection) Get(index int) interfaces.Item {
	if index < 0 || index >= len(c.items) {
		return nil
	}
	return c.items[index]
}

func (c *InMemoryCollection) IndexOf(item interfaces.Item) int {
	for i, v := range c.items {
		if item.Equals(v) {
			return i
		}
	}
	return -1
}

func (c *InMemoryCollection) Find(predicate func(i interfaces.Item) bool) interfaces.Item {
	for _, item := range c.items {
		if predicate(item) {
			return item
		}
	}
	return nil
}

func (c *InMemoryCollection) FindIndex(predicate func(i interfaces.Item) bool) int {
	for i, v := range c.items {
		if predicate(v) {
			return i
		}
	}
	return -1
}

func (c *InMemoryCollection) SortWith(comparator func(i, j interfaces.Item) bool) interfaces.Collection {
	newCollection := &InMemoryCollection{}

	for _, item := range c.items {
		newCollection.Add(item)
	}

	sort.Slice(newCollection.items, func(i, j int) bool {
		return comparator(newCollection.Get(i), newCollection.Get(j))
	})

	return newCollection
}

func (c *InMemoryCollection) Items() []interfaces.Item {
	return c.items
}
