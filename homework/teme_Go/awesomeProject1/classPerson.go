package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type T struct{
	f_1, F_2 int
}

func (p *Person) changeAge(newAge int){
	p.Age = newAge
}

func main() {

	var t T
	t.f_1 = 3


	//var p Person
	var p1 = Person{"Vlad", "Monescu", 12}

	var p2 = Person{
		FirstName: "Dan",
		LastName:  "Popescu",
		Age:       45,
	}



	pointer1 := &p1
	pointer2 := &p2
	pointer3 := new(Person)

	pointer3.Age = 8

	fmt.Println(pointer1,*pointer2, *pointer3)

	p1.changeAge(4)
	fmt.Println(p1)

}