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
	screenWidth    = 576
	screenHeight   = 720
	tileLength     = 24
	numRows        = screenHeight / tileLength
	numCol         = screenWidth / tileLength
	colorReduction = 50
)

type game struct {
	tiles  [numRows][numCol]tile
	player player
}

func NewGame() *game {
	game := &game{
		player: *newPlayer(numRows/2, numCol/2, colornames.White),
	}

	for iRow := uint8(0); iRow < numRows; iRow++ {
		for iCol := uint8(0); iCol < numCol; iCol++ {
			game.tiles[iRow][iCol] = *newTile(iRow, iCol, colornames.Black)
		}
	}

	return game
}

func (g *game) Update() error {
	prevPosX, prevPosY := g.player.posX, g.player.posY
	playerMoved := g.player.update()
	if playerMoved {
		// Paint player's tile
		// g.tiles[g.player.posY][g.player.posX].paint(color.RGBA{colorReduction, colorReduction, colorReduction, 1.0})
		g.tiles[prevPosY][prevPosX].paint(color.RGBA{colorReduction, colorReduction, colorReduction, 1.0})
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Gray)

	// Draw tiles
	for iRow := uint8(0); iRow < numRows; iRow++ {
		for iCol := uint8(0); iCol < numCol; iCol++ {
			g.tiles[iRow][iCol].draw(screen)
		}
	}

	// Draw player
	screen.DrawRectShader(tileLength, tileLength, playerShader, &g.player.drawOpts)

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
