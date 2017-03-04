package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	ebitenImage  *ebiten.Image
	ebitenImageX = 0 // X座標
)

func update(screen *ebiten.Image) error {
	// 描画先のX座標を更新する。
	ebitenImageX++
	if ebiten.IsRunningSlowly() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(ebitenImageX), 0)
	screen.DrawImage(ebitenImage, op)
	return nil
}

func main() {
	var err error
	ebitenImage, _, err = ebitenutil.NewImageFromFile("./ebiten.png",
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, 320, 240, 2, "Image Move"); err != nil {
		log.Fatal(err)
	}
}
