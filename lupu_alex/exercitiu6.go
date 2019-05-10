package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

var white = color.RGBA{255, 255, 255, 255}
var black = color.RGBA{0, 0, 0, 255}
var angle60 = math.Pi / 3;

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 600, 400))
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	x1:=float64(150)
	y1:=float64(200)
	x2:=float64(350)
	y2:=float64(200)
	x3 := x1 + (x2-x1) * math.Cos(angle60) + (y2-y1) * math.Sin(angle60)
	y3 := y1 - (x2-x1) * math.Sin(angle60) + (y2-y1) * math.Cos(angle60)

	koch(x2, y2, x1, y1, 0, *img)
	koch(x3, y3, x2, y2, 0, *img)
	koch(x1, y1, x3, y3, 0, *img)
	
	file, err := os.Create("theImage.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func drawline(x0, y0, x1, y1 int, img image.RGBA) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		img.Set(x0, y0, white)
		img.Set(x0+1, y0, white)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func koch(x1, y1, x2, y2 float64, iter int, img image.RGBA) {

	x3 := (x1*2 + x2) / 3
	y3 := (y1*2 + y2) / 3
	x4 := (x1 + x2*2) / 3
	y4 := (y1 + y2*2) / 3
	x5 := x3 + (x4-x3) * math.Cos(angle60) + (y4-y3) * math.Sin(angle60)
	y5 := y3 - (x4-x3) * math.Sin(angle60) + (y4-y3) * math.Cos(angle60)

	if iter > 0 {
		iter--
		koch(x1, y1, x3, y3, iter, img)
		koch(x3, y3, x5, y5, iter, img)
		koch(x5, y5, x4, y4, iter, img)
		koch(x4, y4, x2, y2, iter, img)
	} else {
		drawline(int(x1), int(y1), int(x3), int(y3), img)
		drawline(int(x3), int(y3), int(x5), int(y5), img)
		drawline(int(x5), int(y5), int(x4), int(y4), img)
		drawline(int(x4), int(y4), int(x2), int(y2), img)
	}

}