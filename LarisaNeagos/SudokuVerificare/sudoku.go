package main

import "fmt"

func Solve(sudoku [9][9]uint16) bool {

	a := [3][9]uint16{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

		position := sudoku[i][j]
		row := a[0][i]
		column := a[1][j]
		square := a[2][(3*(i/3))+j/3]

		if (row & (1 << position)) == (1 << position) || (column & (1 << position)) == (1 << position) || (square & (1 << position)) == (1 << position) {
			return false
		}
		  a[0][i] = row | (1 << position)
		  a[1][j] = column | (1 << position)
		  a[2][(3*(i/3))+j/3] = square | (1 << position)
		}
	}

	return true

}

func main() {
	sudoku := [9][9]uint16{
		{6, 8, 2, 1, 9, 4, 3, 5, 7},
		{7, 3, 1, 5, 6, 8, 9, 2, 4},
		{4, 9, 5, 7, 2, 3, 8, 6, 1},
		{8, 2, 7, 9, 3, 5, 1, 4, 6},
		{5, 1, 9, 6, 4, 7, 2, 8, 3},
		{3, 6, 4, 2, 8, 1, 5, 7, 9},
		{9, 5, 6, 4, 1, 2, 7, 3, 8},
		{2, 4, 8, 3, 7, 9, 6, 1, 5},
		{1, 7, 3, 8, 5, 6, 4, 9, 2},
}
	if(Solve(sudoku)) {
		fmt.Println("Sudoku corect")
	} else {
		fmt.Println("Sudoku incorect")
	}
}