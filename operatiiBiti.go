package main

import "fmt"

func main() {
	fmt.Println("The number is", FindNumber(33,4,2))
}

func FindNumber(x int, p uint, t uint) int {
	x = x >> t
	x = x & (1<<p - 1)

	return x
}