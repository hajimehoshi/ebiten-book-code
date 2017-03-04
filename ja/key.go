package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	// 入力は描画ではなく論理更新の範疇なので、
	// IsRunningSlowlyを呼ぶ前にIsKeyPressedを呼ぶ。

	// Aキーが押されたかどうかの判別を行う。
	if !ebiten.IsKeyPressed(ebiten.KeyA) {
		return nil
	}
	if ebiten.IsRunningSlowly() {
		return nil
	}
	ebitenutil.DebugPrint(screen, "A key is pressed!")
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Key"); err != nil {
		log.Fatal(err)
	}
}
