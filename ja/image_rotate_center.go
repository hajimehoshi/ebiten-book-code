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
	// 画像の大きさを得る。
	w, h := ebitenImage.Size()
	// ebitenImageを時計回りの向きに、画像の中心を回転中心として
	// 45度回転させる。
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(math.Pi/4)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
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
