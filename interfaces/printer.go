package interfaces

type Printer interface {
	PrintHeader()
	Print(track Item)
}
