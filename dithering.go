package main

import (
	"image"
	"image/color"
	"math"
)

// Just find the closest color to the pallete and swap it
func rgbFindClosest(pixels [][]color.Color, size image.Point) {
	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			r, g, b, a := pixels[i][j].RGBA()
			pixels[i][j] = findClosestColorFrom(r/256, g/256, b/256, a/256)
		}
	}
}

func blackWhiteBayer(pixels [][]color.Color, size image.Point, bayerMatrixLvl int) {

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

func rgbBayer(pixels [][]color.Color, size image.Point, bayerMatrixLvl int) {

	var bayerMatrix = NewNormalizedMatrix(bayerMatrixLvl)
	var col, row int = 0, 0

	for i := 0; i < size.X; i++ {
		row = i % int(math.Pow(2, (float64(bayerMatrixLvl+1)))) // x % N
		for j := 0; j < size.Y; j++ {
			col = j % int(math.Pow(2, (float64(bayerMatrixLvl+1)))) // y % N
			r, g, b, a := pixels[i][j].RGBA()
			var attempt float64 = ((bayerMatrix[row][col] - 0.5) * float64(255/bayerMatrixLvl)) // r * M(x % N, y % N)

			pixels[i][j] = findClosestColorFrom(
				r/256+uint32(attempt),
				g/256+uint32(attempt),
				b/256+uint32(attempt),
				a/256+uint32(attempt),
			)
		}
	}
}
