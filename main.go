// Copyright 2022 Anıl Konaç

package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

const (
	screenWidth  = 576
	screenHeight = 720
	tileLength   = 24
	numRows      = screenHeight / tileLength
	numCol       = screenWidth / tileLength
)

const (
	playerStartX   = numRows / 2
	playerStartY   = numCol / 2
	targetX        = numCol / 2.0
	targetY        = 0
	colorReduction = 10
)

var (
	gameOver     bool
	gameFinished bool
)

type game struct {
	tiles  [numRows][numCol]tile
	player player
	target target
}

func NewGame() *game {
	game := &game{
		player: *newPlayer(playerStartX, playerStartY, colornames.White),
		target: *newTarget(targetX, targetY),
	}

	for iRow := uint8(0); iRow < numRows; iRow++ {
		for iCol := uint8(0); iCol < numCol; iCol++ {
			game.tiles[iRow][iCol] = *newTile(iRow, iCol, colornames.Black)
		}
	}

	return game
}

func (g *game) restart() {
	*g = game{
		player: *newPlayer(playerStartX, playerStartY, colornames.White),
		target: *newTarget(targetX, targetY),
	}

	for iRow := uint8(0); iRow < numRows; iRow++ {
		for iCol := uint8(0); iCol < numCol; iCol++ {
			g.tiles[iRow][iCol] = *newTile(iRow, iCol, colornames.Black)
		}
	}

	gameOver = false
	gameFinished = false
}

func (g *game) Update() error {
	if gameOver || gameFinished {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.restart()
		}
		return nil
	}

	g.target.update()

	// prevPosX, prevPosY := g.player.posX, g.player.posY
	playerMoved := g.player.update()
	if playerMoved {

		// Paint player's tile
		var reduR, reduG, reduB uint8 = colorReduction, colorReduction, colorReduction
		if g.player.R < colorReduction {
			reduR = g.player.R
		}
		if g.player.G < colorReduction {
			reduG = g.player.G
		}
		if g.player.B < colorReduction {
			reduB = g.player.B
		}
		g.tiles[g.player.posY][g.player.posX].paint(color.RGBA{reduR, reduG, reduB, 1.0})
		// g.tiles[prevPosY][prevPosX].paint(color.RGBA{reduR, reduG, reduB, 1.0})

		if g.player.posX == g.target.posX && g.player.posY == g.target.posY {
			gameFinished = true
		}

		if g.player.R == 0 && g.player.G == 0 && g.player.B == 0 {
			gameOver = true
		}
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
	screen.DrawRectShader(tileLength, tileLength, shaderPlayer, &g.player.drawOpts)

	// Draw target
	screen.DrawRectShader(tileLength, tileLength, shaderTarget, &g.target.drawOpts)

	// Draw GameOver image
	if gameOver {
		screen.DrawImage(imageTextGameOver, &drawOptionsTextGameOver)
		screen.DrawImage(imageTextRestart, &drawOptionsRestart)
	}
	if gameFinished {
		screen.DrawImage(imageTextSuccess, &drawOptionsTextSuccess)
		screen.DrawImage(imageTextRestart, &drawOptionsRestart)

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
