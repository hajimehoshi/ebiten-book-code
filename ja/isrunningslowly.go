package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	// 状態更新処理
	if ebiten.IsRunningSlowly() {
		// IsRunningSlowlyがtrueのときは、screenに対する
		// 描画結果は実際の画面に反映されない。
		// よって描画処理を行わず終了してよい。
		return nil
	}
	// 描画処理
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Test"); err != nil {
		log.Fatal(err)
	}
}
