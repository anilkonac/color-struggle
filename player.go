package main

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	breadthSpeedMax = 1.5
	breadthSpeedMin = 0.5
)

var (
	timePassedSec float32
	playerShader  *ebiten.Shader
	//go:embed player.kage.go
	shaderBytes []byte
)

func init() {
	var err error
	playerShader, err = ebiten.NewShader(shaderBytes)
	if err != nil {
		panic(err)
	}
}

type player struct {
	color.RGBA
	posX, posY uint8
	drawOpts   ebiten.DrawRectShaderOptions
}

func newPlayer(posX, posY uint8, color color.RGBA) *player {
	player := &player{
		posX: posX,
		posY: posY,
		RGBA: color,
		drawOpts: ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				"TileRadius":  float32(tileLength / 2.0),
				"Time":        float32(timePassedSec),
				"BreathSpeed": float32(breadthSpeedMax),
				"Color":       []float32{float32(color.R), float32(color.G), float32(color.B), float32(color.A)},
			},
		},
	}

	player.drawOpts.GeoM.Translate(float64(posX)*tileLength, float64(posY)*tileLength)

	return player
}

// Returns true if player moves
func (p *player) update() bool {
	timePassedSec += 1.0 / float32(ebiten.MaxTPS())
	p.drawOpts.Uniforms["Time"] = timePassedSec
	// fmt.Printf("timePassedSec: %v\n", timePassedSec)

	pressedUp := inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp)
	pressedDown := inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown)
	pressedLeft := inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft)
	pressedRight := inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight)

	if pressedUp {
		if p.posY != 0 {
			p.posY -= 1
		}
	} else if pressedDown {
		if p.posY != numRows-1 {
			p.posY += 1
		}
	} else if pressedRight {
		if p.posX != numCol-1 {
			p.posX += 1
		}
	} else if pressedLeft {
		if p.posX != 0 {
			p.posX -= 1
		}
	}

	if pressedUp || pressedDown || pressedLeft || pressedRight {
		p.drawOpts.GeoM.Reset()
		p.drawOpts.GeoM.Translate(float64(p.posX)*tileLength, float64(p.posY)*tileLength)
		return true
	}

	return false
}
