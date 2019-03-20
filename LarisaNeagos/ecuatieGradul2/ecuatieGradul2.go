package main

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func solve(a, b, c float64) string {
	var delta, x1, x2 float64
	delta = b*b - 4*a*c
	if a == 0 && b == 0 {
		return "Coeficientul lui x este 0"
	} else {
		if a == 0 {
			x1 = -c / b
			return "Ecuatia de grad 1 are solutia: " + strconv.FormatFloat(x1, 'f', 6, 64)
		} else {
			if b == 0 && (-c/a) >= 0 {
				x1 = math.Sqrt(-c / a)
				return "Ecuatia are solutia: " + strconv.FormatFloat(x1, 'f', 6, 64)
			} else {
				if a != 0 && b != 0 {
					delta = b*b - 4*a*c
					if delta < 0 {
						var realPart = -b / (2 * a)
						var imaginaryPart = math.Sqrt(math.Abs(delta)) / (2 * a)
						return "Ecuatia are solutii complexe  x1= " + strconv.FormatFloat(realPart, 'f', 6, 64) + " + " + strconv.FormatFloat(imaginaryPart, 'g', 3, 64) + "i si x2= " + strconv.FormatFloat(realPart, 'f', 6, 64) + " - " + strconv.FormatFloat(imaginaryPart, 'g', 3, 64) + "i"
					} else {
						if delta == 0 {
							x1 = -b / (2 * a)
							return "Ecuatia are solutia: " + strconv.FormatFloat(x1, 'f', 6, 64)
						} else {
							if delta > 0 {
								x1 = (-b + math.Sqrt(delta)) / (2 * a)
								x2 = (-b - math.Sqrt(delta)) / (2 * a)
								return "Ecuatia are solutia: x1= " + strconv.FormatFloat(x1, 'f', 6, 64) + " si x2= " + strconv.FormatFloat(x2, 'f', 6, 64)
							} else {
								return "Ecuatia nu are solutii"
							}
						}
					}
				}
			}

		}
	}
	return ""
}

type templateParams struct {
	Result string
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8081", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	params := templateParams{}

	if r.Method == "GET" {
		tmpl.Execute(w, params)
		return
	}

	a := r.FormValue("a")
	b := r.FormValue("b")
	c := r.FormValue("c")

	a1, err := strconv.ParseFloat(a, 64)
	b1, err := strconv.ParseFloat(b, 64)
	c1, err := strconv.ParseFloat(c, 64)

	if err == nil {
		result := solve(a1, b1, c1)
		params.Result = result
		tmpl.Execute(w, params)
	}
}
