package main

import (
	"fmt" 				
	"html/template"  	
	"math"
	"net/http"			
	"strconv"			
)

func main() {
	
	tmpl := template.Must(template.ParseFiles("C:\\Users\\andreea\\go\\src\\ecuatie\\ecuatie.html"))

	http.HandleFunc("/", func  (w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		a1, _ := strconv.ParseFloat(r.FormValue("a"), 64)
		b1, _ := strconv.ParseFloat(r.FormValue("b"), 64)
		c1, _ := strconv.ParseFloat(r.FormValue("c"), 64)

		x1,x2:=solutii(a1,b1,c1)


		fmt.Fprintf(w,"Solutiile ecuatiei sunt:\n X1= %f \n X2= %f", x1, x2 )
	})


	http.ListenAndServe(":8000", nil)
}


func solutii(a float64,b float64, c float64)(sol1 float64, sol2 float64){
	if(a == 0 && b == 0) {

		sol1=math.NaN()
		sol2=math.NaN()
		return sol1,sol2
	}

	if(a == 0){
		sol1 = (c/b)*(-1)
		sol2 = (c/b)*(-1)
		return sol1,sol2
	}
	sol1= ((b * -1) - math.Sqrt((math.Pow(b, 2))-(4*a*c)))/(2*a)
	sol2= ((b * -1) + math.Sqrt((math.Pow(b, 2))-(4*a*c)))/(2*a)
	return sol1,sol2
}


