//go:build ignore

package main

const pi = 3.14159

var (
	TileRadius float
	Time       float
	Color      vec4
)

func Fragment(pos vec4, tex vec2, col vec4) vec4 {
	radius := TileRadius / 2.0 * (1.0 + sin(Time*pi))
	if distance(tex, vec2(TileRadius)) >= radius {
		return vec4(0.0)
	}

	return Color
}
