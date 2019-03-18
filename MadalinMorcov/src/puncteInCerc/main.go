package main

import (
	"fmt"
	"math"
	"puncteInCerc/model"
)

func calculateDistances(points []model.Point2D) map[model.Point2D][]float64 {
	var distances = map[model.Point2D][]float64{}

	for _, point := range points {
		for _, otherPoint := range points {
			if !point.Equals(otherPoint) {
				distances[point] = append(distances[point], point.Distance(otherPoint))
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
	var points = []model.Point2D{
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 3},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 2, Y: 3},
		{X: 3, Y: 1},
		{X: 3, Y: 2},
		{X: 3, Y: 3},
		{X: 0, Y: 0},
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
