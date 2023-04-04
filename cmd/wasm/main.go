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
	screenWidth  = 320
	screenHeight = 240
	maxAngle     = 256
)

var (
	ebitenImage *ebiten.Image
)

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	fmt.Println("plip ", len(g.keys))
	var keyStrs []string
	var keyNames []string
	for _, k := range g.keys {
		keyStrs = append(keyStrs, k.String())

		fmt.Println("touch ", k.String())
		ebitenutil.DebugPrint(screen, k.String())
		if name := ebiten.KeyName(k); name != "" {
			keyNames = append(keyNames, name)
			ebitenutil.DebugPrint(screen, name)
		}
	}

	//screen.DrawImage(ebitenImage, nil)

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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
