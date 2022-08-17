package main

import "math/rand"

type World struct {
	area   []bool
	width  int
	height int
}

func NewWorld(width int, height int, maxLiveCells int) *World {

	world := &World{
		area:   make([]bool, width*height),
		width:  width,
		height: height,
	}
	world.init(maxLiveCells)

	return world
}

func (w *World) init(maxLiveCells int) {

	for i := 0; i < maxLiveCells; i++ {

		x := rand.Intn(w.width)
		y := rand.Intn(w.height)

		w.area[y*w.width+x] = true
	}
}

func (w *World) Draw(pixels []byte) {

	for i, value := range w.area {
		if value {
			pixels[4*i] = 0xff //255//color blanco
			pixels[4*i+1] = 0xff
			pixels[4*i+2] = 0xff
			pixels[4*i+3] = 0xff
		} else {
			pixels[4*i] = 0
			pixels[4*i+1] = 0
			pixels[4*i+2] = 0
			pixels[4*i+3] = 0
		}
	}
}
