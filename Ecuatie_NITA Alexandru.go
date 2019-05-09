package main

import (
	"fmt" 				//pentru input/output
	"html/template"  	//pentru a crea o pagina html
	"math"
	"net/http"			//pentru a crea server ul
	"strconv"			//pentru a transforma  datele  primite de la client, site web, din biti in string
)

//  :=  ii da o valoare unei variabile si cu : ii defineste si tipul de data in functie de ce valoare primeste

func main() {
	//pagina de start pt a putea comunica cu serverul
	tmpl := template.Must(template.ParseFiles("C:\\Users\\nital\\go\\src\\ecuatie\\index.html"))

	//handleFunc se ocupa de gestionarea cererilor/ oferirea raspunsurilor de catre site/pentru site
	// reaponseWriter este cel care trimite raspunsul la client, site web,
	// iar request este cel care trimite cererile la server cand apasam pa submit
	//post este folosit pt a trimite date la server
	// daca cerinta clientului,ce se intampla in pag html,  nu este un post, atunci serverul nu face nimic, (return)
	http.HandleFunc("/", func  (w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
//a1 ia valoarea celui care are numele "a" in html , si o transforma din string trimis pe 64 de biti in float
//in caul in care nu reuseste sa le transforme , rezulta o eroare pe care nu o folosim  "_"
		a1, _ := strconv.ParseFloat(r.FormValue("a"), 64)
		b1, _ := strconv.ParseFloat(r.FormValue("b"), 64)
		c1, _ := strconv.ParseFloat(r.FormValue("c"), 64)

		//w este raspunsul pe care il trimite serverul , transformat in string de fprintf
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
//pana aici sunt regulile si functiile serverului, site ului

//acum incepem sa asultam cererile de la client , serverul este la adresa ip localhost:8000,
// handler nil este cel de unde se trimit cererile catre server , dar nu am implementat nimic pt el si se va folosi cel
//oferit de go by default de aia e nil

	http.ListenAndServe(":8080", nil)
}


func solutiileEcuatiei(a float64,b float64,c float64){

}