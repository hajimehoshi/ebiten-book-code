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
	// オフスクリーンバッファをクリアする。
	// screenと違い自動的にクリアされないので、明示的に呼ぶ必要がある。
	offscreenImage.Clear()
	// オフスクリーンバッファに対し、2個海老天の画像を描画する。
	// うち1個は回転させる。
	op := &ebiten.DrawImageOptions{}
	offscreenImage.DrawImage(ebitenImage, op)
	op = &ebiten.DrawImageOptions{}
	w, h := ebitenImage.Size()
	// 画像の中心を回転中心としてangle度分回転する。
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(angle) * math.Pi / 180)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
	op.GeoM.Translate(float64(h)/2, 0)
	offscreenImage.DrawImage(ebitenImage, op)
	// オフスクリーンバッファを画面に描画する。
	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(offscreenImage, op)
	// 同じオフスクリーンバッファを拡大して描画する。
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(100, 100)
	screen.DrawImage(offscreenImage, op)
	return nil
}

func main() {
	var err error
	ebitenImage, _, err = ebitenutil.NewImageFromFile("./ebiten.png",
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	// オフスクリーンレンダリング用のバッファを作成する。
	// 今回はmain関数内で生成した。
	// なおebiten.Imageを生成する処理は重い処理であり、
	// update関数内で毎フレーム生成するのは推奨されない。
	offscreenImage, err = ebiten.NewImage(100, 100, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, 320, 240, 2, "Offscreen"); err != nil {
		log.Fatal(err)
	}
}
