//go:build ignore

package main

const pi = 3.14159

var (
	// TileRadius float
	Time float
)

func Fragment(pos vec4, tex vec2, col vec4) vec4 {

	r := 0.5 * (1 + sin(Time*pi))
	g := 0.5 * (1 + sin(Time*pi+pi/2.0))
	b := 0.5 * (1 + sin(Time*pi-pi/2.0))
	return vec4(r, g, b, 1.0)
}
