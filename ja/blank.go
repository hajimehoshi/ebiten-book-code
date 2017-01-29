package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	// ebiten.Run関数の呼び出しにこの関数が渡される。
	// ebiten.Run関数呼出し後、この関数は定期的に (1秒間に60回) 呼ばれる。
	return nil
}

func main() {
	// ゲームのエントリーポイント。
	// Runは一度呼ばれると基本的には制御が戻ってこない。
	if err := ebiten.Run(update, 320, 240, 2, "Test"); err != nil {
		log.Fatal(err)
	}
}
