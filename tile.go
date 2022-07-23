// Copyright 2022 Anıl Konaç

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var whiteImage = ebiten.NewImage(1, 1)

func init() {
	whiteImage.Fill(color.White)
}

type tile struct {
	// posRow, posCol uint8
	// color          color.RGBA
	drawOpt ebiten.DrawImageOptions
}

func newTile(posRow, posCol uint8, color color.RGBA) *tile {
	tile := &tile{
		// posRow: posRow,
		// posCol: posCol,
		// color:  color,
	}

	// Set geometry matrix
	tile.drawOpt.GeoM.Scale(tileLength-1, tileLength-1)
	tile.drawOpt.GeoM.Translate(float64(posCol)*tileLength, float64(posRow)*tileLength)

	// Set color matrix
	tile.drawOpt.ColorM.ScaleWithColor(color)

	return tile
}

func (t tile) draw(dst *ebiten.Image) {
	dst.DrawImage(whiteImage, &t.drawOpt)
}
