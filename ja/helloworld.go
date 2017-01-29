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
	if err := ebitenutil.DebugPrint(screen, "Hello, World!"); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello"); err != nil {
		log.Fatal(err)
	}
}
