package main
//elementele de deasupra diagonalei principale

import "fmt"


func deasupraDiagPrincipala(n int, matrice [][]int){
	var i = 0
	var j int

	for (i < n){
		j = i
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
	fmt.Println("exercitiu matrice-afisare elementele aflate deasupra diagonalei principale")
	fmt.Println(matrice)

	deasupraDiagPrincipala(4, matrice)

}
