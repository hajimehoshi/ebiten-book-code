package main

import (
	_ "image/jpeg"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var ebitenImage *ebiten.Image

// 座標を表す構造体。
type position struct {
	x int
	y int
}

// ebiten.ImagePartsインターフェイスを実装する。
type sprites struct {
	image     *ebiten.Image
	positions []position
}

func (s *sprites) Len() int {
	return len(positions)
}

func (s *sprites) Src(i int) (int, int, int, int) {
	// 描画元矩形は画像そのまま使う。
	w, h := s.image.Size()
	return 0, 0, w, h
}

func (s *sprites) Dst(i int) (int, int, int, int) {
	// X座標とY座標をiに合わせて変更する。
	w, h := s.image.Size()
	p := s.positions[i]
	return p.x, p.y, p.x + w, p.y + h
}

var positions []position

func init() {
	// 乱数シードを設定する。
	rand.Seed(time.Now().UnixNano())
	// 各スプライトの座標を乱数で決定する。
	positions = make([]position, 100)
	for i := range positions {
		positions[i].x = rand.Intn(320)
		positions[i].y = rand.Intn(240)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.ImageParts = &sprites{
		image:     ebitenImage,
		positions: positions,
	}
	screen.DrawImage(ebitenImage, op)
	return nil
}

func main() {
	var err error
	ebitenImage, _, err = ebitenutil.NewImageFromFile("./ebiten.png",
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, 320, 240, 2, "Sprites"); err != nil {
		log.Fatal(err)
	}
}
