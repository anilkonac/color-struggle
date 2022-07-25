// Copyright 2022 Anıl Konaç

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const paintMultiplier = 10

var whiteImage = ebiten.NewImage(1, 1)

func init() {
	whiteImage.Fill(color.White)
}

type tile struct {
	// posRow, posCol uint8
	color   color.RGBA
	drawOpt ebiten.DrawImageOptions
}

func newTile(posRow, posCol uint8, color color.RGBA) *tile {
	tile := &tile{
		// posRow: posRow,
		// posCol: posCol,
		color: color,
	}

	// Set geometry matrix
	tile.drawOpt.GeoM.Scale(tileLength-1, tileLength-1)
	tile.drawOpt.GeoM.Translate(float64(posCol)*tileLength, float64(posRow)*tileLength)

	// Set color matrix
	tile.drawOpt.ColorM.ScaleWithColor(color)

	return tile
}

func (t *tile) draw(dst *ebiten.Image) {
	dst.DrawImage(whiteImage, &t.drawOpt)
}

func (t *tile) paint(clr color.RGBA) {
	var paintR uint8 = clr.R * paintMultiplier
	var paintG uint8 = clr.G * paintMultiplier
	var paintB uint8 = clr.B * paintMultiplier
	if (255 - t.color.R) >= paintR {
		t.color.R += paintR
	} else {
		t.color.R = 255
	}
	if (255 - t.color.G) >= paintG {
		t.color.G += paintG
	} else {
		t.color.G = 255
	}
	if (255 - t.color.B) >= paintB {
		t.color.B += paintB
	} else {
		t.color.B = 255
	}
	// t.color.A += clr.A
	t.drawOpt.ColorM.Reset()
	t.drawOpt.ColorM.ScaleWithColor(t.color)
}
