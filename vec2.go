package main

import "math"

type vec2 struct {
	x, y float64
}

func (v vec2) add(other vec2) vec2 {
	return vec2{v.x + other.x, v.y + other.y}
}

func (v vec2) sub(other vec2) vec2 {
	return vec2{v.x - other.x, v.y - other.y}
}

func (v vec2) mul(other vec2) vec2 {
	return vec2{v.x * other.x, v.y * other.y}
}

func (v vec2) div(other vec2) vec2 {
	return vec2{v.x / other.x, v.y / other.y}
}

func vec2Sphere(ro, rd vec3, r float64) vec2 {
	b := vec3Dot(ro, rd)
	c := vec3Dot(ro, ro) - r*r
	h := b*b - c
	if h < 0 {
		return vec2{-1, -1}
	}
	h = math.Sqrt(h)
	return vec2{-b - h, -b + h}
}

func vec2Sub(a, b vec2) vec2 {
	return vec2{a.x - b.x, a.y - b.y}
}

func vec2Mul(a, b vec2) vec2 {
	return vec2{a.x * b.x, a.y * b.y}
}

func vec2Div(a, b vec2) vec2 {
	return vec2{a.x / b.x, a.y / b.y}
}

func vec2Lenght(v vec2) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
