package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	number, err := validare(reader,"x= ")
	if err != nil{
		fmt.Println(err)
	} else {
		p, errP := validare(reader,"p= ")
		if errP != nil{
			fmt.Println(errP)
		} else {
			t, errT := validare(reader,"t= ")
			if errT != nil{
				fmt.Println(errT)
			} else {
				x := number >> t
				x = x & (1<<p - 1)
				fmt.Println(x)
			}
		}
	}
}

func validare(reader *bufio.Reader, message string) (uint64, error) {
	fmt.Print(message)
	numberStr, _ := reader.ReadString('\n')
	numberStr = strings.Replace(numberStr, "\n", "", -1)
	number, err := strconv.ParseUint(numberStr, 10, 64)

	return number, err
}
