package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

//Create a struct that holds information to be displayed in our HTML file
type Equation struct { 
	x_squared float64
	x_first   float64
	free_term float64
}

var result string

func solveEquation(equation Equation) string {
	delta := equation.x_first*equation.x_first - 4*(equation.x_squared*equation.free_term)
	if delta > 0 {
		x1 := (-equation.x_first + math.Sqrt(delta)) / 2 * equation.x_squared
		x2 := (-equation.x_first - math.Sqrt(delta)) / 2 * equation.x_squared
		result := fmt.Sprintf("x1:%d x2:%d", x1, x2)
		return result
	}
	if delta == 0 {
		x1 := (-equation.x_first + math.Sqrt(delta)) / 2 * equation.x_squared
		result := fmt.Sprint("x: %d", x1)
		return result
	}
	if delta < 0 {
		real := (-equation.x_first) / (2 * equation.x_squared)
		img := math.Sqrt(-delta) / (2 * equation.x_squared)
		cmplex := complex(real, img)
		cmplex2 := complex(real, -img)
		result := fmt.Sprintf("x1: %v x2: %v", cmplex, cmplex2)
		return result
	}
	return ""
}

//Go application entrypoint
func main() {

	equation := Equation{}
	templates := template.Must(template.ParseFiles("D:/facultate/golang/repo/Ruxanda_Tudor/Tema5/templates/equation-template.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "equation-template.html", equation); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/solveEquation", func(w http.ResponseWriter, r *http.Request) {
		if x_squared := r.FormValue("x_squared"); x_squared != "" {
			equation.x_squared, _ = strconv.ParseFloat(x_squared, 64)
		}
		if x_first := r.FormValue("x_first"); x_first != "" {
			equation.x_first, _ = strconv.ParseFloat(x_first, 64)
		}
		if free_term := r.FormValue("free_term"); free_term != "" {
			equation.free_term, _ = strconv.ParseFloat(free_term, 64)
		}
		result = solveEquation(equation)
		fmt.Println(result)
	})

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
