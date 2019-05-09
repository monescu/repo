package main

import "fmt"

func BacktrackSudokuSolution(board [9][9]uint) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = uint(candidate)
					if isBoardValid(board) {
						if BacktrackSudokuSolution(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func main() {
	configuration := [9][9]uint{
		{6, 0, 9, 5, 7, 4, 1, 8, 2},
		{5, 4, 1, 0, 2, 9, 3, 7, 6},
		{7, 8, 2, 6, 1, 3, 9, 0, 4},
		{1, 9, 8, 4, 6, 0, 5, 2, 3},
		{3, 6, 5, 9, 8, 2, 4, 1, 7},
		{4, 2, 7, 1, 3, 5, 8, 6, 9},
		{9, 5, 6, 7, 4, 8, 2, 3, 1},
		{8, 1, 3, 0, 9, 6, 7, 4, 0},
		{2, 7, 4, 3, 5, 1, 6, 9, 8},
	}
	if BacktrackSudokuSolution(configuration) {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Printf("%d", uint64(configuration[i][j]))
				fmt.Print(" ")
			}
			fmt.Println()
		}
	}
}

func isBoardValid(board [9][9]uint) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}
func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func hasEmptyCell(board [9][9]uint) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}
