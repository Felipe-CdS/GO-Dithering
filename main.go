package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {

	// Each lvl n equals 2^(n+1) matrix.
	lvlFlag := flag.Int("lvl", 3, "Matrix Lvl")
	inputPathFlag := flag.String("i", "", "Input Name")
	outputPathFlag := flag.String("o", "", "Output Name")
	flag.Parse()
	bayerMatrixLvl := *lvlFlag
	if *inputPathFlag == "" || *outputPathFlag == "" {
		fmt.Println("Please provide correct input and output names with the -i -o flags.")
		os.Exit(1)
	}
	fmt.Printf("Generating image with a %d level bayer matrix.\n", bayerMatrixLvl)

	pixels, size := ReadImgFileToTensor(fmt.Sprintf("imageIO/%s", *inputPathFlag))
	// Image transformation
	rgbBayer(pixels, size, bayerMatrixLvl)

	WriteTensorToImgFile(pixels, fmt.Sprintf("imageIO/%s", *outputPathFlag))
}
