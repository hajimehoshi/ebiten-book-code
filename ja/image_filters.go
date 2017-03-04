package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	ebitenImageNearest *ebiten.Image
	ebitenImageLinear  *ebiten.Image
)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// 4倍拡大してそれぞれ描画する。
	// 大きめに拡大したのは、フィルタの効果をわかりやすくするため。
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4)
	screen.DrawImage(ebitenImageNearest, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4)
	op.GeoM.Translate(0, 100)
	screen.DrawImage(ebitenImageLinear, op)
	return nil
}

func main() {
	var err error
	// ebiten.Imageオブジェクト作成時にFilterLinearを指定する。
	ebitenImageNearest, _, err = ebitenutil.NewImageFromFile("./ebiten.png",
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	ebitenImageLinear, _, err = ebitenutil.NewImageFromFile("./ebiten.png",
		ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, 320, 240, 2, "Image Filters"); err != nil {
		log.Fatal(err)
	}
}
