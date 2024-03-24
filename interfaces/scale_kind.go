package interfaces

import "fmt"

type ScaleKind rune

const (
	Minor ScaleKind = 'A'
	Major ScaleKind = 'B'
)

func (kind ScaleKind) Swap() ScaleKind {
	if kind == Minor {
		return Major
	}
	return Minor
}

func (kind ScaleKind) String() string {
	return fmt.Sprintf("%c", kind)
}
