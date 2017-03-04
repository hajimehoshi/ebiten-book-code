package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// 画面 (screen) を赤色で塗りつぶす。
	// color.RGBAは (premultiplied-alphaな) 8bitカラーを表す。
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Fill"); err != nil {
		log.Fatal(err)
	}
}
