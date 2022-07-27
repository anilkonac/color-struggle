// Copyright 2022 Anıl Konaç

package main

import (
	"image/color"

	"github.com/beefsack/go-astar"
	"github.com/hajimehoshi/ebiten/v2"
)

const paintMultiplier = 10

var whiteImage = ebiten.NewImage(1, 1)

func init() {
	whiteImage.Fill(color.White)
}

type tile struct {
	posRow, posCol uint8
	color          color.RGBA
	drawOpt        ebiten.DrawImageOptions
}

func newTile(posRow, posCol uint8, color color.RGBA) *tile {
	tile := &tile{
		posRow: posRow,
		posCol: posCol,
		color:  color,
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

func (t *tile) up() *tile {
	if t.posRow == 0 {
		// return nil
		return t
	}

	return &tiles[t.posRow-1][t.posCol]
}

func (t *tile) down() *tile {
	if t.posRow == numRows-1 {
		// return nil
		return t
	}

	return &tiles[t.posRow+1][t.posCol]
}

func (t *tile) left() *tile {
	if t.posCol == 0 {
		// return nil
		return t
	}

	return &tiles[t.posRow][t.posCol-1]
}

func (t *tile) right() *tile {
	if t.posCol == numCol-1 {
		// return nil
		return t
	}

	return &tiles[t.posRow][t.posCol+1]
}

func (t *tile) PathNeighbors() []astar.Pather {
	return []astar.Pather{
		t.up(),
		t.right(),
		t.down(),
		t.left(),
	}
}

func (t *tile) PathNeighborCost(to astar.Pather) float64 {
	return 1.0
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*tile)
	absX := float64(toT.posCol) - float64(t.posCol)
	if absX < 0 {
		absX = -absX
	}
	absY := float64(toT.posRow) - float64(t.posRow)
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}
