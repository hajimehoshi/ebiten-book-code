package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	// 現在のマウスカーソル座標を取得する。
	x, y := ebiten.CursorPosition()
	msg := fmt.Sprintf("(%d, %d)\n", x, y)
	// マウス左ボタンが押されていたら、メッセージを追加する。
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		msg += "Left button is pressed"
	}
	if ebiten.IsRunningSlowly() {
		return nil
	}
	if err := ebitenutil.DebugPrint(screen, msg); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Mouse"); err != nil {
		log.Fatal(err)
	}
}
