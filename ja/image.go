package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	if err := screen.Fill(color.RGBA{0xff, 0, 0, 0xff}); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Image"); err != nil {
		log.Fatal(err)
	}
}
