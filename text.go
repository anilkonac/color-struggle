package main

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	dpi              = 72
	fontSizeGameOver = 96
	fontSizeRestart  = 32
	textGameOver     = "Game Over!"
	textGameSuccess  = "You Win!"
	textGameRestart  = "Press R to restart"
)

var (
	//go:embed Moonlight-Regular.ttf
	bytesFontMoonlight      []byte
	imageTextGameOver       *ebiten.Image
	imageTextSuccess        *ebiten.Image
	imageTextRestart        *ebiten.Image
	drawOptionsTextGameOver ebiten.DrawImageOptions
	drawOptionsTextSuccess  ebiten.DrawImageOptions
	drawOptionsRestart      ebiten.DrawImageOptions
)

func init() {
	tt, err := opentype.Parse(bytesFontMoonlight)
	panicErr(err)

	fontFaceGameOver, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSizeGameOver,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	panicErr(err)

	fontFaceRestart, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSizeRestart,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	panicErr(err)

	// Prepare Text Images
	imageTextGameOver, drawOptionsTextGameOver = prepareTextImage(fontFaceGameOver, textGameOver, color.White, 0, 0)
	imageTextSuccess, drawOptionsTextSuccess = prepareTextImage(fontFaceGameOver, textGameSuccess, color.White, 0, 0)
	imageTextRestart, drawOptionsRestart = prepareTextImage(fontFaceRestart, textGameRestart, color.White, 0, 64)
}

func prepareTextImage(fontFace font.Face, txt string, color color.Color, shiftX, shiftY int) (image *ebiten.Image, opt ebiten.DrawImageOptions) {
	boundText := text.BoundString(fontFace, txt)
	boundTextSize := boundText.Size()
	image = ebiten.NewImage(boundTextSize.X, boundTextSize.Y)
	text.Draw(image, txt, fontFace, -boundText.Min.X, -boundText.Min.Y, color)
	opt.GeoM.Translate(
		float64((screenWidth-boundTextSize.X)/2.0-boundText.Min.X+shiftX),
		float64((screenHeight-boundTextSize.Y)/2.0 /*-boundText.Min.Y*/ +shiftY))

	return
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
