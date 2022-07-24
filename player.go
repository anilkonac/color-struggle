package main

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
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

func (p *player) update() {
	timePassedSec += 1.0 / float32(ebiten.MaxTPS())
	p.drawOpts.Uniforms["Time"] = timePassedSec
	// fmt.Printf("timePassedSec: %v\n", timePassedSec)
}
