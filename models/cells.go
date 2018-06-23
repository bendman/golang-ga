package models

type Cell int

const (
	Empty Cell = iota
	Can
	Wall
)

var Cells = [...]Cell{Empty, Can, Wall}

// Return a specific string based on cell type
func (cell Cell) String() string {
	switch cell {
	case Wall:
		return "X"
	case Can:
		return ","
	default:
		return " "
	}
}
