package main

import (
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	tick         = 0
	gophersImage *ebiten.Image
)

func update(screen *ebiten.Image) error {
	// tickは1フレームごとに1増える値である。
	// 120フレーム (2秒) で0に戻す。
	tick++
	tick %= 120
	if ebiten.IsRunningSlowly() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	// 1秒ごとにセピア色にする。
	if 60 <= tick {
		op.ColorM.ChangeHSV(0, 5.0/15.0, 1)
		op.ColorM.Translate(2.0/15.0, -2.0/15.0, -4.0/15.0, 0)
	}
	screen.DrawImage(gophersImage, op)
	return nil
}

func main() {
	var err error
	gophersImage, _, err = ebitenutil.NewImageFromFile("./gophers_photo.jpg",
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, 320, 240, 2, "Sepia"); err != nil {
		log.Fatal(err)
	}
}
