package main

import "fmt"

var msg string ="Salut"
var x int

/*var(
	a int
	b string
	c int
)
*/
const(
	message = "Salutare clasa a treia!"
	a = 3
)


func ceva(x float64, y float64) (r1 float64,r2 float64){
	r1 = x+y
	r2 = x-y
	return
}

func cmmdc(a int, b int)int {
	for a%b != 0 {
		r := a % b
		a = b
		b = r
	}
	return b
}
/*
const (
	xx = iota
	yy = 5
	zz
	tt
)


const (
	first = 1 << (10 * iota)
	second
	third
)
*/

func main(){
	//a= 5
	//a,b := ceva(4,5)
   //println(a,b)
	//println(ceva(5,6))
   //fmt.Println(ceva(2,3))

   x:=[][]int{{2,6,8,2,1,6,9,0},{4,5,4}}



   fmt.Println(x)
}

