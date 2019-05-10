package main

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type ExuationDetails struct {
	firstCoef string
	secondCoef string
	thirdCoef string
}

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		values := ExuationDetails{
			firstCoef:   r.FormValue("firstCoef"),
			secondCoef: r.FormValue("secondCoef"),
			thirdCoef: r.FormValue("thirdCoef"),
		}

		a, _ := strconv.Atoi(values.firstCoef)
		b, _ := strconv.Atoi(values.secondCoef)
		c, _ := strconv.Atoi(values.thirdCoef)

		delta := math.Sqrt(math.Pow(float64(b), 2) - (4*float64(a)*float64(c)))
		x1 := (-float64(b) + delta)/ (2 * float64(a))
		x2 := (-float64(b) - delta)/ (2 * float64(a))

		tmpl.Execute(w, struct{
			Success bool
			FirstX float64
			SecondX float64
		}{true, x1, x2})
	})

	http.ListenAndServe(":8080", nil)
}