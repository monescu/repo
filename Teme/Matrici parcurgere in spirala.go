package main

import "fmt"


func spirala(n int) []int {
	stanga, sus, dreapta, jos := 0, 0, n-1, n-1
	dim := n * n
	s := make([]int, dim)
	i := 0
	for stanga < dreapta {
		for c := stanga; c <= dreapta; c++ {
			s[sus*n+c] = i
			i++
		}
		sus++
		for r := sus; r <= jos; r++ {
			s[r*n+dreapta] = i
			i++
		}
		dreapta--
		if sus== jos {
			break
		}

		for c := dreapta; c >= stanga; c-- {
			s[jos*n+c] = i
			i++
		}
		jos--

		for r := jos; r >= sus; r-- {
			s[r*n+stanga] = i
			i++
		}
		stanga++
	}
	//centru
	s[sus*n+stanga] = i

	return s
}

func main() {
	num := 5
	len := 2
	for i, draw := range spirala(num) {
		fmt.Printf("%*d ", len, draw)
		if i%num == num-1 {
			fmt.Println("")
		}
	}
}