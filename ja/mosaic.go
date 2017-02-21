package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	// Gopherの画像。
	gopherImage *ebiten.Image
	// 一旦縮小したGopherを描画するためのオフスクリーンバッファ。
	tmpImage *ebiten.Image
)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	const scale = 4.0
	// Gopherをオフスクリーンバッファに縮小して描画する。
	if err := tmpImage.Clear(); err != nil {
		return err
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1/scale, 1/scale)
	if err := tmpImage.DrawImage(gopherImage, op); err != nil {
		return err
	}
	// Gopherを画面に描画する。
	op = &ebiten.DrawImageOptions{}
	if err := screen.DrawImage(gopherImage, op); err != nil {
		return err
	}
	// オフスクリーンバッファ画面に違う位置に描画する。
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(100, 0)
	if err := screen.DrawImage(tmpImage, op); err != nil {
		return err
	}
	return nil
}

func main() {
	var err error
	gopherImage, _, err = ebitenutil.NewImageFromFile("./gopher.png",
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	// オフスクリーンバッファを生成する。
	// 大きさはgopherImageの大きさ以上であればなんでも良い。
	// とりあえず画面と同じ大きさにした。
	// モザイクのためには、フィルタはNearestフィルタである必要がある。
	tmpImage, err = ebiten.NewImage(320, 240, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, 320, 240, 2, "Mosaic"); err != nil {
		log.Fatal(err)
	}
}
