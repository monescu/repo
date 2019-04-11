package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)


func enterValues(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("values.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		a,_:= strconv.ParseFloat(r.FormValue("a"),64)
		b,_:=strconv.ParseFloat(r.FormValue("b"),64)
		c,_:=strconv.ParseFloat(r.FormValue("c"),64)
		fmt.Fprintf(w,calculateSolutions(a,b,c))
		// logic part of log in
		//fmt.Fprintln(w, r.Form["a"])
		//fmt.Fprintln(w, r.Form["b"])
		//fmt.Fprintln(w, r.Form["c"])
	}
}


func calculateSolutions(a float64,b float64, c float64)string{

	var delta float64 = math.Pow(b, 2) - 4.0*a*c
	var x1 float64 = (-b + math.Sqrt(delta))/(2.0*a)
	var x2 float64 = (-b - math.Sqrt(delta))/(2.0*a)
	if(delta>0.0){
		return fmt.Sprintf("Radacinile sunt %.2f si %.2f",x1,x2)
	}

	if delta == 0.0{
		var result float64 = -b/(2.0*a)
		return fmt.Sprintf("x1=x2, adica radacina este %.2f",result)
	}

	if delta<0.0{
		return fmt.Sprintf("Ecuatia nu are solutii reale!!!")
	}
	return fmt.Sprintf("Radacinile sunt %.2f si %.2f",x1,x2)
}

func main() {

	http.HandleFunc("/", enterValues) // setting router rule
	//http.HandleFunc("/entervalues", enterValues)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	/*var a,b,c float64

	fmt.Println("Enter a:")
	fmt.Scanf("%f", &a)

	fmt.Println("Enter b:")
	fmt.Scanf("%f", &b)

	fmt.Println("Enter c:")
	fmt.Scanf("%f", &c)

	var delta float64 = math.Pow(b, 2) - 4.0*a*c

	if(delta>0.0){
		var x1 float64 = (-b + math.Sqrt(delta))/(2.0*a)
		var x2 float64 = (-b - math.Sqrt(delta))/(2.0*a)
		fmt.Printf("Radacinile sunt %.2f si %.2f",x1,x2)
	}

	if delta == 0.0{
		var result float64 = -b/(2.0*a)
		fmt.Printf("x1=x2, adica radacina este %.2f",result)
	}

	if delta<0.0{
		fmt.Println("Ecuatia nu are solutii reale!!!")
	}*/
}