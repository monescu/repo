package main

import (
	"fmt"
)

func main() {
	var x, p, t uint

	x = 250
	p = 4
	t = 3

	fmt.Println( (x>>t) & (1<<p-1))
}
