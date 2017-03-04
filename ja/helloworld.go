package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello"); err != nil {
		log.Fatal(err)
	}
}
