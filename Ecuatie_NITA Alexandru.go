package main

import (
	"fmt" 				//pentru input/output
	"html/template"  	//pentru a crea o pagina html
	"math"
	"net/http"			//pentru a crea server ul
	"strconv"			//pentru a transforma  datele  primite de la client, site web, din biti in string
)



func main() {

	tmpl := template.Must(template.ParseFiles("C:\\Users\\nital\\go\\src\\ecuatie\\index.html"))


	http.HandleFunc("/", func  (w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		a1, _ := strconv.ParseFloat(r.FormValue("a"), 64)
		b1, _ := strconv.ParseFloat(r.FormValue("b"), 64)
		c1, _ := strconv.ParseFloat(r.FormValue("c"), 64)


		if(a1 == 0 && b1 == 0){
			fmt.Fprintf(w,"Ecuatia nu are solutii.")
		}
		if(a1 == 0){
			fmt.Fprintf(w,"Ecuatia are o solutie: %f.",(c1/b1)*(-1))
		}

		sol1:= ((b1 * -1) - math.Sqrt((math.Pow(b1, 2))-(4*a1*c1)))/(2*a1)
		sol2:= ((b1 * -1) + math.Sqrt((math.Pow(b1, 2))-(4*a1*c1)))/(2*a1)


		fmt.Fprintf(w,"Ecuatia are doua solutii: %f , %f.",sol1,sol2)

	})

	http.ListenAndServe(":8080", nil)
}


func solutiileEcuatiei(a float64,b float64,c float64){

}