// Copyright 2022 Anıl Konaç

package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	minDistPlayerTarget = numRows
	colorReduction      = 10
	numSources          = 4
)

var (
	gameOver     bool
	gameFinished bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type game struct {
	tiles   [numRows][numCol]tile
	player  player
	target  target
	sources [numSources]source
}

func NewGame() *game {
	var playerX, playerY, targetX, targetY int
	var tooNear bool = true
	for tooNear {
		playerX = rand.Intn(numCol)
		playerY = rand.Intn(numRows)
		targetX = rand.Intn(numCol)
		targetY = rand.Intn(numRows)
		tooNear = ((playerX-targetX)*(playerX-targetX) + (playerY-targetY)*(playerY-targetY)) < minDistPlayerTarget*minDistPlayerTarget

	}

	game := &game{
		player: *newPlayer(playerX, playerY, colornames.White),
		target: *newTarget(targetX, targetY),
	}

	// Create tiles
	for iRow := uint8(0); iRow < numRows; iRow++ {
		for iCol := uint8(0); iCol < numCol; iCol++ {
			game.tiles[iRow][iCol] = *newTile(iRow, iCol, colornames.Black)
		}
	}

	createColorSources(playerX, playerY, &game.sources)

	return game
}

func (g *game) restart() {

	var playerX, playerY, targetX, targetY int
	var tooNear bool = true
	for tooNear {
		playerX = rand.Intn(numCol)
		playerY = rand.Intn(numRows)
		targetX = rand.Intn(numCol)
		targetY = rand.Intn(numRows)
		tooNear = ((playerX-targetX)*(playerX-targetX) + (playerY-targetY)*(playerY-targetY)) < minDistPlayerTarget*minDistPlayerTarget

	}

	*g = game{
		player: *newPlayer(playerX, playerY, colornames.White),
		target: *newTarget(targetX, targetY),
	}

	for iRow := uint8(0); iRow < numRows; iRow++ {
		for iCol := uint8(0); iCol < numCol; iCol++ {
			g.tiles[iRow][iCol] = *newTile(iRow, iCol, colornames.Black)
		}
	}

	createColorSources(playerX, playerY, &g.sources)

	gameOver = false
	gameFinished = false
}

func createColorSources(playerX, playerY int, sources *[numSources]source) {
	// Create color sources
	for iSource := 0; iSource < numSources; iSource++ {
		sourceColorIndex := rand.Intn(3)
		var sourceColor color.RGBA
		switch sourceColorIndex {
		case 0:
			sourceColor = colornames.Red
		case 1:
			sourceColor = colornames.Green
		case 2:
			sourceColor = colornames.Blue
		}
		sourceX, sourceY := playerX, playerY
		for sourceX == playerX && sourceY == playerY {
			sourceX = rand.Intn(numRows)
			sourceY = rand.Intn(numCol)
			sources[iSource] = *newSource(sourceX, sourceY, sourceColor)
		}
	}
}

func (g *game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.restart()
		return nil
	}

	if gameOver || gameFinished {
		return nil
	}

	g.target.update()

	prevPosX, prevPosY := g.player.posX, g.player.posY
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
		// g.tiles[g.player.posY][g.player.posX].paint(color.RGBA{reduR, reduG, reduB, 1.0})
		g.tiles[prevPosY][prevPosX].paint(color.RGBA{reduR, reduG, reduB, 1.0})

		// Gather color source
		for iSource := 0; iSource < numSources; iSource++ {
			source := &g.sources[iSource]
			if g.player.posX == source.posX && g.player.posY == source.posY {
				source.eaten = true
				g.player.gatherColor(source.RGBA)
			}
		}

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
	screen.Fill(colornames.Darkgray)

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

	// Draw color sources
	for iSource := 0; iSource < numSources; iSource++ {
		source := &g.sources[iSource]
		if !source.eaten {
			screen.DrawRectShader(tileLength, tileLength, shaderSource, &source.drawOpt)
		}

	}

	// Draw GameOver image
	if gameOver {
		screen.DrawImage(imageTextGameOver, &drawOptionsTextGameOver)
		screen.DrawImage(imageTextRestart, &drawOptionsRestart)
	}
	if gameFinished {
		screen.DrawImage(imageTextSuccess, &drawOptionsTextSuccess)
		screen.DrawImage(imageTextRestart, &drawOptionsRestart)

	}

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f  FPS: %.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowTitle("Color Struggle")
	// ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	// ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	// ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMinimum)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
