package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func ReadImgFileToTensor(filepath string) ([][]color.Color, image.Point) {
	//open image
	reader, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	img, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	// Convert image to tensor
	size := img.Bounds().Size()
	var tensor [][]color.Color

	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		tensor = append(tensor, y)
	}
	return tensor, size
}

// Convert tensor back to image
func WriteTensorToImgFile(tensor [][]color.Color, filepath string) {
	rect := image.Rect(0, 0, len(tensor), len(tensor[0]))
	nImg := image.NewRGBA(rect)

	for x := 0; x < len(tensor); x++ {
		for y := 0; y < len(tensor[0]); y++ {
			q := tensor[x]
			if q == nil {
				continue
			}
			p := tensor[x][y]
			if p == nil {
				continue
			}
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}

	fg, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Creating file:", err)
	}
	defer fg.Close()
	png.Encode(fg, nImg)
}
