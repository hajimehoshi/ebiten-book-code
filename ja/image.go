package main

import (
	// PNGファイルを扱うため、デコーダを有効にする。
	// この、デコーダ登録のためだけにパッケージをインポートする手法については
	// 標準ライブラリのimageパッケージを参照すること:
	// https://golang.org/pkg/image/
	// TODO: Bug fix in ebitenutil
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var ebitenImage *ebiten.Image

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// ebitenImageを画面に描画する。
	// 第2引数であるDrawImageOptionsは今回は指定を省略する (nilを指定する)。
	screen.DrawImage(ebitenImage, nil)
	return nil
}

func main() {
	var err error
	// ebitenutilパッケージのNewImageFromFile関数を使って、
	// ファイル名から直接*ebiten.Imageを作成する。
	// 2つめの戻り値は*image.Imageだが、今回は使わないので無視する。
	// フィルタとして今回はFilterNearestを使用した。
	// この値は画像の拡大縮小の際のフィルタを表すが、実際に画像を拡大縮小したり
	// しない限りは使用されないので、今回はとりあえずなんでも良い。
	ebitenImage, _, err = ebitenutil.NewImageFromFile("./ebiten.png",
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, 320, 240, 2, "Image"); err != nil {
		log.Fatal(err)
	}
}
