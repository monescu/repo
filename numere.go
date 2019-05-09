package main

import (
	"fmt" 				
	"html/template"  	
	"math"
	"net/http"			
	"strconv"			
)

func main() {
	
	tmpl := template.Must(template.ParseFiles("C:\\Users\\turto\\numere\\index.html"))

	http.HandleFunc("/", func  (w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		a1, _ := strconv.ParseFloat(r.FormValue("a"), 64)
		b1, _ := strconv.ParseFloat(r.FormValue("b"), 64)
		c1, _ := strconv.ParseFloat(r.FormValue("c"), 64)

		sol1:= ((b1 * -1) - math.Sqrt((math.Pow(b1, 2))-(4*a1*c1)))/(2*a1)
		sol2:= ((b1 * -1) + math.Sqrt((math.Pow(b1, 2))-(4*a1*c1)))/(2*a1)


		fmt.Fprintf(w,"Solutiile ecuatiei: %f , %f.", sol1, sol2 )
	})

	http.ListenAndServe(":8000", nil)
}
