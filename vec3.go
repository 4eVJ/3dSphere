package main

import "math"

type vec3 struct {
	x, y, z float64
}

func (v vec3) add(other vec3) vec3 {
	return vec3{v.x + other.x, v.y + other.y, v.z + other.z}
}

func (v vec3) sub(other vec3) vec3 {
	return vec3{v.x - other.x, v.y - other.y, v.z - other.z}
}

func (v vec3) mul(other vec3) vec3 {
	return vec3{v.x * other.x, v.y * other.y, v.z * other.z}
}

func (v vec3) div(other vec3) vec3 {
	return vec3{v.x / other.x, v.y / other.y, v.z / other.z}
}

func (v vec3) normalize() vec3 {
	return v.div(vec3{vec3Lenght(v), vec3Lenght(v), vec3Lenght(v)})
}

func vec3Dot(a, b vec3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func vec3Normalize(v vec3) vec3 {
	return v.div(vec3{vec3Lenght(v), vec3Lenght(v), vec3Lenght(v)})
}

func vec3Mul(a, b vec3) vec3 {
	return vec3{a.x * b.x, a.y * b.y, a.z * b.z}
}

func vec3Div(a, b vec3) vec3 {
	return vec3{a.x / b.x, a.y / b.y, a.z / b.z}
}

func vec3Add(a, b vec3) vec3 {
	return vec3{a.x + b.x, a.y + b.y, a.z + b.z}
}

func vec3Sub(a, b vec3) vec3 {
	return vec3{a.x - b.x, a.y - b.y, a.z - b.z}
}

func vec3Negative(v vec3) vec3 {
	return vec3{-v.x, -v.y, -v.z}
}

func vec3Lenght(v vec3) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}
