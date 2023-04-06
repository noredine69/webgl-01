package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth           = 1024
	screenHeight          = 800
	maxAngle              = 256
	turningLeftRectIndex  = 0
	centerRectIndex       = 1
	turningRightRectIndex = 2
	backwardRectIndex     = 0
	forwardSpeedRectIndex = 1
	forwardRectIndex      = 2
)

var (
	ebitenImage *ebiten.Image
)

type Game struct {
	keys           []ebiten.Key
	op             ebiten.DrawImageOptions
	spaceShipsRect [][]image.Rectangle
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "")
	//fmt.Println("plip ", len(g.keys))

	//w, h := ebitenImage.Bounds().Dx(), ebitenImage.Bounds().Dy()
	w, h := float64(1.5), float64(1.5)
	txTranslation := float64(0)
	tyTranslation := float64(0)
	translation := false
	speedIndexRec := forwardRectIndex
	slidingIndexRec := centerRectIndex
	for _, k := range g.keys {
		if k == ebiten.KeyArrowLeft {
			slidingIndexRec = turningLeftRectIndex
			txTranslation = -w
			translation = true
		} else if k == ebiten.KeyArrowRight {
			slidingIndexRec = turningRightRectIndex
			txTranslation = w
			translation = true
		}
		if k == ebiten.KeyArrowDown {
			speedIndexRec = backwardRectIndex
			tyTranslation = h
			translation = true
		} else if k == ebiten.KeyArrowUp {
			speedIndexRec = forwardSpeedRectIndex
			tyTranslation = -h
			translation = true
		}
	}
	if translation {
		g.op.GeoM.Translate(txTranslation, tyTranslation)
	}
	screen.DrawImage(ebitenImage.SubImage(g.spaceShipsRect[speedIndexRec][slidingIndexRec]).(*ebiten.Image), &g.op)

}
func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func init() {
	fmt.Println("init")
	// Decode an image from the image file's byte slice.
	//spaceshipsprites
	//file, errOpen := os.Open("/home/ubuntu/Dev/supervision-apps/webgl-01/assets/spaceshipsprites.png")
	file, errOpen := os.Open("assets/spaceshipsprites.png")
	if errOpen != nil {
		fmt.Println("err ", errOpen)
	}

	//nolint: staticcheck, errcheck
	defer file.Close()

	fileReader := bufio.NewReader(file)

	//img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	img, _, err := image.Decode(fileReader)
	if err != nil {
		log.Fatal(err)
	}
	origEbitenImage := ebiten.NewImageFromImage(img)

	s := origEbitenImage.Bounds().Size()
	ebitenImage = ebiten.NewImage(s.X, s.Y)

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(1)
	ebitenImage.DrawImage(origEbitenImage, op)
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Xenon 3")
	if err := ebiten.RunGame(&Game{
		spaceShipsRect: [][]image.Rectangle{
			{
				image.Rect(0, 0, 32, 36),
				image.Rect(39, 0, 78, 36),
				image.Rect(86, 0, 116, 36),
			},
			{
				image.Rect(0, 40, 32, 82),
				image.Rect(39, 40, 78, 82),
				image.Rect(86, 40, 116, 82),
			},
			{
				image.Rect(0, 86, 32, 126),
				image.Rect(39, 86, 78, 126),
				image.Rect(86, 86, 116, 126),
			},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
