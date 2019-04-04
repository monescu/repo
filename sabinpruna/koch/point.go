package koch

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

//Point is used to draw Koch curve
type Point struct {
	X float64
	Y float64
}

var img = image.NewRGBA(image.Rect(0, 0, 4000, 4000))
var col = color.White

//Koch is used to generate Koch fractal for the nth degree starting with t1 and t2
func Koch(x1, y1, x2, y2 float64, n int, polarity bool) {

	if n == 0 {
		drawLine(img, int(x1), int(y1), int(x2), int(y2), col)
	} else {

		if x1 > x2 {
			x1, y1, x2, y2 = x2, y2, x1, y1

			if polarity == false {
				polarity = true
			} else {
				polarity = false
			}
		}

		xa, ya, xb, yb, xc, yc := getABCcoords(x1, y1, x2, y2, polarity)

		n = n - 1
		Koch(x1, y1, xa, ya, n, polarity)
		Koch(xa, ya, xb, yb, n, polarity)
		Koch(xb, yb, xc, yc, n, polarity)
		Koch(xc, yc, x2, y2, n, polarity)
	}
}

//DrawSnowflake makes a snowflake using Koch curve
func DrawSnowflake(x1, y1, x2, y2 float64, n int, polarity bool) {
	x3 := (x1 + x2 + math.Sqrt(3)*(y2-y1)) / 2
	y3 := (y1 + y2 - math.Sqrt(3)*(x2-x1)) / 2

	Koch(x1, y1, x2, y2, n, polarity)
	Koch(x2, y2, x3, y3, n, polarity)
	Koch(x3, y3, x1, y1, n, polarity)
}

//MakePng creates a pgn with a snowflake
func MakePng() {
	n := 5
	x1, y1, x2, y2 := 500.0, 3000.0, 3500.0, 3000.0

	DrawSnowflake(x1, y1, x2, y2, n, true)
	f, err := os.Create("snowflake.png")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	png.Encode(f, img)

	img = image.NewRGBA(image.Rect(0, 0, 4000, 4000))

	Koch(x1, y1, x2, y2, n, false)

	f, err = os.Create("koch.png")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	png.Encode(f, img)
}

//----------------------PRIVATES----------------------------------------

func getABCcoords(x1, y1, x2, y2 float64, polarity bool) (float64, float64, float64, float64, float64, float64) {
	var xa, ya, xb, yb, xc, yc float64

	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	L := distance / 3

	alfa := math.Asin((y2 - y1) / (3 * L))

	xa = x1 + L*math.Cos(alfa)
	ya = y1 + L*math.Sin(alfa)

	if polarity == true {
		xb = x1 + L*math.Sqrt(3)*math.Cos(alfa+(math.Pi/6))
		yb = y1 + L*math.Sqrt(3)*math.Sin(alfa+(math.Pi/6))
	} else {
		xb = x1 + L*math.Sqrt(3)*math.Cos(alfa-(math.Pi/6))
		yb = y1 + L*math.Sqrt(3)*math.Sin(alfa-(math.Pi/6))
	}

	xc = x1 + 2*L*math.Cos(alfa)
	yc = y1 + 2*L*math.Sin(alfa)

	return xa, ya, xb, yb, xc, yc
}

//uses the github/stephanebunel/bresenham algorithm to draw a line between 2 points
func drawLine(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
	var dx, dy, e, slope int

	// have lower point first
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {
	//line is point
	case x1 == x2 && y1 == y2:
		img.Set(x1, y1, col)

	//horizontal line
	case y1 == y2:
		for ; dx != 0; dx-- {
			img.Set(x1, y1, col)
			x1++
		}
		img.Set(x1, y1, col)

	//vertical line
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			img.Set(x1, y1, col)
			y1++
		}
		img.Set(x1, y1, col)

	//diagonal line
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				y1--
			}
		}
		img.Set(x1, y1, col)

	//wider than high
	case dx > dy:
		if y1 < y2 {
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		img.Set(x2, y2, col)

	// higher than wide.
	default:
		if y1 < y2 {
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, col)
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, col)
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		img.Set(x2, y2, col)
	}
}

//----------------------------------------------------------------------
