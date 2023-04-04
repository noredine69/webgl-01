package main

import (
	"bytes"
	"fmt"
	"image"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 800
	screenHeight = 600
	maxAngle     = 256
)

var (
	ebitenImage *ebiten.Image
)

type Game struct {
	keys []ebiten.Key
	op   ebiten.DrawImageOptions
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	//fmt.Println("plip ", len(g.keys))

	w, h := ebitenImage.Bounds().Dx(), ebitenImage.Bounds().Dy()
	txTranslation := float64(0)
	tyTranslation := float64(0)
	translation := false
	for _, k := range g.keys {
		if k == ebiten.KeyArrowLeft {
			txTranslation = -float64(w) / 2
			translation = true
		} else if k == ebiten.KeyArrowRight {
			txTranslation = float64(w) / 2
			translation = true
		}
		if k == ebiten.KeyArrowDown {
			tyTranslation = float64(h) / 2
			translation = true
		} else if k == ebiten.KeyArrowUp {
			tyTranslation = -float64(h) / 2
			translation = true
		}
	}
	if translation {
		fmt.Println("plip ", txTranslation, tyTranslation, w, h)
		//g.op.GeoM.Reset()
		g.op.GeoM.Translate(txTranslation, tyTranslation)
	}

	screen.DrawImage(ebitenImage, &g.op)

}
func (g *Game) Update() error {

	//fmt.Println("plip")
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	//fmt.Println("plip ", len(g.keys))

	return nil

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func init() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}
	origEbitenImage := ebiten.NewImageFromImage(img)

	s := origEbitenImage.Bounds().Size()
	ebitenImage = ebiten.NewImage(s.X, s.Y)

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(0.5)
	ebitenImage.DrawImage(origEbitenImage, op)
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
