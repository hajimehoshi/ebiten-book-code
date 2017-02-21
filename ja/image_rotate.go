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
	// ebitenImageを45度時計回りの向きに(原点中心に)回転させる。
	// 45度はラジアンでπ/4である。
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(math.Pi / 4)
	if err := screen.DrawImage(ebitenImage, op); err != nil {
		return err
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
	if err := ebiten.Run(update, 320, 240, 2, "Image Rotating"); err != nil {
		log.Fatal(err)
	}
}
