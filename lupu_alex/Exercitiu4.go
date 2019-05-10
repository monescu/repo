package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Date struct {
	year int
	month int
	day int
}

var months = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var reader = bufio.NewReader(os.Stdin)
var weekDays = [...]string{"Duminica", "Luni", "Marti", "Miercuri", "Joi", "Vineri", "Sambata"}

func main() {
	rgx, _ := regexp.Compile("([12]\\d{3}\\.(0[1-9]|1[0-2])\\.(0[1-9]|[12]\\d|3[01]))")

	fmt.Println("Please provide the first date(yyyy.mm.dd):")
	firstDate, _ := reader.ReadString('\n')

	if(!rgx.MatchString(firstDate)) {
		fmt.Println("Data nu este valida")
		return
	}

	fmt.Println("Please provide the second date(yyyy.mm.dd):")
	secondDate, _ := reader.ReadString('\n')

	if(!rgx.MatchString(secondDate)) {
		fmt.Println("Data nu este valida")
		return
	}

	fDate := parseDate(firstDate)
	sDate := parseDate(secondDate)

	for {
		clearWindowsScreen()
		fmt.Println("Selectati o optiune:")
		fmt.Println("1. Diferenta dintre date in zile")
		fmt.Println("2. Diferenta dintre date in ani, luni, zile")
		fmt.Println("3. Adauga zile la prima data")
		fmt.Println("4. Scade zile din prima data")
		fmt.Println("5. Ziua saptamanii pentru prima data")
		fmt.Println("0. Iesire")

		var option int
		fmt.Scanf("%d", &option)

		switch option {
		case 0:
			return;
		case 1:
			days := getDiffInDays(fDate, sDate)
			fmt.Println("Diferenta in zile este ", days)
		case 2:
			years, months, days := getDiffInYMD(fDate, sDate)
			fmt.Println("Diferenta in ani, luni si zile este ", years, " ", months, " ", days)
		case 3:
			days := getDaysInDate(fDate)
			var cnt int
			fmt.Scanf("%d", &cnt)
			days = days + cnt
			date := getDateFromDays(days)
			fmt.Println("Noua data: ", date.year, ".", date.month, ".", date.day)
		case 4:
			days := getDaysInDate(fDate)
			var cnt int
			fmt.Scanf("%d", &cnt)
			days = days - cnt
			date := getDateFromDays(days)
			fmt.Println("Noua data: ", date.year, ".", date.month, ".", date.day)
		case 5:
			days := getDaysInDate(fDate)
			dayOweek := getDayOfWeek(days)
			fmt.Println("Ziua este :", dayOweek)
		default:
		}
	}

}

func getDiffInYMD(firstDate, secondDate Date) (int, int, int) {
	totalDays := getDiffInDays(firstDate, secondDate)
	years :=  totalDays/365
	months := 0
	days := 0

	totalDays = totalDays%365

	if(totalDays > 0){
		months = totalDays/30
		days = totalDays%30
	}
	return years, months, days
}

func clearWindowsScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func parseDate(line string) Date {

	str := strings.Split(line, ".")
	year, _ := strconv.Atoi(str[0])
	month, _ := strconv.Atoi(str[1])
	day, _ := strconv.Atoi(strings.ReplaceAll(str[2], "\n", "" ))

	return Date{ year, month, day}
}

func getDiffInDays(firstDate, secondDate Date) int {
	diff := getDaysInDate(firstDate) - getDaysInDate(secondDate)

	if(diff < 0) {
		return -diff;
	}

	return diff;
}

func compareDate(firstDate, secondDate Date) int {
	diff := getDaysInDate(firstDate) - getDaysInDate(secondDate)

	if(diff < 0) {
		return -1
	} else if diff > 0 {
		return 1
	} else {
		return 0
	}
}

func getDaysInDate(date Date) int {
	return getDaysTillYear(date.year) + getDaysTillMonth(date.month, isLeapYear(date.year)) + date.day
}

func getDaysTillYear(year int) int {
	year--
	leapYears := float64(year)/4 - float64(year)/100 + float64(year)/400
	totalDays := int(leapYears) + year*365;
	return totalDays
}

func getDaysTillMonth(month int, leapYear bool) int {
	month--
	days:=0
	for i:=0;i<month ;i++  {
		days += months[i]
	}
	if(leapYear) {
		days++
	}
	return days;
}

func isLeapYear(year int) bool {
	return (year%100 != 0 || year%400 == 0) && year%4 == 0
}

func getDateFromDays(days int) Date {
	year := days/365
	days = days%365 + 365
	leapYears := float64(year)/4 - float64(year)/100 + float64(year)/400
	days = days - int(leapYears)
	if(days > 365) {
		days -= 365
		year++
	}
	month := 1
	for i:=0;i <12;i++ {
		if days > months[i] {
			month++
			days = days - months[i]
		} else {
			break;
		}
	}

	return Date{year, month, days}
}

func getDayOfWeek(days int) string {
	return weekDays[days%7]
}