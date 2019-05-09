package main

import "fmt"

func main() {

	sudoku := [9][9]uint{
		{4, 3, 5, 2, 6, 9, 7, 8, 1},
		{6, 8, 2, 5, 7, 1, 4, 9, 3},
		{1, 9, 7, 8, 3, 4, 5, 6, 2},
		{8, 2, 6, 1, 9, 5, 3, 4, 7},
		{3, 7, 4, 6, 8, 2, 9, 1, 5},
		{9, 5, 1, 7, 4, 3, 6, 2, 8},
		{5, 1, 9, 3, 2, 6, 8, 7, 4},
		{2, 4, 8, 9, 5, 7, 1, 3, 6},
		{7, 6, 3, 4, 1, 8, 2, 5, 9},
	}

	if check(sudoku) {
		fmt.Println("Rezolvare corecta!")
	}
}

func check(sudoku [9][9]uint) bool {
	var a = [3][9]uint{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row := a[0][i]
			column := a[1][j]
			square := a[2][3*(i/3)+(j/3)]
			position := sudoku[i][j]
			if (row & (1 << position)) == (1 << position) {
				fmt.Println("Greseala pe randul ", i)
				return false
			} else {
				a[0][i] = row | (1 << position)
			}
			if (column & (1 << position)) == (1 << position) {
				fmt.Println("Greseala pe coloana ", j)
				return false
			} else {
				a[1][j] = column | (1 << position)
			}
			if (square & (1 << position)) == (1 << position) {
				fmt.Println("Greseala in patrat")
				return false
			} else {
				a[2][3*(i/3)+(j/3)] = square | (1 << position)
			}
		}
	}
	return true
}