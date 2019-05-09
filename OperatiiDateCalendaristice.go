package main

import "fmt"
import "time"

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
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

	now := time.Now()
	fmt.Println("Data curenta: ", now)

	choseDate := time.Date(
		2019, 3, 5, 10, 34, 58, 651387237, time.UTC)
	fmt.Println("Data aleasa: ", choseDate)

	diferenta := now.Sub(choseDate)
	zile := int16 (diferenta.Hours()/24)
	fmt.Println("Diferenta intre doua date in zile: ", zile)

	year, month, day, hour, min, sec := diff(choseDate, now)
	fmt.Printf("Diferenta %d years, %d months, %d days, %d hours, %d mins and %d seconds.", year, month, day, hour, min, sec)
	fmt.Println(" ")

	fmt.Println("Ziua saptamanii a datei alese: " ,choseDate.Weekday())

	daysToAdd := 5
	daysToSubstract := 3
	fmt.Println("Data plus numar de zile: ", choseDate.AddDate(0, 0, daysToAdd))
	fmt.Println("Data minus numar de zile: ", choseDate.AddDate(0, 0, (-daysToSubstract)))

}
