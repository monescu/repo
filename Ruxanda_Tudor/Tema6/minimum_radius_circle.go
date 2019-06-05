package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func calculateDistanceVector(n int, pointsVector []Point) [][]float64 {

	distanceVector := make([][]float64, n)
	var currentValue float64 = 0
	for i := 0; i < n; i++ {
		distanceVector[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j { 
				distanceVector[i][j] = currentValue
			} else {
				distanceVector[i][j] = calculateDistance(pointsVector[i], pointsVector[j])
			}
		}
	}
	return distanceVector
}

func getMaximumList(matrix [][]float64, n int) []float64 {

	var maximumValuesArray = make([]float64, n)
	var intermediaryMaximumValue float64
	for i := 0; i < n; i++ {
		intermediaryMaximumValue = matrix[i][0]
		for j := 1; j < n-1; j++ {
			if matrix[i][j] > intermediaryMaximumValue {
				intermediaryMaximumValue = matrix[i][j]
			}
		}
		maximumValuesArray[i] = intermediaryMaximumValue
	}
	return maximumValuesArray

}

func getMinumumDistance(n int, maximumDistancesList []float64) float64 {

	var minim float64 = maximumDistancesList[0]
	for index := 1; index < len(maximumDistancesList); index++ {
		if maximumDistancesList[index] < minim {
			minim = maximumDistancesList[index]
		}
	}
	return minim
}

func calculateDistance(firstPoint Point, secondPoint Point) float64 {
	return math.Sqrt((secondPoint.y-firstPoint.y)*(secondPoint.y-firstPoint.y) + (secondPoint.x-firstPoint.x)*(secondPoint.x-firstPoint.x))
}

func main() {
	var points = []Point{
		{x: 4, y: 0},
		{x: 3, y: 1},
		{x: 2, y: 2},
		{x: 1, y: 3},
		{x: 0, y: 4},
	}
	//generate empty matrix of len len(points)x len(points)
	matrix := make([][]float64, len(points))
	for i := 0; i < len(points); i++ {
		matrix[i] = make([]float64, len(points))
	}
	var maxArray []float64
	var radius float64
	matrix = calculateDistanceVector(len(points), points)
	maxArray = getMaximumList(matrix, len(matrix))
	radius = getMinumumDistance(len(points), maxArray)
	var radiusInteger = int(radius)
	fmt.Println("Center coordinates:", points[radiusInteger])

}
