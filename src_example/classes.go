package main

import (
	"fmt"
	"image/color"
)

type ColoredPoint struct {
	color.Color // Anonymous field (embedding)
	x, y int // Named fields (aggregation)
}

func NewColoredPoint(x int, y int, color color.Color) *ColoredPoint {
	return &ColoredPoint{x: x, y: y, Color: color}
}

func (c ColoredPoint)doSomething() int{
	return c.x + c.y
}


type Count int
func (count *Count) Increment() { *count++ }
func (count *Count) Decrement() { *count-- }
func (count Count) IsZero() bool { return count == 0 }


func main() {

	var count Count
	i := int(count)
	count.Increment()
	j := int(count)
	count.Decrement()
	k := int(count)
	fmt.Println(count, i, j, k, count.IsZero())



}
