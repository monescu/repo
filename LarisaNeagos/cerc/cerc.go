package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

type point2D struct{
	x float64
	y float64
}

type minOrMax int
const (
	min minOrMax = 0
	max minOrMax = 1
)

func main() {
	http.HandleFunc("/", calculate)
	http.ListenAndServe(":9090", nil)
}

func calculate(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.ServeFile(w, r, "puncte.html")
	}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		n := len(r.PostForm) / 2

		var points []point2D
		distances := make([][]float64, n)

		for i := 0; i < n; i++ {
			distances[i] = make([]float64, n)
			var p point2D

			xName := fmt.Sprintf("x-%d", i)
			yName := fmt.Sprintf("y-%d", i)
			p.x, _ = strconv.ParseFloat(r.FormValue(xName), 64)
			p.y, _ = strconv.ParseFloat(r.FormValue(yName), 64)
			points = append(points, p)
		}

		makeDistancesMatrix(points, distances)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {

				fmt.Printf("| %-20g|", distances[i][j])
			}
			fmt.Println()
		}
		radius := getRadius(distances)
		fmt.Printf("Radius = %g\n", radius)
		fmt.Fprintf(w, "Radius = %g\n", radius)

	}
}

func makeDistancesMatrix(points []point2D, distances [][]float64){
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			distances[i][j] = math.Sqrt(math.Pow(points[i].x-points[j].x, 2) + math.Pow(points[i].y-points[j].y, 2))
			distances[j][i] = distances[i][j]
		}
	}
}

func getRadius(distances [][]float64) float64{
	arrayMax := make([]float64, len(distances))
	for i := 0; i < len(distances); i++ {
		arrayMax[i] = getMinOrMaxFromArray(distances[i], max)
	}
	radius := getMinOrMaxFromArray(arrayMax, min)
	return radius
}

func getMinOrMaxFromArray (array []float64, operation minOrMax) float64{
	result := array[0]
	for _, value := range array {
		if operation == max {
			if value > result {
				result = value
			}
		} else {
			if value < result{
				result = value
			}
		}
	}
	return result
}