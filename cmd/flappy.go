package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten"

	"github.com/Cluas/go-flappy/internal/flappy"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Flappy Gopher!!!")
	if err := ebiten.RunGame(flappy.NewGame()); err != nil {
		panic(err)
	}
}
