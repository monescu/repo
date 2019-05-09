package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"math"
)

// fractali plecand de la un pattern: curba lui Koch, triunghi echilateral
// desenarea la un pas se va face cu o culoare data

var img = image.NewRGBA(image.Rect(0, 0, 100, 100))
var col color.Color

func main() {
	col = color.RGBA{255, 0, 0, 255} // Red
	//HLine(10, 20, 80)
	Koch(0, 0, 100, 0, 1)
	//col = color.RGBA{0, 255, 0, 255} // Green
	//Rect(10, 10, 80, 50)

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

func DrawLine(x1, y1, x2, y2 int) {
	dx, dy := x2-x1, y2-y1
	a := float64(dy) / float64(dx)
	b := int(float64(y1) - a*float64(x1))

	img.Set(x1, y1, col)
	for x := x1 + 1; x <= x2; x++ {
		y := int(a*float64(x)) + b
		img.Set(x, y, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
	HLine(x1, y1, x2)
	HLine(x1, y2, x2)
	VLine(x1, y1, y2)
	VLine(x2, y1, y2)
}

func Koch(x1, y1, x2, y2 float64, n int)  {
	angle := math.Atan((y2 - y1) / (x2 - x1))
	l := Dist(x1, y1, x2, y2) / 3
	xa := x1 + l * math.Cos(angle)
	ya := y1 + l * math.Sin(angle)
	xb := x1 + l * math.Sqrt(3) * math.Cos(angle + math.Pi / 6)
	yb := y1 + l * math.Sqrt(3) * math.Sin(angle + math.Pi / 6)
	xc := x1 + 2 * l * math.Cos(angle)
	yc := y1 + 2 * l * math.Sin(angle)
	if n > 0 {
		Koch(x1, y1, xa, ya, n - 1)
		Koch(xa, ya, xb, yb, n - 1)
		Koch(xb, yb, xc, yc, n - 1)
		Koch(xc, yc, x2, y2, n - 1)
	} else {
		DrawLine(int(x1), int(y1), int(x2), int(y2))
	}

}

func Dist(x1, y1, x2, y2 float64) float64{
	dist := math.Sqrt((x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1))
	return dist
}
