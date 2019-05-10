package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Println("Hello!")
		fmt.Println("1. Difference")
		fmt.Println("2. Add a date + number of days")
		fmt.Println("3. Add a date - number of days")
		fmt.Println("4. Get the day of the week from a date")
		fmt.Println("5. Exit")

		option, err := readMenuOption(reader, "Option: ")
		fmt.Println(option)
		if err != nil {
			fmt.Println(err)
		} else {
			if option <= 0 && option >= 6 {
				fmt.Println("Invalid option, please try again!")
				option = -1
			} else {
				switch option {
				case 1:
					getDifferenceBetweenDates(reader)
				case 2:
					addDateForOperation(reader, "add")
				case 3:
					addDateForOperation(reader, "substract")
				case 4:
					addDateForOperation(reader, "getWeekday")
				case 5:
					os.Exit(0)
				}
			}
		}
	}
}

func diffMonths(startDate, endDate time.Time) int {
	var months = 0
	for startDate.Before(endDate) {
		startDate = startDate.AddDate(0, 1, 0)
		if startDate.Year() != endDate.Year() {
			months++
		} else {
			if startDate.Month() != endDate.Month() {
				months++
			}
		}
	}
	return months
}

func checkDate(startDate, endDate time.Time) (bool, error) {
	if startDate.After(endDate) {
		return false, fmt.Errorf("Final date bigger than the start date!")
	}
	return true, nil
}

func addDateForOperation(reader *bufio.Reader, operation string) {

	fmt.Println("Date format: YYYY-MM-DD")
	var auxDate time.Time

	date, err := readAndValidateDate(reader, "Date")
	if err == nil {
		if operation == "getWeekday" {
			fmt.Println("The weekday is: ", date.Weekday())
		} else {
			numberDays, err := readMenuOption(reader, "Number of dys to"+operation+":")
			if err != nil {
				fmt.Println(err)
			} else {
				switch operation {
				case "add":
					auxDate = date.AddDate(0, 0, numberDays)
				case "substract":
					auxDate = date.AddDate(0, 0, -numberDays)
				}
				fmt.Println("Result date: ", auxDate)
			}
		}
	} else {
		fmt.Println(err)
	}

}

func readAndValidateDate(reader *bufio.Reader, message string) (time.Time, error) {
	fmt.Print(message)
	date, _ := reader.ReadString('\n')
	date = strings.Replace(date, "\n", "", -1)
	tdate, errorOccurred := validateDate(date)

	return tdate, errorOccurred
}

func getDifferenceBetweenDates(reader *bufio.Reader) {

	fmt.Println("Date format: YYYY-MM-DD")
	startDate, startDateErr := readAndValidateDate(reader, "Start date")
	endDate, endDateErr := readAndValidateDate(reader, "End date:")

	if startDateErr != nil {
		fmt.Println(startDateErr)
	}

	if endDateErr != nil {
		fmt.Println(endDateErr)
	}

	if startDateErr == nil && endDateErr == nil {
		res, err := checkDate(startDate, endDate)
		if res == true {
			var diff = endDate.Sub(startDate)
			var days = int(diff.Hours() / 24)
			fmt.Println("Difference in days:  days", days, "\n")

			var weeks = int(diff.Hours() / (24 * 7))
			fmt.Printf("Diffrence in weeks : %d weeks\n", weeks)

			fmt.Printf("Diffrence in monts : %d monts\n", diffMonths(startDate, endDate))
			if startDate.Month() > endDate.Month() {
				fmt.Printf("Diffrence in years : %d years\n", endDate.Year()-startDate.Year()-1)
			} else {
				fmt.Printf("Diffrence in years : %d years\n", endDate.Year()-startDate.Year())
			}
		} else {
			fmt.Println(err)
		}
	}
}

func readMenuOption(reader *bufio.Reader, message string) (int, error) {
	fmt.Print(message)
	stringNumber, _ := reader.ReadString('\n')
	stringNumber = strings.Replace(stringNumber, "\n", "", -1)
	number, err := strconv.Atoi(stringNumber)
	return number, err
}

func validateDate(date string) (time.Time, error) {
	tdate, err := transformStringToDate(date)
	if err != nil {
		return tdate, err
	}
	return tdate, err
}

func transformStringToDate(date string) (time.Time, error) {
	const template = "2006-01-02"
	t, err := time.Parse(template, date)
	if err != nil {
		return t, err
	}
	return t, err
}
