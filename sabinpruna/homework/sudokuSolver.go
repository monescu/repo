package homework

import (
	"fmt"
)

//PrintSudokuBoard returns a string that displays sudoku solution nicely
func PrintSudokuBoard(board [9][9]int) string {
	var sudokuSolution string
	sudokuSolution = fmt.Sprintln("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		sudokuSolution += fmt.Sprint("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				sudokuSolution += fmt.Sprint("| ")
			}
			sudokuSolution += fmt.Sprintf("%d ", board[row][col])
			if col == 8 {
				sudokuSolution += fmt.Sprint("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			sudokuSolution += fmt.Sprintln("\n+-------+-------+-------+")
		} else {
			sudokuSolution += fmt.Sprintln()
		}
	}
	return sudokuSolution
}

//BacktrackSudokuSolution find sudoku solution if it exists
//@nticaric
func BacktrackSudokuSolution(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
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

//----------------------------PRIVATES------------------------------

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func arrayToSlice(board [9][9]int) []uint16 {
	var slice []int

	temp := board[:]

	for _, arr := range temp {
		slice = append(slice, arr[:]...)
	}

	var uintSlice []uint16
	for _, elem := range slice {
		uintSlice = append(uintSlice, uint16(elem))
	}

	return uintSlice
}

func isBoardValid(board *[9][9]int) bool {

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

//------------------------------------------------------------------
