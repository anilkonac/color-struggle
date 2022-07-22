package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

const (
	screenWidth  = 576
	screenHeight = 720
	tileLength   = 4
	numRows      = screenHeight / tileLength
	numCol       = screenWidth / tileLength
)

var emptyImage = ebiten.NewImage(1, 1)

type game struct {
	tiles [numRows][numCol]color.RGBA
}

func NewGame() *game {
	game := new(game)

	for iRow := uint16(0); iRow < numRows; iRow++ {
		for iCol := uint16(0); iCol < numCol; iCol++ {
			game.tiles[iRow][iCol] = colornames.Gray
		}
	}

	return game
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

}

func (g *game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	// ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
