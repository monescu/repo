package main

import "fmt"




func main() {
	var a = [3][4]int{
		{0, 1, 2, 3} ,
		{4, 5, 6, 7} ,
		{8, 9, 10, 11}}


	var i, k, l, m, n int
	//k - starting row index
	//m - ending row index
	m = 3
	//l - starting column index
	//n - ending column index
	n = 4
	//i - iterator


	fmt.Println("Matricea afisata in spirala: ")
	for k < m && l < n {
		// Print the first row from the remaining rows
		for i = l; i < n; i++ {
			fmt.Printf("a[%d][%d] = %d\n", k,i, a[k][i])
		}
		k++

		// Print the last column from the remaining columns
		for i = k; i < m; i++ {
			fmt.Printf("a[%d][%d] = %d\n", i,n-1, a[i][n-1])
		}
		n--

		// Print the last row from the remaining rows */
		if k < m {
			for i = n-1; i >= l; i-- {
				fmt.Printf("a[%d][%d] = %d\n", m-1,i, a[m-1][i])
			}
			m--
		}

		// Print the first column from the remaining columns */
		if l < n{
			for i = m-1; i >= k; i--{
				fmt.Printf("a[%d][%d] = %d\n", i, l, a[i][l])
			}
			l++;
		}
	}


}