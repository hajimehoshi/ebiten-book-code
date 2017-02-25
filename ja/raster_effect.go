package main

import (
	_ "image/jpeg"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var gophersImage *ebiten.Image

type rasterEffectParts struct {
	image *ebiten.Image
	angle int
}

func (r *rasterEffectParts) Update() {
	// 角度を更新する。
	// この更新方法で角速度が決まる。
	r.angle++
	r.angle %= 360
}

func (r *rasterEffectParts) Len() int {
	// 画像を横に細長い形に分割する。
	// パーツの高さは1ピクセルであり、数も高さのピクセル数と同じである。
	_, h := r.image.Size()
	return h
}

func (r *rasterEffectParts) Src(i int) (int, int, int, int) {
	// 描画元矩形をiに合わせて指定する。
	// 高さは1ピクセルであることに注意。
	w, _ := r.image.Size()
	return 0, i, w, i + 1
}

func (r *rasterEffectParts) Dst(i int) (int, int, int, int) {
	// 波形の振幅。
	const amplitude = 8
	w, _ := r.image.Size()
	// 波形の位相。
	// 位相の計算にiを使うことで、角行ごとに微妙に異なる
	// ズレを生じさせる。
	a := r.angle + 4*i
	// ズレdを計算する。
	d := int(math.Floor(math.Sin(float64(a)*math.Pi/180) * amplitude))
	return d, i, w + d, i + 1
}

var parts *rasterEffectParts

func update(screen *ebiten.Image) error {
	parts.Update()
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// ImagePartsとしてrasterEffectPartsを指定する。
	op := &ebiten.DrawImageOptions{}
	op.ImageParts = parts
	if err := screen.DrawImage(gophersImage, op); err != nil {
		return err
	}
	return nil
}

func main() {
	var err error
	gophersImage, _, err = ebitenutil.NewImageFromFile("./gophers_photo.jpg",
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	parts = &rasterEffectParts{
		image: gophersImage,
	}
	if err := ebiten.Run(update, 320, 240, 2, "Raster Effect"); err != nil {
		log.Fatal(err)
	}
}
