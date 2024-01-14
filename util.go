package main

import (
	"image/color"
	"math"
	"strconv"
)

func findClosestColorFrom(pixelR uint32, pixelG uint32, pixelB uint32, pixelA uint32) color.Color {

	// pallete := palette.Plan9
	// pallete := []color.Color{HexColor("#1a5fb4"), HexColor("#ffbe6f")}
	pallete := []color.Color{color.Black, color.White}

	var closestDistance float64 = math.MaxFloat64
	var closest color.Color

	for _, v := range pallete {

		r, g, b, a := v.RGBA()
		r = r / 256
		g = g / 256
		b = b / 256
		a = a / 256
		distR := (r - pixelR) * (r - pixelR)
		distG := (g - pixelG) * (g - pixelG)
		distB := (b - pixelB) * (b - pixelB)
		distA := (a - pixelA) * (a - pixelA)

		sum := distR + distG + distB + distA

		var dist float64 = math.Sqrt(float64(sum))
		if dist <= closestDistance {
			closestDistance = dist
			closest = v
		}
	}
	return closest
}

func HexColor(hex string) color.RGBA {
	values, _ := strconv.ParseUint(string(hex[1:]), 16, 32)
	return color.RGBA{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
		A: 255,
	}
}
