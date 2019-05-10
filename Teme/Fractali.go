package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// fractali -se pleaca de la curba lui Koch, triunghi echilateral
// desenarea la un pas se va face cu o  culoare care este  data



var img = image.NewRGBA(image.Rect(0, 0, 550, 300))
var col color.Color

func main() {
	col = color.RGBA{255, 100, 255, 255} // purple
	fmt.Println("Student Sraier Alina-Gabriela anul III ID")
	fmt.Println("Desenul este afisat in desen.png")
	angle := math.Pi / 3 // 60 grade
	x1:=float64(150)
	y1:=float64(200)
	x2:=float64(350)
	y2:=float64(200)
	x3 := x1 + (x2-x1) * math.Cos(angle) + (y2-y1) * math.Sin(angle)
	y3 := y1 - (x2-x1) * math.Sin(angle) + (y2-y1) * math.Cos(angle)
//aici modific numarul de iteratii
	koch(x2, y2, x1, y1, 2)
	koch(x3, y3, x2, y2, 3)
	koch(x1, y1, x3, y3, 1)


	f, err := os.Create("desen.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Bresenham's algorithm de pe wikipedia
func drawline(x0, y0, x1, y1 int) {
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
		img.Set(x0, y0, col)
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

//functia Koch
func koch(x1, y1, x2, y2 float64, iter int) {
	angle := math.Pi / 3 // unghi de 60 grade

	x3 := (x1*2 + x2) / 3
	y3 := (y1*2 + y2) / 3
	x4 := (x1 + x2*2) / 3
	y4 := (y1 + y2*2) / 3
	x5 := x3 + (x4-x3) * math.Cos(angle) + (y4-y3) * math.Sin(angle)
	y5 := y3 - (x4-x3) * math.Sin(angle) + (y4-y3) * math.Cos(angle)

	if iter > 0 {//iteratii
		iter--
		koch(x1, y1, x3, y3, iter)
		koch(x3, y3, x5, y5,iter)
		koch(x5, y5, x4, y4,iter)
		koch(x4, y4, x2, y2,iter)
	} else { //deseneaza
		drawline(int(x1), int(y1), int(x3), int(y3))
		drawline(int(x3), int(y3), int(x5), int(y5))
		drawline(int(x5), int(y5), int(x4), int(y4))
		drawline(int(x4), int(y4), int(x2), int(y2))
	}

}
