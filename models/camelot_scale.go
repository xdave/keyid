package models

import (
	"fmt"

	"github.com/xdave/keyid/interfaces"
)

type CamelotScale struct {
	Index int
	Kind  interfaces.ScaleKind
}

type ScaleTransition func() interfaces.Scale

func NewKey(key string) interfaces.Scale {
	return ParseCamelotKey(key)
}

func (key *CamelotScale) GetIndex() int {
	return key.Index
}

func (key *CamelotScale) GetKind() interfaces.ScaleKind {
	return key.Kind
}

func (key *CamelotScale) ChangeIndex(n int) interfaces.Scale {
	return &CamelotScale{
		Index: interfaces.ModCyclic(key.Index+n, 12),
		Kind:  key.Kind,
	}
}

func (key *CamelotScale) SwapKind() interfaces.Scale {
	return &CamelotScale{Index: key.Index, Kind: key.Kind.Swap()}
}

func (key *CamelotScale) String() string {
	return fmt.Sprintf("%d%s", key.Index, key.Kind.String())
}

func (key *CamelotScale) Horizontal(direction int) interfaces.Scale {
	return key.ChangeIndex(direction)
}

func (key *CamelotScale) Vertical() interfaces.Scale {
	return key.SwapKind()
}

func (key *CamelotScale) Diagonal() interfaces.Scale {
	if key.Kind == interfaces.Major {
		return key.ChangeIndex(1).SwapKind()
	}
	return key.ChangeIndex(-1).SwapKind()
}

// Not sure about this one
func (key *CamelotScale) FlatToMinor() interfaces.Scale {
	if key.Kind == interfaces.Minor {
		return key.ChangeIndex(-4).SwapKind()
	}
	return key.ChangeIndex(4).SwapKind()
}

func (key *CamelotScale) MajorToMinor() interfaces.Scale {
	if key.Kind == interfaces.Minor {
		return key.ChangeIndex(3).SwapKind()
	}
	return key.ChangeIndex(-3).SwapKind()
}

func (key *CamelotScale) IsEqual(other interfaces.Scale) bool {
	return key.Index == other.GetIndex() && key.Kind == other.GetKind()
}

func (key *CamelotScale) IsCompatible(other interfaces.Scale) bool {
	return key.IsEqual(other) ||
		key.IsEqual(other.MajorToMinor()) ||
		key.IsEqual(other.Horizontal(1)) ||
		key.IsEqual(other.Horizontal(-1)) ||
		key.IsEqual(other.Diagonal()) ||
		key.IsEqual(other.Vertical()) // ||
	// key.IsEqual(other.FlatToMinor())
}
