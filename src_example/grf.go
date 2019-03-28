package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// fractali plecand de la un pattern: curba lui Koch, triunghi echilateral
// desenarea la un pas se va face cu o culoare data



var img = image.NewRGBA(image.Rect(0, 0, 100, 100))
var col color.Color

func main() {
	col = color.RGBA{255, 0, 0, 255} // Red
	HLine(10, 20, 80)
	col = color.RGBA{0, 255, 0, 255} // Green
	Rect(10, 10, 80, 50)

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

// HLine draws a horizontal line
func HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
	HLine(x1, y1, x2)
	HLine(x1, y2, x2)
	VLine(x1, y1, y2)
	VLine(x2, y1, y2)
}