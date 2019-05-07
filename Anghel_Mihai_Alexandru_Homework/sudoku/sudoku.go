package main

import (
	"fmt"
	"os"
)

func checkSudoku(S [9][9]uint){
	var a[3][9] uint

	for i := 0; i<9; i++ {
		for j := 0; j<9; j++ {
			p := S[i][j] - 1
			k := 3 * (i/3) + (j/3)
			if a[0][i] & (1 << p) == 0 {
				a[0][i] = a[0][i] | (1 << p)
			} else {
				fmt.Println("Solutie incorecta!")
				os.Exit(0)
			}
			if a[1][j] & (1 << p) == 0 {
				a[1][j] = a[1][j] | (1 << p)
			} else {
				fmt.Println("Solutie incorecta!")
				os.Exit(0)
			}
			if a[2][k] & (1 << p) == 0 {
				a[2][k] = a[2][k] | (1 << p)
			} else {
				fmt.Println("Solutie incorecta!")
				os.Exit(0)
			}
		}
	}
	fmt.Println("Solutie corecta!")
}

func main() {
	var S = [9][9] uint { {1,2,3,4,5,6,7,8,9}, {7,8,9,1,2,3,4,5,6}, {4,5,6,7,8,9,1,2,3}, {3,1,2,8,4,5,9,6,7}, {6,9,7,3,1,2,8,4,5}, {8,4,5,6,9,7,3,1,2}, {2,3,1,5,7,4,6,9,8}, {9,6,8,2,3,1,5,7,4}, {5,7,4,9,6,8,2,3,1} }
	checkSudoku(S)
}