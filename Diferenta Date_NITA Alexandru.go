package main

import(
	"fmt"
	"time"
)

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func main() {

	firstDate := time.Date(2019, 4, 9, 0, 0, 0, 0, time.UTC)
	fmt.Println("Data curenta: ", firstDate.Format(time.RFC3339))

	secondDate := time.Date(2018, 8, 26, 5, 10, 32, 235899241, time.UTC)
	fmt.Println("Data aleasa: ", secondDate)

	diferenta := firstDate.Sub(secondDate)
	zile := int16 (diferenta.Hours()/24)
	fmt.Println("Diferenta intre doua date in zile: ", zile)

	year, month, day, hour, min, sec := diff(secondDate, firstDate)
	fmt.Printf("Diferenta %d ani, %d luni, %d zile, %d ore, %d minute si %d secunde.", year, month, day, hour, min, sec)
	fmt.Println(" ")

	fmt.Println("Ziua saptamanii a datei alese: " ,secondDate.Weekday())

	daysToAdd := 6
	daysToSubstract := 27
	fmt.Println("Data plus numar de zile: ", secondDate.AddDate(0, 0, daysToAdd))
	fmt.Println("Data minus numar de zile: ", secondDate.AddDate(0, 0, (-daysToSubstract)))

}