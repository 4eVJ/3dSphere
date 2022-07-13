package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	fhp := flag.Int("fh", 19, "fonts character height")
	fwp := flag.Int("fw", 9, "fonts character width")

	flag.Parse()

	fw := float64(*fwp)
	fh := float64(*fhp)

	width, height, err := terminal.GetSize(0)
	if err != nil {
		log.Fatal(err)
	}

	pixelAspect := fw / fh
	screen := make([]rune, width*height)
	aspect := float64(width) / float64(height)
	gradient := []rune{' ', '.', '-', ':', ';', '+', '>', '*', '&', '%', '#', '@'}
	gradientSize := len(gradient) - 1
	for t := 0; t < 100000; t++ {
		light := vec3{math.Sin(float64(t) * 0.001), math.Cos(float64(t) * 0.001), -1.0}.normalize()
		// light := vec3{-0.5, 0.5, -1.0}.normalize()

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				uv := vec2Sub(vec2Mul(vec2Div(vec2{float64(i), float64(j)}, vec2{float64(width), float64(height)}), vec2{2, 2}), vec2{1, 1})

				uv.x *= aspect * pixelAspect

				ro := vec3{-2, 0, 0}
				rd := vec3Normalize(vec3{1, uv.x, uv.y})
				intersection := vec2Sphere(ro, rd, 1)

				pixel := ' '
				color := 0
				if intersection.x > 0 {
					itPoint := vec3Add(ro, vec3Mul(rd, vec3{intersection.x, intersection.x, intersection.x}))
					n := itPoint.normalize()
					diff := vec3Dot(n, light)
					color = int(diff * 17)
				}
				color = clamp(color, 0, gradientSize)
				pixel = gradient[color]
				screen[i+j*width] = pixel
			}
		}
		fmt.Println(string(screen))
		// time.Sleep(time.Millisecond * 75)
	}

}

func clamp(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
