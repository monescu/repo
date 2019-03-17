package main

import (
	"fmt"
	"os"
	//"math/cmplx"

)

var (

	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}


func test(x, y, z int)(int, int, int){

	return x + y, x + z, y + z
}


func main() {

	var v1, v3 int
	v1, _, v3=test(1,2,3)
	fmt.Println(v1, v3)

	//fmt.Println(a,b,c,d)

	//fmt.Printf("%X", 7768)

	//x:=0

	//x,_=fmt.Scanf("%d", &x)


	//fmt.Println(cmplx.Sqrt(-4.))
}
