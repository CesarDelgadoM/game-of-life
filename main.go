package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

type Game struct {
	world *World
	cells []Cell
}

func (g *Game) Layout(width int, height int) (int, int) {
	return width, height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// if g.pixels == nil {
	// 	g.pixels = make([]byte, screenWidth*screenHeight*4)
	// }

	// g.world.Draw(g.pixels)
	// screen.ReplacePixels(g.pixels)
	g.cells[0].DrawCell(screen, 30, 30)
}

func main() {

	cells := make([]Cell, screenWidth*screenHeight*4)
	cells[0] = NewCell(4, 4)

	game := &Game{
		world: NewWorld(screenWidth, screenHeight, 10000),
		cells: cells,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
