package main

import (
	"fmt"
	"math"
)

func calculateDistances(points []Point2D) map[Point2D][]float64 {
	var distances = map[Point2D][]float64{}

	for _, point := range points {
		for _, otherPoint := range points {
			if !point.equals(otherPoint) {
				distances[point] = append(distances[point], point.distance(otherPoint))
			}
		}
	}
	return distances
}

func calculateMax(line []float64) float64 {
	var max float64
	for _, nr := range line {
		if nr > max {
			max = nr
		}
	}
	return max
}

func calculateMin(maxims []float64) (float64, int) {
	var min = float64(math.MaxFloat64)
	var index int
	for i, nr := range maxims {
		if nr <= min {
			min = nr
			index = i
		}
	}
	return min, index
}

func main() {
	var points = []Point2D{
		{x: 1, y: 1},
		{x: 1, y: 2},
		{x: 1, y: 3},
		{x: 2, y: 1},
		{x: 2, y: 2},
		{x: 2, y: 3},
		{x: 3, y: 1},
		{x: 3, y: 2},
		{x: 3, y: 3},
		{x: 0, y: 0},
	}

	fmt.Println(points)

	var distances = calculateDistances(points)
	var maxims []float64

	for _, point := range points {
		maxims = append(maxims, calculateMax(distances[point]))
	}

	min, index := calculateMin(maxims)

	fmt.Println("Centre: ", points[index])
	fmt.Println("Radius: ", min)

}
