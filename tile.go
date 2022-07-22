// Copyright 2022 Anıl Konaç

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var colorImageMap map[color.RGBA]*ebiten.Image

func init() {
	colorImageMap = make(map[color.RGBA]*ebiten.Image)
}

type tile struct {
	posRow, posCol uint8
	color          color.RGBA
	image          *ebiten.Image
}

func newTile(posRow, posCol uint8, color color.RGBA) *tile {
	tile := &tile{
		posRow: posRow,
		posCol: posCol,
		color:  color,
	}

	// Assign/create color image
	if img, ok := colorImageMap[color]; ok {
		tile.image = img
	} else {
		// image := ebiten.NewImage(1, 1)
		image := ebiten.NewImage(tileLength-1, tileLength-1)
		image.Fill(color)
		colorImageMap[color] = image
		tile.image = image
	}

	return tile

}
