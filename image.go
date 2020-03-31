package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/Nykakin/quantize"
)

func getImageFromFilePath(path string) (image.Image, error) {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	image, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	return image, err
}

func getPaletteFromImage(img image.Image, num int) []color.Color {
	quantizer := quantize.NewHierarhicalQuantizer()
	colors, _ := quantizer.Quantize(img, num)

	palette := make([]color.Color, len(colors))
	for index, clr := range colors {
		palette[index] = clr
	}

	return palette
}
