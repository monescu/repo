package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func calculateDistance(ax float64, bx float64, ay float64, by float64) float64 {
	return math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
}

func createPointsMatrix(numberOfPoints int, points [] Point) [][]float64 {

	matrixPoints := make([][] float64, numberOfPoints)
	for i := 0; i < numberOfPoints; i++ {
		matrixPoints[i] = make([]float64,numberOfPoints)
		for j := 0; j < numberOfPoints; j++ {
			if i == j {
				matrixPoints[i][j] = 0
			} else {
				matrixPoints[i][j] = calculateDistance(points[i].x, points[j].x, points[i].y, points[j].y)
			}
		}
	}
	return matrixPoints
}

func selectMaxElementsFromMatrix(initialMatrix [][]float64, n int) []float64 {
	var maxElements = make([]float64, n)
	var maxLineElem float64
	for i := 0; i < n; i++ {
		maxLineElem = initialMatrix[i][0]
		for j := 1; j < n-1; j++ {
			if initialMatrix[i][j] > maxLineElem {
				maxLineElem = initialMatrix[i][j]
			}
		}
		maxElements[i] = maxLineElem
	}

	return maxElements
}

func minElementFromList(listOfMaxElements[] float64) float64{
	var minElem float64 = 100000
	for i:=0;i<len(listOfMaxElements);i++{
		if listOfMaxElements[i]<minElem {
			minElem = listOfMaxElements[i]
		}
	}
	return minElem
}
func main() {

	var points = []Point{
		{x: 1, y: 3},
		{x: 3, y: 2},
		{x: 1, y: 4},
		{x: 1, y: 6},
		{x: 3, y: 3},
	}

	matrix := make([][] float64, len(points));
	for i := 0; i < len(points); i++ {
		matrix[i] = make([]float64, len(points));
	}

	var maxesArray []float64
	var radius float64

	matrix = createPointsMatrix(len(points),points)
	maxesArray = selectMaxElementsFromMatrix(matrix,len(points))
	radius = minElementFromList(maxesArray)
	fmt.Println("Centrul cercului este ", radius)
	fmt.Println("Puncte:")
	fmt.Println(points)
	var convertRadiusToInt = int(radius)
	fmt.Print("Centrul cercului se afla la coordonatele:",points[convertRadiusToInt])

}
