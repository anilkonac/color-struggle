package main

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed target.kage.go
	bytesShaderTarget []byte
	shaderTarget      *ebiten.Shader
)

func init() {
	var err error
	shaderTarget, err = ebiten.NewShader(bytesShaderTarget)
	panicErr(err)
}

type target struct {
	posX, posY int
	drawOpts   ebiten.DrawRectShaderOptions
}

func newTarget(posX, posY int) *target {
	target := &target{
		posX: posX,
		posY: posY,
		drawOpts: ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				// "TileRadius": float32(tileLength / 2.0),
			},
		},
	}

	target.drawOpts.GeoM.Translate(float64(posX)*tileLength, float64(posY)*tileLength)
	return target
}

func (t *target) update() {
	t.drawOpts.Uniforms["Time"] = timePassedSec
}
