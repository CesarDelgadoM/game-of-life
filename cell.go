package main

type Cell struct {
	posX  int
	posY  int
	size  int
	alive bool
}

func NewCell(posx, posy, size int, alive bool) Cell {
	return Cell{
		posX:  posx,
		posY:  posy,
		size:  size,
		alive: alive,
	}
}

func InitCells(width int, height int, numCells int) [][]Cell {

	sizeCell := height / numCells
	cells := make([][]Cell, height+sizeCell)

	for c := range cells {
		cells[c] = make([]Cell, width+sizeCell)
	}

	for y := 0; y <= screenHeight; y += sizeCell {
		for x := 0; x <= screenWidth; x += sizeCell {

			cells[y][x] = NewCell(x, y, sizeCell, false)
		}
	}
	return cells
}
