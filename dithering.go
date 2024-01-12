package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"math"
)

func FirstCorrectIterarion(pixels [][]color.Color, size image.Point, bayerMatrixLvl int) {

	var bayerMatrix = NewNonNormalizedMatrix(bayerMatrixLvl)
	var col, row int = 0, 0

	for i := 0; i < size.X; i++ {
		row = i % int(math.Pow(2, (float64(bayerMatrixLvl+1))))
		for j := 0; j < size.Y; j++ {
			col = j % int(math.Pow(2, (float64(bayerMatrixLvl+1))))
			r, g, b, _ := pixels[i][j].RGBA()
			if int((r/256+g/256+b/256)/3) < bayerMatrix[row][col] {
				pixels[i][j] = color.Black
			} else {
				pixels[i][j] = color.White
			}
		}
	}
}

func secondIteration(pixels [][]color.Color, size image.Point, bayerMatrixLvl int) {

	var bayerMatrix = NewNormalizedMatrix(bayerMatrixLvl)
	var col, row int = 0, 0

	for i := 0; i < size.X; i++ {
		row = i % int(math.Pow(2, (float64(bayerMatrixLvl+1))))
		for j := 0; j < size.Y; j++ {
			col = j % int(math.Pow(2, (float64(bayerMatrixLvl+1))))
			r, g, b, a := pixels[i][j].RGBA()
			var factor float64 = bayerMatrix[row][col]
			var attempt float64 = (factor * float64(255/bayerMatrixLvl))

			// fmt.Printf("%.2f %.2f\n", factor, attempt)
			r = r * uint32(attempt)
			g = g * uint32(attempt)
			b = b * uint32(attempt)
			pixels[i][j] = findClosestColorFrom(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
			pixels[i][j] = findClosestColorFrom(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}
}

func findClosestColorFrom(pixel color.Color) color.Color {

	pallete := palette.Plan9
	// pallete := []color.Color{color.Black, color.White}
	pixelR, pixelG, pixelB, pixelA := pixel.RGBA()

	var closestDistance float64 = math.MaxFloat64
	var closest color.Color

	for _, v := range pallete {

		r, g, b, a := v.RGBA()
		distR := (r - pixelR) * (r - pixelR)
		// distR := math.Pow((float64(pixelR) - float64(r)), 2)
		distG := (g - pixelG) * (g - pixelG)
		// distG := math.Pow((float64(pixelG) - float64(g)), 2)
		distB := (b - pixelB) * (b - pixelB)
		// distB := math.Pow((float64(pixelB) - float64(b)), 2)
		distA := (a - pixelA) * (a - pixelA)
		// distA := math.Pow((float64(pixelA) - float64(a)), 2)
		sum := distR + distG + distB + distA

		var dist float64 = math.Sqrt(float64(sum))
		if dist <= closestDistance {
			closestDistance = dist
			closest = v
		}
	}
	// r, g, b, _ := closest.RGBA()
	// fmt.Printf("%d %d %d %.2f\n", r, g, b, closestDistance)
	return closest
}
