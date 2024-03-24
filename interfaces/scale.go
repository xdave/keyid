package interfaces

type Scale interface {
	GetIndex() int
	GetKind() ScaleKind
	ChangeIndex(amount int) Scale
	SwapKind() Scale
	String() string
	Vertical() Scale
	Horizontal(direction int) Scale
	Diagonal() Scale
	FlatToMinor() Scale
	MajorToMinor() Scale
	IsEqual(other Scale) bool
	IsCompatible(other Scale) bool
}

func ModCyclic(num, modulus int) int {
	index := ((num % modulus) + modulus) % modulus
	if index == 0 {
		index = modulus
	}
	return index
}
