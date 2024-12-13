package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	textv1 "github.com/hajimehoshi/ebiten/v2/text"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var mplusFace font.Face
var mplusFaceSource *textv2.GoTextFaceSource

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	textv1.Draw(screen, "Hello, \nv1 world", mplusFace, 10, 40, color.White)

	msg := "Hello,\nv2 world"
	op := &textv2.DrawOptions{}
	op.GeoM.Translate(500, 40)
	op.ColorScale.ScaleWithColor(color.White)
	textv2.Draw(screen, msg, &textv2.GoTextFace{Source: mplusFaceSource, Size: 32}, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	var err error

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	fatalIfErr(err)

	mplusFace, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	fatalIfErr(err)

	mplusFaceSource, err = textv2.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	fatalIfErr(err)

	ebiten.SetWindowSize(1280, 720)
	if err = ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
