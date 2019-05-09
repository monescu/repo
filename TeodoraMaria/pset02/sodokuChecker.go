package main

import "fmt"

func main() {
	configuration := [9][9]uint{
		{6, 3, 9, 5, 7, 4, 1, 8, 2},
		{5, 4, 1, 8, 2, 9, 3, 7, 6},
		{7, 8, 2, 6, 1, 3, 9, 5, 4},
		{1, 9, 8, 4, 6, 7, 5, 2, 3},
		{3, 6, 5, 9, 8, 2, 4, 1, 7},
		{4, 2, 7, 1, 3, 5, 8, 6, 9},
		{9, 5, 6, 7, 4, 8, 2, 3, 1},
		{8, 1, 3, 2, 9, 6, 7, 4, 5},
		{2, 7, 4, 3, 5, 1, 6, 9, 8},
	}
	if isValidConfiguration(configuration) {
		fmt.Println("Given configuration is valid!")
	} else {
		fmt.Println("Given configuration is invalid!")
	}
}

func isValidConfiguration(configuration [9][9]uint) bool {
	a := [3][9]uint{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			position := configuration[i][j]
			line := a[0][i]
			column := a[1][j]
			square := a[2][(3*(i/3))+j/3]
			var bitPosition uint = 1 << position
			if (line&bitPosition) == bitPosition || (column&bitPosition) == bitPosition || (square&bitPosition) == bitPosition {
				return false
			}
			a[0][i] = line | bitPosition
			a[1][j] = column | bitPosition
			a[2][(3*(i/3))+j/3] = square | bitPosition
		}
	}
	return true
}
