package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)
func rezolvare(a, b, c float64) string {
	var Delta, x1, x2 float64
	Delta = b*b - 4*a*c
	if a == 0 && b == 0 {
		return "Coeficientul lui x este 0" //cazul in care coeficientul este 0=>ec gr1
	} else {
		if a == 0 {
			x1 = -c / b
			return "Nu este de ecuatie de gradul 2 si avem solutia pentru ecuatie de gradul 1: " + strconv.FormatFloat(x1, 'f', 6, 64)
		} else {
			if b == 0 && (-c/a) >= 0 {
				x1 = math.Sqrt(-c / a)
				return "Ecuatia de gradul 2 are solutia: " + strconv.FormatFloat(x1, 'f', 6, 64)
			} else {
				//cazul in care am solutii complexe
				if a != 0 && b != 0 {
					Delta = b*b - 4*a*c
					if Delta < 0 {//Cazul pentru Delta >0
						var real = -b / (2 * a)
						var imaginar = math.Sqrt(math.Abs(Delta)) / (2 * a)
						return "Ecuatia de gradul 2 are solutii complexe  x1= " + strconv.FormatFloat(real, 'f', 6, 64) + " + " + strconv.FormatFloat(imaginar, 'g', 3, 64) + "i " + "si x2= " + strconv.FormatFloat(real, 'f', 6, 64) + " - " + strconv.FormatFloat(imaginar, 'g', 3, 64) + "i"
					} else {
						if Delta > 0 { //Cazul pentru Delta >0
							x1 = (-b + math.Sqrt(Delta)) / (2 * a)
							x2 = (-b - math.Sqrt(Delta)) / (2 * a)
							return "Ecuatia are solutia: x1= " + strconv.FormatFloat(x1, 'f', 6, 64) + " si x2= " + strconv.FormatFloat(x2, 'f', 6, 64)

						} else { //Cazul pentru Delta =0
							if Delta == 0 {
								x1 = -b / (2 * a)
								return "Ecuatia de gradul 2 are solutia: " + strconv.FormatFloat(x1, 'f', 6, 64)

							} else {
								return "Ecuatia de gradul 2 nu are solutii!"
							}
						}
					}
				}
			}

		}
	}
	return ""
}

type templateParam struct {
	Rezultat string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

	a := r.FormValue("a")
	b := r.FormValue("b")
	c := r.FormValue("c")

	x, err := strconv.ParseFloat(a, 64)
	y, err := strconv.ParseFloat(b, 64)
	z, err := strconv.ParseFloat(c, 64)

	if err == nil {
		rezultat := rezolvare(x, y,z)
		fmt.Println(rezultat)
		Rezultat = rezultat
		tmpl.Execute(w, nil)
	}
}