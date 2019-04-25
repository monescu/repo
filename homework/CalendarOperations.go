package main

import (
	"fmt"
	"time"
)

func main() {
	currentDate := time.Now()

	otherDate := time.Date(2030, time.December, 9, 5, 0, 0, 0, time.UTC)

	DateDifferenceInDays(currentDate, otherDate)
	DateDifferenceComplete(currentDate, otherDate)

	fmt.Println("The current date is", currentDate.String())
	fmt.Println("The sum between a date and a number of days is", AddDays(currentDate, 10).String())
	fmt.Println("The difference between a date and a number of days is", RemoveDays(currentDate, -10).String())

	PrintWeekDay(currentDate)
}

func MyDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func DateDifferenceInDays(date time.Time, otherDate time.Time) {
	days := int(otherDate.Sub(date).Hours() / 24)
	fmt.Println("The difference in days is", days)
}

func DateDifferenceComplete(date time.Time, otherDate time.Time) {
	duration := otherDate.Sub(date)

	years := int(duration.Seconds() / 31207680)

	months := int(duration.Seconds()/2600640) % 12
	weeks := (int(duration.Seconds()/604800) % 12) % 4
	days := ((int(duration.Seconds()/86400) % 12) % 4) % 7

	fmt.Println("Years : ", years)
	fmt.Println("Months : ", months)
	fmt.Println("Weeks : ", weeks)
	fmt.Println("Days : ", days)
}

func AddDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

func RemoveDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, -days)
}

func PrintWeekDay(date time.Time) {
	fmt.Println("The current week day is",date.Weekday().String())
}
