// Copyright 2022 Anıl Konaç

package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

const (
	screenWidth  = 576
	screenHeight = 720
	tileLength   = 24
	numRows      = screenHeight / tileLength
	numCol       = screenWidth / tileLength
)

type game struct {
	tiles [numRows][numCol]tile
}

func NewGame() *game {
	game := new(game)

	for iRow := uint8(0); iRow < numRows; iRow++ {
		for iCol := uint8(0); iCol < numCol; iCol++ {
			game.tiles[iRow][iCol] = *newTile(iRow, iCol, colornames.Gray)
		}
	}

	return game
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	for iRow := uint8(0); iRow < numRows; iRow++ {
		for iCol := uint8(0); iCol < numCol; iCol++ {
			curTile := &g.tiles[iRow][iCol]
			screen.DrawImage(emptyImage, &curTile.drawOpt)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f  FPS: %.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))

}

func (g *game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowTitle("Color Struggle")
	// ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	// ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMinimum)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
