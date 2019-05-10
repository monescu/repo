package main

import (
	"html/template"
	"net/http"
	"strconv"
	"math"
)

type Polinom struct {
	A   string
	B   string
	C   string

}
type Rezultat struct{
	A   string
	B   string
	C   string
	Success  bool
	IsReal bool
	IsComlex bool
	IsEcGr1 bool
	IsAltfel bool
	Sol1 string
	Sol2 string
	TipRezultat string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		dateCitite := Polinom{
			A:   r.FormValue("a"),
			B: 	 r.FormValue("b"),
			C:   r.FormValue("c"),
		}
		var rezEc Rezultat
		println(dateCitite.A,dateCitite.B,dateCitite.C)
		rezEc=ecGrad2(dateCitite)
		// do something with details
		//_ = details

		//tmpl.Execute(w, struct{ Success bool  }{true})
		tmpl.Execute(w, rezEc)
	})

	http.ListenAndServe(":8080", nil)
}

func ecGrad2(datePol Polinom ) Rezultat{
	//initializare rezultat cu datele polinomului
	var rezEc Rezultat
	rezEc.A=datePol.A
	rezEc.B=datePol.B
	rezEc.C=datePol.C
	rezEc.Success=true

	var a,b,c, delta,sol1,sol2,re,im float64
	//converirea la float64 a coeficientiilor polinomului
	a, _ =strconv.ParseFloat(datePol.A,64)
	b, _ =strconv.ParseFloat(datePol.B,64)
	c, _ =strconv.ParseFloat(datePol.C,64)

	// rezolvarea ecuatiei

	if a!=0{
		delta=math.Pow(b,2)-4*a*c
		if delta>0{
			sol1=(-b+math.Sqrt(delta))/(2*a)
			sol2=(-b-math.Sqrt(delta))/(2*a)
			//rezultate
			rezEc.IsReal=true
			rezEc.TipRezultat="Solutile sunt reale"
			rezEc.Sol1=strconv.FormatFloat(sol1, 'f', 6, 64)
			rezEc.Sol2=strconv.FormatFloat(sol2, 'f', 6, 64)
		}else if delta==0{
			sol1=(-b)/(2*a)
			sol2=sol1
			//rezultate
			rezEc.IsReal=true
			rezEc.TipRezultat="Solutile sunt egale si reale"
			rezEc.Sol1=strconv.FormatFloat(sol1, 'f', 6, 64)
			rezEc.Sol2=strconv.FormatFloat(sol2, 'f', 6, 64)
		}else{
			re=b*(2*a)
			im=math.Sqrt(-delta)/(2*a)
			//rezultate
			rezEc.IsComlex=true
			rezEc.TipRezultat="Este o solutie complexa"
			rezEc.Sol1=strconv.FormatFloat(re, 'f', 6, 64)
			rezEc.Sol2=strconv.FormatFloat(im, 'f', 6, 64)
		}
	}else if b!=0{
		sol1=-c/b
		//rezultate
		rezEc.IsEcGr1=true
		rezEc.TipRezultat="Este o solutie"
		rezEc.Sol1=strconv.FormatFloat(sol1, 'f', 6, 64)
	}else if(c!=0){
		//nu sunt solutii
		rezEc.IsAltfel=true
		rezEc.TipRezultat="Nu sunt solutii"
	}else{
		//o solutie infinita
		rezEc.IsAltfel=true
		rezEc.TipRezultat="Sunt o infinitate de soluti "
	}


	return rezEc
}