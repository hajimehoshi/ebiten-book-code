package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var i = 0

func update(screen *ebiten.Image) error {
	i++
	if ebiten.IsRunningSlowly() {
		return nil
	}
	msg := fmt.Sprintf("%d", i)
	ebitenutil.DebugPrint(screen, msg)
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Increment"); err != nil {
		log.Fatal(err)
	}
}
