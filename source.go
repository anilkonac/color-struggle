package main

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shader/source.kage.go
	bytesShaderSource []byte
	shaderSource      *ebiten.Shader
)

func init() {
	var err error
	shaderSource, err = ebiten.NewShader(bytesShaderSource)
	panicErr(err)
}

type source struct {
	color.RGBA
	drawOpt ebiten.DrawRectShaderOptions
}

func newSource(x, y int, color color.RGBA) *source {
	source := &source{
		drawOpt: ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				"Radius": float32(tileLength / 4.0),
				"Color":  []float32{float32(color.R) / 255.0, float32(color.G) / 255.0, float32(color.B) / 255.0, float32(color.A) / 255.0},
			},
		},
	}

	source.drawOpt.GeoM.Translate(tileLength*(float64(x)+1.0/4.0), tileLength*(float64(y)+1.0/4.0))

	return source
}
