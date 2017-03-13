package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// カウンター
var i = 0

func update(screen *ebiten.Image) error {
	// カウンターをインクリメントする。
	i++
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// カウンターを表示する。
	msg := fmt.Sprintf("%d", i)
	ebitenutil.DebugPrint(screen, msg)
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Increment"); err != nil {
		log.Fatal(err)
	}
}
