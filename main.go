package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	screenWidth  = 800
	screenHeight = 800
	numCells     = 100
)

func main() {
	world := NewWorld(screenWidth, screenHeight, numCells)
	world.generateCellsAlive()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game of life 2.0")

	if err := ebiten.RunGame(world); err != nil {
		panic(err)
	}
}
