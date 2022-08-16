package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct{}

func (g *Game) Layout(width int, height int) (int, int) {
	return width, height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, 0, 0, screenWidth/2, screenHeight/2, color.White)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
