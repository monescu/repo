package homework

import (
	"fmt"
	"time"
)

//PUT THIS IN MAIN
// date := time.Now()

// otherDate := time.Date(2090, time.September, 10, 23, 0, 0, 0, time.UTC)

// homework.DateDifferenceInDays(date, otherDate)
// homework.DateDifferenceFull(date, otherDate)

// fmt.Println(date.String())
// fmt.Println(homework.AddDays(date, 10).String())
// fmt.Println(homework.RemoveDays(date, -10).String())

// homework.PrintWeekDay(date)

//MyDate return a date
func MyDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

//DateDifferenceInDays returns the difference in days between 2 dates
func DateDifferenceInDays(date time.Time, otherDate time.Time) {
	days := int(otherDate.Sub(date).Hours() / 24)
	fmt.Println("The difference in days is", days)
}

//DateDifferenceFull will print years, months, days, minutes, seconds
func DateDifferenceFull(date time.Time, otherDate time.Time) {
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

//AddDays adds days to a date
func AddDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

//RemoveDays substracts days to a date
func RemoveDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, -days)
}

//PrintWeekDay substracts days to a date
func PrintWeekDay(date time.Time) {
	fmt.Println(date.Weekday().String())
}
