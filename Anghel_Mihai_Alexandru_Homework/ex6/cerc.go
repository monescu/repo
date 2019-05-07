package main

import "fmt"

func main(){
	var a[100][100] int
	var row,col int
	fmt.Print("Introduceti numarul de linii: ")
	fmt.Scanln(&row)
	fmt.Print("Introduceti numarul de coloane: ")
	fmt.Scanln(&col)

	fmt.Println()
	fmt.Println("Afisare matrice")
	fmt.Println()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Introduceti elementele matricei %d %d :",i+1,j+1)
			fmt.Scanln(&a[i][j])

		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ",a[i][j])
			if(j==col-1){
				fmt.Println("")
			}
		}
	}

	var min=999999
	var max[100] int
	max[0]=0
	var poz int=0

	for i:=0;i<row;i++ {
		for j:=0;j<col;j++ {
			if max[poz]<a[i][j] {
				max[poz]=a[i][j]
			}
		}
		poz++
	}
	for i:=0;i<row;i++{
		if min>max[i]{
			min=max[i]
		}
	}
	fmt.Printf("Minimul dintre maximele de pe linii este: %d ",min)
}