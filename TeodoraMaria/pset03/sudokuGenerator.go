package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("E:/facultate/go/repo/homework/pset03/index.html"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/sudoku", processSudokuBoasrd)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func processSudokuBoasrd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	initSudoku := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	var sudoku *[9][9]int
	sudoku = &initSudoku

	readBoard(r, sudoku)
	if isGeneratedValidConfig(sudoku, 0) {
		display(sudoku, w)
	}
}

func readBoard(r *http.Request, sudoku *[9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cellName := fmt.Sprintf("%s%d%d", "cell-", i, j)
			sudoku[i][j], _ = strconv.Atoi(r.FormValue(cellName))
		}
	}
}

func absentOnLine(k int, sudoku [9][9]int, x int) bool {
	var y int
	for y = 0; y < 9; y++ {
		if sudoku[x][y] == k {
			return false
		}
	}
	return true
}

func absentOnRow(k int, sudoku [9][9]int, y int) bool {
	var x int
	for x = 0; x < 9; x++ {
		if sudoku[x][y] == k {
			return false
		}
	}
	return true
}

func absentOnBloc(k int, sudoku [9][9]int, x int, y int) bool {
	var firstX, firstY int
	firstX = x - (x % 3)
	firstY = y - (y % 3)
	for x = firstX; x < firstX+3; x++ {
		for y = firstY; y < firstY+3; y++ {
			if sudoku[x][y] == k {
				return false
			}
		}
	}
	return true
}

func isGeneratedValidConfig(sudoku *[9][9]int, position int) bool {
	if position == 9*9 {
		return true
	}
	var x, y, k int
	x = position / 9
	y = position % 9
	if sudoku[x][y] != 0 {
		return isGeneratedValidConfig(sudoku, position+1)
	}
	for k = 1; k <= 9; k++ {
		if absentOnLine(k, *sudoku, x) && absentOnRow(k, *sudoku, y) && absentOnBloc(k, *sudoku, x, y) {
			sudoku[x][y] = k
			if isGeneratedValidConfig(sudoku, position+1) {
				return true
			}
		}
	}
	sudoku[x][y] = 0
	return false
}

func display(sudoku *[9][9]int, w http.ResponseWriter) {
	fmt.Fprintln(w, "Sudoku Solution: ")
	fmt.Fprintln(w)
	for i := 0; i < 9; i++ {
		if i == 3 || i == 6 {
			fmt.Fprintln(w, "---------------------")
		}
		for j := 0; j < 9; j++ {
			if j == 3 || j == 6 {
				fmt.Fprint(w, "| ")
			}
			fmt.Fprint(w, sudoku[i][j])
			fmt.Fprint(w, " ")
		}
		fmt.Fprintln(w)
	}
}
