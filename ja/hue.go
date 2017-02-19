package main

import (
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var ebitenImage *ebiten.Image

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// 色々色相を変えて画像を描画する。
	for i, a := range []int{0, 60, 120, 180, 240, 300} {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i) * 50, 0)
		op.ColorM.ChangeHSV(float64(a) * math.Pi / 180, 1, 1)
		if err := screen.DrawImage(ebitenImage, op); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var err error
	ebitenImage, _, err = ebitenutil.NewImageFromFile("./ebiten.png",
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, 320, 240, 2, "Hue"); err != nil {
		log.Fatal(err)
	}
}
