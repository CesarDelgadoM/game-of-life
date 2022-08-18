package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 800
	sizeCells    = 80
	density      = 10
)

func main() {

	world := NewWorld(screenWidth, screenHeight, sizeCells, density)
	world.GenerateLiveCells()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(world); err != nil {
		panic(err)
	}
}
