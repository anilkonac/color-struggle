//go:build ignore

package main

const (
	pi          = 3.14159
	breathRatio = 1.0 / 5.0
)

var (
	TileRadius  float
	Time        float
	BreathSpeed float
	Color       vec4
)

func Fragment(pos vec4, tex vec2, col vec4) vec4 {
	radius := TileRadius * ((1.0 - breathRatio) + breathRatio*sin(BreathSpeed*Time*pi))
	if distance(tex, vec2(TileRadius)) >= radius {
		return vec4(0.0)
	}

	return Color
}
