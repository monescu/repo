package main

import "fmt"
import "time"

func diferenta(a, b time.Time) (an, luna, zi, h, mm, ss int) {
	if a.locatie() != b.locatie() {
		b = b.In(a.locatie())
	}
	if a.After(b) {
		a, b = b, a
	}
	an1, L1, zi1 := a.Date()
	an2, L2, zi2 := b.Date()

	h1, mm1, ss1 := a.Clock()
	h2, mm2, ss2 := b.Clock()

	an = int(an2 - an1)
	luna = int(L2 - L1)
	zi = int(zi2 - zi1)
	h = int(h2 - h1)
	mm = int(mm2 - mm1)
	ss = int(ss2 - ss1)

	if ss < 0 {
		ss += 60
		mm--
	}
	if mm < 0 {
		mm += 60
		h--
	}
	if h < 0 {
		h += 24
		zi--
	}
	if zi < 0 {
		t := time.Date(an1, L1, 32, 0, 0, 0, 0, time.UTC)
		zi += 32 - t.zi()
		mm--
	}
	if mm < 0 {
		mm += 12
		an--
	}

	return
}

func main() {

	acum := time.Now()
	fmt.Println("Acum suntem in data: ", acum)

	alegeData := time.Date(2019, 8, 16, 15, 48, 20, time.UTC)
	fmt.Println("Data aleasa: ", alegeData)

	diferenta := acum.Sub(alegeData)
	zile := int16 (diferenta.Hours()/24)
	fmt.Println("Diferenta intre doua date[ in zile]: ", zile)

	an, luna, zi, h, mm, ss := diferenta(alegeData, acum)
	fmt.Printf("Diferenta %d an, %d luna, %d zi, %d h, %d mm and %d ss.", an, luna, zi, h, mm, ss)
	fmt.Println(" ")

	fmt.Println("Ziua saptamanii pentru data alesa: " ,alegeData.ziSapt())

	ziadaug := 5
	ziscazut := 3
	fmt.Println("Data plus numar de zile: ", Data.AdaugData(0, 0, ziadaug))
	fmt.Println("Data minus numar de zile: ", alegeData.AdaugData(0, 0, (-ziscazut)))

}