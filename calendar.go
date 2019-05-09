package main

import "fmt"
import "time"

func calcul(a, b time.Time) (an, luna, zi int) {
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	an= int(y2 - y1)
	luna = int(M2 - M1)
	zi = int(d2 - d1)


	if zi < 0 {
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		zi += 32 - t.Day()
		luna--
	}
	if luna < 0 {
		luna += 12
		an--
	}

	return
}

func main() {

	now := time.Now()
	fmt.Println("Data curenta: ", now)

	choseDate := time.Date(
		2017, 3, 15, 11, 34, 58, 651387237, time.UTC)
	fmt.Println("Data aleasa: ", choseDate)

	an, luna, zi := calcul(choseDate, now)
	fmt.Printf("Diferenta %d ani, %d luni, %d zile.", an, luna, zi)
	fmt.Println(" ")

	diferenta := now.Sub(choseDate)
	zile := int16 (diferenta.Hours()/24)
	fmt.Println("Diferenta intre doua date in zile: ", zile)

	fmt.Println("Ziua saptamanii a datei alese: " ,choseDate.Weekday())

	daysToAdd := 7
	daysToSubstract := 15
	fmt.Println("Data plus numar de zile: ", choseDate.AddDate(0, 0, daysToAdd))
	fmt.Println("Data minus numar de zile: ", choseDate.AddDate(0, 0, (-daysToSubstract)))

}
