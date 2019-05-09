package main

import "fmt"

func main() {
   /* an array with 5 rows and 2 columns*/
   var matrix = [5][5]int{ {1,  2,  3,  4,  5} ,{6,  7,  8,  9,  10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
   var i, j, k, l, m, n int

   /* output each array element's value */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 5; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, matrix[i][j] )
      }
   }
   fmt.Printf("matrix length = %d\n",len(matrix))
   // Spiral iteration
   var rowStart int = 0
   var rowLength int = len(matrix) - 1
   
   var colStart int = 0
   var colLength int = len(matrix) - 1
   
   /*for rowStart <= rowLength && colStart <= colLength{
	
   }*/
   
   for ok := true; ok; ok = rowStart <= rowLength && colStart <= colLength {
    for k = rowStart; k <= colLength; k++{
		fmt.Printf("a[%d][%d] = %d\n", rowStart,k, matrix[rowStart][k] )
	}
	for l = rowStart + 1; l <= rowLength; l++{
		fmt.Printf("a[%d][%d] = %d\n", l,colLength, matrix[l][colLength] )
	}
	if (rowStart + 1 <= rowLength){
		for m = colLength - 1; m >= colStart; m--{
			fmt.Printf("a[%d][%d] = %d\n", rowLength,m, matrix[rowLength][m] )
		}
	}
	if (colStart + 1 <= colLength){
		for n = rowLength - 1; n > rowStart; n--{
			fmt.Printf("a[%d][%d] = %d\n", n,colStart, matrix[n][colStart] )
		}
	}
   rowStart++
   rowLength--
   colStart++
   colLength--
}
   
}