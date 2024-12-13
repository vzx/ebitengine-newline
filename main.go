package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var mplusFaceSource *text.GoTextFaceSource

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := "The quick\n\tbrown fox jumps\n\n\t\tover the lazy\r\n\rdog."
	op := &text.DrawOptions{}
	op.GeoM.Translate(10, 40)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   32,
	}, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	var err error
	if mplusFaceSource, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf)); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(1280, 720)
	if err = ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
