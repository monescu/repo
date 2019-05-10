package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func spiral(m int, n int, matrice [][]int)  {
	/*
		k - indexul primului rand
		m - indexul ultimului rand
		l - index primei coloane
		n - indexul ultimei coloane
		i - iterator
	*/
	var i int
	var k = 0
	var l = 0

	for (k < m && l < n) {
		i = l;
		/* Printeaza primul rand dintre randurile care au mai ramas */
		for (i < n) {
			fmt.Print(matrice[k][i], " ")
			i = i + 1
		}
		k = k + 1
		fmt.Println("")

		i = k
		/* Printeaza ultima coloana din coloanele ramase */
		for (i < m) {
			fmt.Print(matrice[i][n-1], " ");
			i = i + 1
		}
		n = n - 1
		fmt.Println("")

		/* Printeaza ultimul rand din randurile ramase */
		if ( k < m) {
			i = n - 1
			for (i >= l) {
				fmt.Print(matrice[m-1][i], " ");
				i = i - 1
			}
			m = m - 1;
			fmt.Println(" ")
		}

		/* Printeaza prima coloana din coloanele ramase */
		if (l < n) {
			i = m-1
			for (i >= k) {
				fmt.Print(matrice[i][l], " ")
				i = i - 1;
			}
			l = l + 1;
			fmt.Println(" ")
		}
	}


}

func main(){
	matrice := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println("Student Sraier Alina-Gabriela exercitiu afisarea unei matrici in spirala")
	fmt.Println(matrice)

	spiral(4,4, matrice)

}