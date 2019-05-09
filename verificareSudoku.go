package main


func main() {

	var sudoku =[][] uint16{
		{1,2,3,4,5,6,7,8,9},
		{2,3,4,5,6,7,8,9,1},
		{3,4,5,6,7,8,9,1,2},
		{4,5,6,7,8,9,1,2,3},
		{5,6,7,8,9,1,2,3,4},
		{6,7,8,9,1,2,3,4,5},
		{7,8,9,1,2,3,4,5,6},
		{8,9,1,2,3,4,5,6,7},
		{9,1,2,3,4,5,6,7,8},
	}
	println(check_sudoku_biti(sudoku))


}


func check_sudoku_biti(sudoku [][]uint16) uint8{
	var a =[3][9] uint16{}
	var p uint16
	var k uint16
	for i := 0; i<9;i++ {
		for j := 0; j < 9; j++ {
			p = sudoku[i][j] - 1
			k = uint16(i/3*3 + j/3)
			if a[0][i]&(1<<p) == 0 {
				a[0][i] = a[0][i] | (1 << p)
			} else {
				return 0
			}
			if a[1][j]& (1<<p) == 0{
				a[1][j] = a[1][j] | (1<<p)
			}else{
				return 0
			}
			if a[2][k]& (1<<p) == 0{
				a[2][j] = a[2][k]
			}else {
				return 0
			}
		}

	}
	return 1
}
