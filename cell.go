package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Cell struct {
	width  int
	height int
	state  bool
}

func NewCell() Cell {

	return Cell{}
}

func (c Cell) DrawCell(screen *ebiten.Image, x int, y int) {

}
