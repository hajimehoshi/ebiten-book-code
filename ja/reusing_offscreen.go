package main

import (
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	ebitenImage    *ebiten.Image
	offscreenImage *ebiten.Image
	angle          int
)

func update(screen *ebiten.Image) error {
	angle++
	angle %= 360
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// オフスクリーンバッファを初期化する。
	// 前節で見た通り、オフスクリーンは自前で初期化する必要がある。
	if err := offscreenImage.Clear(); err != nil {
		return err
	}
	// オフスクリーンバッファに対し、1個海老天の画像を描画する。
	op := &ebiten.DrawImageOptions{}
	if err := offscreenImage.DrawImage(ebitenImage, op); err != nil {
		return err
	}
	// オフスクリーンバッファを画面に描画する。
	op = &ebiten.DrawImageOptions{}
	if err := screen.DrawImage(offscreenImage, op); err != nil {
		return err
	}

	// オフスクリーンバッファをまた消去する。
	// 直前のDrawImageについては、その時点でのオフスクリーンバッファの
	// 内容が使われるので、影響はない。
	if err := offscreenImage.Clear(); err != nil {
		return err
	}
	// 海老天を回転して描画する。
	op = &ebiten.DrawImageOptions{}
	w, h := ebitenImage.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(angle) * math.Pi / 180)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
	op.GeoM.Translate(float64(h)/2, 0)
	if err := offscreenImage.DrawImage(ebitenImage, op); err != nil {
		return err
	}
	// オフスクリーンバッファを画面に別の場所に描画する。
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 0)
	if err := screen.DrawImage(offscreenImage, op); err != nil {
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
	offscreenImage, err = ebiten.NewImage(100, 100, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, 320, 240, 2, "Reusing Offscreen"); err != nil {
		log.Fatal(err)
	}
}
