package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func parseDate(date string) (time.Time, error) {
	return time.Parse("01/02/2006", date)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter date in the mm/dd/yyyy format: ")
	firstDateString, _ := reader.ReadString('\n')
	firstDateString = strings.Replace(firstDateString, "\n", "", -1)

	fmt.Print("Enter date in the mm/dd/yyyy format: ")
	secondDateString, _ := reader.ReadString('\n')
	secondDateString = strings.Replace(secondDateString, "\n", "", -1)

	firstDate, err := parseDate(firstDateString)
	secondDate, err := parseDate(secondDateString)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(firstDate)
	fmt.Println(secondDate)

	diff := secondDate.Sub(firstDate)
	fmt.Println(diff.Hours())
}
