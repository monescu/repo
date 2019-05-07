package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))

	solutions := make(chan float64, 2)

	http.HandleFunc("/", func  (w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		a1, _ := strconv.ParseFloat(r.FormValue("a"), 64)
		b1, _ := strconv.ParseFloat(r.FormValue("b"), 64)
		c1, _ := strconv.ParseFloat(r.FormValue("c"), 64)

		go s1(a1, b1, c1, solutions)
		go s2(a1, b1, c1, solutions)

		fmt.Fprintf(w,"Solutions:\n X1= %f \n X2= %f", <-solutions, <-solutions )
	})

	http.ListenAndServe(":8000", nil)
}

func handler (w http.ResponseWriter, r *http.Request){

}

func s1(a float64, b float64, c float64, solutions chan float64) {
	var div float64
	var quo float64

	div = (b * -1) + math.Sqrt((math.Pow(b, 2))-(4*a*c))
	quo = div / (2 * a)

	solutions <- quo
}

func s2(a float64, b float64, c float64, solutions chan float64) {
	var div float64
	var quo float64

	div = (b * -1) - math.Sqrt((math.Pow(b, 2))-(4*a*c))
	quo = div / (2 * a)

	solutions <- quo
}