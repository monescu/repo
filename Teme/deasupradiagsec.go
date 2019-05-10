package main

import "fmt"
//elementele de deasupra diagonalei secundare

func deasupraDiagSec(n int, matrice [][]int){
	var i = 0
	var j int

	for (i < n) {
		j = n - i - 1
		for (j < n) {
			fmt.Print(matrice[i][j], " ")
			j = j + 1
		}
		fmt.Println()
		i = i + 1
	}
}

func main(){

	matrice := [][]int{
		{1,  2,  3,  4},
		{5,  6,  7,  8},
		{9,  10,  11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println("Student Sraier Alina-Gabriela anul III ID")
	fmt.Println("exercitiu matrice-afisare elementele aflate deasupra diagonalei secundare")
	fmt.Println(matrice)

	deasupraDiagSec(4, matrice)

}