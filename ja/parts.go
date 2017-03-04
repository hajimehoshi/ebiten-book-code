package main

import (
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var gophersImage *ebiten.Image

// ebiten.ImagePartsインターフェイスを実装する構造体。
// 4つのパーツに分割して描画する。
type fourParts struct {
	// 描画対象となる画像。
	// この構造体内ではサイズを取得するためだけに使う。
	image *ebiten.Image
}

// 描画するパーツの数を返す関数。
func (f *fourParts) Len() int {
	// 今回は4つのパーツを描画するので、定数4を返す。
	return 4
}

// 描画元矩形の範囲を返す関数。
func (f *fourParts) Src(i int) (int, int, int, int) {
	// Lenが4を返すので、iの範囲は0から3である。
	// それぞれに対応した描画元矩形の範囲を返す。
	// それぞれ左上、右上、左下、右下のパーツに分解している。
	w, h := f.image.Size()
	switch i {
	case 0:
		return 0, 0, w / 2, h / 2
	case 1:
		return w / 2, 0, w, h / 2
	case 2:
		return 0, h / 2, w / 2, h
	case 3:
		return w / 2, h / 2, w, h
	default:
		panic("not reach")	
	}
}

// 描画先矩形の範囲を返す関数。
func (f *fourParts) Dst(i int) (int, int, int, int) {
	// Lenが4を返すので、iの範囲は0から3である。
	w, h := f.image.Size()
	// 少し隙間を開けて表示する。
	const margin = 10
	switch i {
	case 0:
		return 0, 0, w / 2, h / 2
	case 1:
		return w/2 + margin, 0, w + margin, h / 2
	case 2:
		return 0, h/2 + margin, w / 2, h + margin
	case 3:
		return w/2 + margin, h/2 + margin, w + margin, h + margin
	default:
		panic("not reach")	
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// ImagePartsとしてfourPartsを指定する。
	op := &ebiten.DrawImageOptions{}
	op.ImageParts = &fourParts{gophersImage}
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
	if err := ebiten.Run(update, 320, 240, 2, "Parts"); err != nil {
		log.Fatal(err)
	}
}
