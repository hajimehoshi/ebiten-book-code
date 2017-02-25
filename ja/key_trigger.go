package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Aキーの状態を表す整数値。
// 具体的には、どのくらい長い間押されたかを表す値。
var keyAState = 0

func update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		keyAState++
	} else {
		keyAState = 0
	}
	if ebiten.IsRunningSlowly() {
		return nil
	}
	if keyAState != 1 {
		return nil
	}
	// keyAStateが1のときにのみ文字を描画する。
	if err := ebitenutil.DebugPrint(screen, "A key is triggered!"); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Key Trigger"); err != nil {
		log.Fatal(err)
	}
}
