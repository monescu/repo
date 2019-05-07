package main

import (
	"fmt"
)

func main(){
	var x,p,t uint
	fmt.Print("Introduceti X: ")
	fmt.Scan(&x)
	fmt.Print("Introduceti P: ")
	fmt.Scan(&p)
	fmt.Print("Introduceti T: ")
	fmt.Scan(&t)
	x = x >> t
	x = x & (1<<p - 1)
	fmt.Println("X-ul determinat:",x)
}