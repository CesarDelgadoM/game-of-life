package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type World struct {
	cells    [][]bool
	sizeCell int
	width    int
	height   int
	density  int
}

func NewWorld(screenWidth int, screenHeight int, numberCells int, density int) *World {

	sizeCell := screenHeight / numberCells

	cells := make([][]bool, screenHeight+sizeCell)
	for i := range cells {
		cells[i] = make([]bool, screenWidth+sizeCell)
	}

	return &World{
		cells:    cells,
		width:    screenWidth,
		height:   screenHeight,
		sizeCell: sizeCell,
		density:  density,
	}
}

func (w *World) GenerateLiveCells() {

	rand.Seed(time.Now().UnixNano())

	for y := 0; y <= w.height; y += w.sizeCell {
		for x := 0; x <= w.width; x += w.sizeCell {

			if rand.Intn(w.density) == 1 {
				w.cells[y][x] = true
			}
		}
	}
}

func (w *World) Layout(width int, height int) (int, int) {
	return width, height
}

func (w *World) Update() error {

	w.executeRules()

	return nil
}

func (w *World) executeRules() {

	newCells := NewWorld(w.width, w.height, (w.height / w.sizeCell), w.density).cells

	for y := 0; y <= w.height; y += w.sizeCell {
		for x := 0; x <= w.width; x += w.sizeCell {

			count := w.countNeighbours(x, y)
			alive := w.cells[y][x]

			if alive && count < 2 {
				newCells[y][x] = false
			}

			if alive && (count == 2 || count == 3) {
				newCells[y][x] = true
			}

			if alive && count > 3 {
				newCells[y][x] = false
			}

			if !alive && count == 3 {
				newCells[y][x] = true
			}
		}
	}

	w.cells = newCells
}

func (w *World) countNeighbours(x int, y int) int {

	directions := []int{0, w.sizeCell, -w.sizeCell}
	count := 0

	for _, dx := range directions {
		for _, dy := range directions {

			if dx == 0 && dy == 0 {
				continue
			}

			if w.isBounds(x+dx, y+dy) && w.cells[y+dy][x+dx] == true {
				count++
			}
		}
	}

	return count
}

func (w *World) isBounds(x int, y int) bool {

	xOK := x >= 0 && x < w.width+w.sizeCell
	yOK := y >= 0 && y < w.height+w.sizeCell

	return xOK && yOK
}

func (w *World) Draw(screen *ebiten.Image) {

	for y := range w.cells {
		for x := range w.cells[y] {

			if w.cells[y][x] {
				ebitenutil.DrawRect(screen, float64(x), float64(y),
					float64(w.sizeCell), float64(w.sizeCell), color.White)
			}
		}
	}
}
