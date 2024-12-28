package util

import "slices"

type StringSlice []string

func (s StringSlice) ContainsAnyOf(other []string) bool {
	for _, item := range other {
		return slices.Contains(s, item)
	}
	return false
}
