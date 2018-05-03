package main

import (
	"image"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

var player_anim0 = `------+++++-----
----+++++++++---
---+++++++++++--
--+++++++++++++-
--++++--+++--++-
-+++++--+++--++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++-+++++-+++
+++++++-----++++
+-++++++++++++++
--++++++++++++-+
---++++++++++---
---++-----++----
--+++++--+++++--`

var player_anim1 = `------+++++-----
----+++++++++---
---+++++++++++--
--+++++++++++++-
--++++--+++--++-
-+++++--+++--++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++-+++++-+++
+++++++-----++++
+-++++++++++++++
--++++++++++++-+
---++++++++++---
--+++++---++----
---------+++++--`

var player_anim2 = `------+++++-----
----+++++++++---
---+++++++++++--
--+++++++++++++-
--++++--+++--++-
-+++++--+++--++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++-+++++-+++
+++++++-----++++
+-++++++++++++++
--++++++++++++-+
---++++++++++---
---++----+++++--
--+++++---------`

var (
	charWidth   = 16
	charHeight  = 16
	tmpImage    *image.RGBA
	playerAnim0 *ebiten.Image
	playerAnim1 *ebiten.Image
	playerAnim2 *ebiten.Image
)

func createImageFromString(charString string, img *image.RGBA) {
	for indexY, line := range strings.Split(charString, "\n") {
		for indexX, str := range line {
			pos := 4*indexY*charWidth + 4*indexX
			if string(str) == "+" {
				img.Pix[pos] = 0xff
				img.Pix[pos+1] = 0xff
				img.Pix[pos+2] = 0xff
				img.Pix[pos+3] = 0xff
			} else {
				img.Pix[pos] = 0
				img.Pix[pos+1] = 0
				img.Pix[pos+2] = 0
				img.Pix[pos+3] = 0
			}
		}
	}
}

type position struct {
	X int
	Y int
}

type Sprite struct {
	//	アニメーションさせる画像配列
	Images []*ebiten.Image
	//	総イメージ数
	ImageNum int
	//	現在何枚目の画像が描画されているか
	CurrentNum int
	//	現在描画位置
	Position position
	//	フレーム数カウンター
	count int
}

func NewSprite(images []*ebiten.Image) *Sprite {
	return &Sprite{
		Images:   images,
		ImageNum: len(images),
	}
}

//	currentImage 描画する画像選択し返す
func (s *Sprite) currentImage() *ebiten.Image {
	//	5フレームごとの画像更新
	if s.count > 5 {
		s.count = 0
		s.CurrentNum++
		s.CurrentNum %= s.ImageNum
	}
	return s.Images[s.CurrentNum]
}

func (s *Sprite) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.Position.X -= 1
		s.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.Position.X += 1
		s.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.Position.Y -= 1
		s.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.Position.Y += 1
		s.count++
	}
}

func (s *Sprite) DrawImage(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Position.X), float64(s.Position.Y))
	screen.DrawImage(s.currentImage(), op)
}

type Game struct {
	Char *Sprite
}

func (g *Game) Init() {
	tmpImage = image.NewRGBA(image.Rect(0, 0, charWidth, charHeight))

	createImageFromString(player_anim0, tmpImage)
	playerAnim0, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	playerAnim0.ReplacePixels(tmpImage.Pix)

	createImageFromString(player_anim1, tmpImage)
	playerAnim1, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	playerAnim1.ReplacePixels(tmpImage.Pix)

	createImageFromString(player_anim2, tmpImage)
	playerAnim2, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	playerAnim2.ReplacePixels(tmpImage.Pix)

	//	生成した画像からスプライトを作成
	images := []*ebiten.Image{
		playerAnim0,
		playerAnim1,
		playerAnim2,
	}
	g.Char = NewSprite(images)
	g.Char.Position.X = 10
	g.Char.Position.Y = 10
}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	g.Char.Move()

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Char.DrawImage(screen)
	return nil
}

func main() {
	game := Game{}
	game.Init()

	if err := ebiten.Run(game.MainLoop, 320, 240, 2, "go-pixelman3"); err != nil {
		log.Fatal(err)
	}
}
