package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("sudoku.html"))
}

func main() {
	http.HandleFunc("/", render)
	http.HandleFunc("/process", handle)
	http.ListenAndServe(":9090", nil)

}

func render(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "sudoku.html", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var sudoku [9][9] int
	var i,j int
	for i = 0; i<9; i++ {
		for j = 0; j<9; j++ {
			strI := strconv.Itoa(i)
			strJ := strconv.Itoa(j)
			var cell,_ = strconv.Atoi(r.FormValue("cell-" + strI + strJ))
			 if cell != 0 {
			 	sudoku[i][j] = cell
			 }
		}
	}

	IsValid(&sudoku, 0)
	Display(sudoku,w)
}
func Display(sudoku [9][9] int,w http.ResponseWriter,) {

	fmt.Fprintln(w,"Rezolvarea este: ")
	var x, y int
	for x = 0; x < 9; x++ {
		fmt.Fprintln(w,"")
		if (x == 3 || x == 6) {
			fmt.Fprintln(w," ")
		}
		for y = 0; y < 9; y++ {
			if (y == 3 || y == 6) {
				fmt.Fprint(w,"|")
			}
			fmt.Fprint(w,sudoku[x][y])
		}
	}
}
func AbsentOnLine(k int, sudoku [9][9]int, x int) bool {
	var y int
	for y = 0; y < 9; y++ {
		if (sudoku[x][y] == k) {
			return false
		}
	}
	return true
}
func AbsentOnRow(k int, sudoku [9][9]int, y int) bool {
	var x int
	for x = 0; x < 9; x++ {
		if (sudoku[x][y] == k) {
			return false;
		}
	}
	return true;
}
func AbsentOnBloc(k int, sudoku [9][9]int, x int, y int) bool {
	var firstX, firstY int;
	firstX = x - (x % 3)
	firstY = y - (y % 3)
	for x = firstX; x < firstX+3; x++ {
		for y = firstY; y < firstY+3; y++ {
			if (sudoku[x][y] == k) {
				return false;
			}
		}
	}
	return true;
}
func IsValid(sudoku *[9][9]int, position int) bool {
	if (position == 9*9) {
		return true;
	}
	var x, y, k int
	x = position / 9
	y = position % 9
	if (sudoku[x][y] != 0) {
		return IsValid(sudoku, position+1);
	}
	for k = 1; k <= 9; k++ {
		if (AbsentOnLine(k, *sudoku, x) && AbsentOnRow(k, *sudoku, y) && AbsentOnBloc(k, *sudoku, x, y)) {
			sudoku[x][y] = k;
			if (IsValid(sudoku, position+1)) {
				return true;
			}
		}
	}
	sudoku[x][y] = 0;
	return false;
}
