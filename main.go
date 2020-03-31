package main

import (
	"flag"
	"fmt"
	"image/color"
)

func main() {
	numberFlag := flag.Int("number", 5, "number of colors to include in palette")
	imgFlag := flag.String("file", "", "target image to extract colors from")
	flag.Parse()

	if *imgFlag == "" {
		fmt.Println("USAGE: $ paleta -f <filename.jpg> [-n int (default: 5)]")
		return
	}

	img, err := getImageFromFilePath(*imgFlag)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	palette := getPaletteFromImage(img, *numberFlag)
	for _, color := range palette {
		fmt.Println(hexColor(color))
	}
}

func hexColor(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}
