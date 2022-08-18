package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type World struct {
	cells       [][]Cell
	width       int
	height      int
	numberCells int
}

func NewWorld(screenWidth int, screenHeight int, numCells int) *World {

	cells := InitCells(screenWidth, screenHeight, numCells)

	return &World{
		cells:       cells,
		width:       screenWidth,
		height:      screenHeight,
		numberCells: numCells,
	}
}

func (w *World) generateCellsAlive() {

	rand.Seed(time.Now().UnixNano())
	alive := false

	for y := 0; y <= w.height; y++ {
		for x := 0; x <= w.width; x++ {

			if rand.Intn(6) == 1 {
				alive = true
			} else {
				alive = false
			}

			w.cells[y][x].alive = alive
		}
	}
}

func (w *World) Layout(width int, height int) (int, int) {
	return width, height
}

func (w *World) Update() error {

	w.executeSimulation()

	return nil
}

func (w *World) executeSimulation() {

	newCells := NewWorld(w.width, w.height, w.numberCells).cells

	for y := 0; y <= w.height; y++ {
		for x := 0; x <= w.width; x++ {

			cell := w.cells[y][x]
			neighbours := w.countNeighbours(cell.posX, cell.posY)

			if cell.alive && neighbours < 2 {
				newCells[cell.posY][cell.posX].alive = false
			}

			if cell.alive && (neighbours == 2 || neighbours == 3) {
				newCells[cell.posY][cell.posX].alive = true
			}

			if cell.alive && neighbours > 3 {
				newCells[cell.posY][cell.posX].alive = false
			}

			if !cell.alive && neighbours == 3 {
				newCells[cell.posY][cell.posX].alive = true
			}
		}
	}
	w.cells = newCells
}

func (w *World) countNeighbours(x int, y int) int {

	sizeCell := w.cells[0][0].size
	directions := []int{0, sizeCell, -sizeCell}
	count := 0

	for _, dirX := range directions {
		for _, dirY := range directions {

			if dirX == 0 && dirY == 0 {
				continue
			}

			if x+dirX >= 0 && x+dirX < w.width+sizeCell &&
				y+dirY >= 0 && y+dirY < w.height+sizeCell &&
				w.cells[y+dirY][x+dirX].alive {

				count++
			}
		}
	}
	return count
}

func (w *World) Draw(screen *ebiten.Image) {

	for y := range w.cells {
		for x, c := range w.cells[y] {

			if c.alive {
				ebitenutil.DrawRect(screen, float64(x), float64(y),
					float64(c.size), float64(c.size), color.White)
			}
		}
	}
}
