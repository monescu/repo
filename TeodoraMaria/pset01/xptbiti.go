package main

import "fmt"

func findX(x int, p uint, t uint) int {
	x = x >> t
	x = x & (1<<p - 1)

	return x
}

func main() {
	var result = findX(100, 4, 3)
	fmt.Println(result)
}
