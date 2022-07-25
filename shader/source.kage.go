//go:build ignore

package main

var (
	Radius float
	Color  vec4
)

func Fragment(pos vec4, tex vec2, col vec4) vec4 {
	if distance(tex, vec2(Radius)) >= Radius {
		return vec4(0.0)
	}
	return Color

}
