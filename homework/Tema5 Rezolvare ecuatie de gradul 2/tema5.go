package main
// Tema 05. Rezolvati ecuatia de gradul 2 folosind o aplicatie web Google GO
// Claudia Marinache - 10LD561
import (
	"html/template"
	"net/http"
	"log"
	"strconv"
	"math"
	"fmt"
)
type Variabile struct {
	numar_a string
	numar_b string
	numar_c string
	Delta	string
	Mesaj_text	string
	TitluPagina	string
	X1		string
	X2		string
	A		string
	B		string
	C		string
}

func main() {
  log.Print("folositi Chrome pentru a accesa http://localhost:8080/")
  http.HandleFunc("/", PaginaPrincipala)
  http.HandleFunc("/rezultate", AfisareRezultate)
  log.Fatal(http.ListenAndServe(":8080", nil)) // pentru accesare se deschide pagina web cu adresa http://localhost:8080
}

func PaginaPrincipala(w http.ResponseWriter, r *http.Request){  // genereaza pagina principala folosind ecuatieGDR2.html
  Titlu := "Ecuatie gradul 2 - main"

  Variabile_in_html := Variabile{
		TitluPagina: Titlu,
		Mesaj_text: "",
		X1: "",
		X2: "",
		Delta: "",
    }
	
	t, err := template.ParseFiles("ecuatieGDR2.html")
    if err != nil { // daca sunt erori se vor afisa pe consola
  	  log.Print("eroare: ", err)
  	}
	err = t.Execute(w, Variabile_in_html) 
    if err != nil { // daca sunt erori se vor afisa pe consola
  	  log.Print("eroare la executie: ", err)
  	}
}

func AfisareRezultate(w http.ResponseWriter, r *http.Request){  // genereaza pagina de raspuns ecuatieGDR2.html/rezultate
  r.ParseForm()

  // deoarece functia math.Sqrt() intoarce ca rezultat un float64, convertim toate cele 3 stringuri (a, b, c) la float64, pentru a fi compatibile cu calculele ulterioare
  a, _ := strconv.ParseFloat(r.Form.Get("numar_a"), 64)  //deoarece strconv.ParseFloat intoarce doi parametrii:  rezultat, erroare
  b, _ := strconv.ParseFloat(r.Form.Get("numar_b"), 64) 
  c, _ := strconv.ParseFloat(r.Form.Get("numar_c"), 64)
  
  var delta = math.Pow(b, 2) - 4 * a * c		// delta = b^2 - 4*a*c
  var x1 = (-1*b + math.Sqrt(delta))/(2*a)		// x1=(-b + SQRT(delta))/(2*a)
  var x2 = (-1*b - math.Sqrt(delta))/(2*a)		// x2=(-b - SQRT(delta))/(2*a)
  
  log.Printf("\n")		// pentru debug se afiseaza parametrii de calcul in consola
  log.Printf("a= %f", a)
  log.Printf("b= %f", b)
  log.Printf("c= %f", c)  
  log.Printf("delta= %f", delta)
  log.Printf("x1= %f", x1)
  log.Printf("x2= %f", x2)

  // declarare si initializare stringuri rezultate ce vor fi afisate ulterior in pagina html
  var mesaj=""
  var mesaj_delta="delta = " + fmt.Sprintf("%f", delta)
  var mesaj_x1=""
  var mesaj_x2=""
  
  if (a ==0 && b==0 && c == 0) || a == 0 {  // daca a=0 atunci ecuatia va fi de gradul 1, liniara
	  mesaj = "Valorile introduse nu sunt valide!"
	  mesaj_delta = ""
  } else if delta < 0 {
	  mesaj = "Ecuatia NU are solutii reale."
  } else if delta == 0 {
	  mesaj = "Ecuatia are 2 solutii reale si egale:"
	  mesaj_x1 = "x1 = x2 = " + fmt.Sprintf("%f", x1)
  } else if delta > 0 {
	  mesaj = "Ecuatia are doua solutii reale distincte:"
	  mesaj_x1 = "x1 = " + fmt.Sprintf("%f", x1)
	  mesaj_x2 = "x2 = " + fmt.Sprintf("%f", x2)
  }

  
  Titlu := "Ecuatie gradul 2 - rezultate"
  Variabile_in_html := Variabile{
    TitluPagina: Titlu,
    Mesaj_text : mesaj,
	Delta: mesaj_delta,  // afisare rezultate
	X1: mesaj_x1,
	X2: mesaj_x2,
	A:	fmt.Sprintf("%f", a), // pastram si afisam valorile introduse
	B:	fmt.Sprintf("%f", b),
	C:	fmt.Sprintf("%f", c),
    }

    t, err := template.ParseFiles("ecuatieGDR2.html")
    if err != nil { // daca sunt erori se vor afisa pe consola
      log.Print("eroare: ", err)
    }

    err = t.Execute(w, Variabile_in_html)
    if err != nil { // daca sunt erori se vor afisa pe consola
      log.Print("eroare la executie: ", err)
    }
}