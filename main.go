package main

import (
	"fmt"
	"time"
)

func anLunaZi (x,y time.Time) (an, luna, zi int){
	if x.Location()!=y.Location(){
		y=y.In(x.Location())
	}
	if x.After(y){
		x,y = y,x
	}
	a1, L1, z1 :=x.Date()
	a2, L2, z2 :=y.Date()

	an = int(a2-a1)
	luna = int(L2-L1)
	zi = int (z2-z1)

	if zi<0{
		t:=time.Date(a1, L1, 31, 0, 0, 0,0 , time.UTC)
		zi+=31-t.Day()
		luna--
	}
	if luna<0{
		luna+=12
		an--
	}
	return

}
func adunareZile(data time.Time, zi int)time.Time{
	return data.AddDate(0,0, zi)
}
func diferentaZile(data time.Time, zi int)time.Time{
	return data.AddDate(0,0, zi)
}

func main() {
	acum := time.Now()
	fmt.Println(" Suntem in data de: ", acum)
	alegeData := time.Date(2019, 04, 05, 18, 22, 30, 0, time.UTC)
	year,month,day:=alegeData.Date()

	fmt.Println(" Data pe care ati ales-o este: ")
	fmt.Println(" ",year, "/", month, "/", day)

	diferentaInZile := acum.Sub(alegeData)
	zile := int16(diferentaInZile.Hours() / 24)
	fmt.Println("\n Diferenta exprimata in zile, dintre cele doua date este: ", zile)

	an, luna, zi:= anLunaZi(alegeData, acum)
	fmt.Println("\n Diferenta exprimata in ani, luni, zile, dintre cele 2 date este: ",
		an, luna, zi)

	fmt.Println("\n Data aleasa + 8 zile: ", adunareZile(alegeData,8).String())
	fmt.Println(" Data aleasa - 5 zile: ", diferentaZile(alegeData,-5).String())

	fmt.Println("\n Data pe care ati ales-o cade intr-o: ", alegeData.Weekday())

}
