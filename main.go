package main

import (
	"flag"
	_ "image/jpeg"
	_ "image/png"
)

func main() {

	// Each lvl n equals 2^(n+1) matrix.
	lvlFlag := flag.Int("lvl", 3, "Matrix Lvl")
	flag.Parse()
	bayerMatrixLvl := *lvlFlag

	pixels, size := ReadImgFileToTensor("input.jpeg")
	// Image transformation
	FirstCorrectIterarion(pixels, size, bayerMatrixLvl)
	// secondIteration(pixels, size, bayerMatrixLvl)

	WriteTensorToImgFile(pixels, "output.jpeg")
}
