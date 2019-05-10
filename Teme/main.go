package main

import "fmt"

const(
	message = "Salutare clasa a treia!"
	a = 3
)

var (
	message_var string = "Servus!"
	q int = 4
)



func cmmdc(a int, b int)int {
	for a%b != 0 {
		r := a % b
		a = b
		b = r
	}
	return b
}


const (
	first = 1 << (10 * iota)
	second
	third
)

const (
	xx = iota
	yy
	zz
)


func main() {
	//message = "Hahaha" //nu functioneaza daca s-a definit constanta message anterior

	//const message = "Hahaha" //functioneaza daca apare din nou message definita ca si constanta
	//println(message_var)
	//fmt.Println(message, q, a, message_var)

	//println(first)
	//println(second)
	//println(third)
	//println(1 << 20)

	println(cmmdc(42,52))
}

func init(){
fmt.Println(xx,yy,zz,"Salutare din init!")
}