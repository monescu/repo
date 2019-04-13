package main
// Tema 04. Operatii cu date calendaristice, aplicatie  Google GO
//			1- diferenta dintre doua date in zile;
//			2- diferenta dintre doua date calendaristice in ani, luni, saptamani, zile
//			3- data + nr zile
//			4- data - nr zile
//			5- in ce zi a saptamanii se gaseste o data calendaristica
// Claudia Marinache - 10LD561
import (
	"fmt"
	"math"
	"time"
	"os"
)

func main() {
	//an_start, luna_start, zi_start := 2016, 4, 16 // declaratii locale cu initializare
	//an_stop, luna_stop, zi_stop := 2019, 4, 12
	//nr_zile_adunare := 10
	//nr_zile_scadere := 9

	fmt.Println()
		
	var an_start, luna_start, zi_start, an_stop, luna_stop, zi_stop, nr_zile_adunare, nr_zile_scadere int

	fmt.Print("Inceput interval [AN]:\t\t")
	fmt.Scanln(&an_start)
	if (an_start < 0 ) {
		fmt.Println("\t\t\t\teroare - interval permis doar e.n.")
		os.Exit(2)
	}	
	
	fmt.Print("Inceput interval [LUNA]:\t")
	fmt.Scanln(&luna_start)
	if (luna_start <= 0 || luna_start > 12) {
		fmt.Println("\t\t\t\teroare - interval permis LUNA[1..12]")
		os.Exit(2)
	}

	fmt.Print("Inceput interval [ZI]:\t\t")
	fmt.Scanln(&zi_start)
	if (zi_start <= 0 || zi_start > 31) {
		fmt.Println("\t\t\t\teroare - interval permis ZI[1..31]")
		os.Exit(2)
	}	

	fmt.Println()
		
	fmt.Print("Sfarsit interval [AN]:\t\t")
	fmt.Scanln(&an_stop)
	if (an_stop < 0 ) {
		fmt.Println("\t\t\t\teroare - interval permis doar e.n.")
		os.Exit(2)
	}	
	
	fmt.Print("Sfarsit interval [LUNA]:\t")
	fmt.Scanln(&luna_stop)
	if (luna_stop <= 0 || luna_stop > 12) {
		fmt.Println("\t\t\t\teroare - interval permis LUNA[1..12]")
		os.Exit(2)
	}


	fmt.Print("Sfarsit interval [ZI]:\t\t")
	fmt.Scanln(&zi_stop)
	if (zi_stop <= 0 || zi_stop > 31) {
		fmt.Println("\t\t\t\teroare - interval permis ZI[1..31]")
		os.Exit(2)
	}	
	
	fmt.Println()	
	fmt.Print("Nr. zile pentru adunare:\t")
	fmt.Scanln(&nr_zile_adunare)

	fmt.Print("Nr. zile pentru scadere:\t")
	fmt.Scanln(&nr_zile_scadere)

	fmt.Println()
	
	// https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
	const layout_en = "2 of January 2006"
	const layout_ro = "02-01-2006"
	//layout := layout_en
	layout := layout_ro 	// formatul ales de afisare a datei

	t1 := Date(an_start, luna_start, zi_start)
	t2 := Date(an_stop, luna_stop, zi_stop)
	//fmt.Println("t1= ",t1)
	//fmt.Println("t2= ",t2)
	
	// cerinta 1
	zile := t2.Sub(t1).Hours() / 24  // variabila zile contine diferenta in zile dintre cele doua date calendaristice definite la inceput
	//fmt.Println("zile= ",t2.Sub(t1))
	
	// cerinta 2	
	ani := math.Floor(zile / 365)  // aproximeaza in jos numarul de ani
	luni := math.Floor((zile - ani*365) / 30)  // aproximeaza in jos numarul de luni
	saptamani := math.Floor((zile - ani*365 - luni*30) / 7)  // aproximeaza in jos numarul de saptamani
	zi := zile - ani*365 - luni*30 - saptamani*7 // calculeaza numarul de zile ramase din calcul

	// cerinta 3
	adunare_zile := t2.AddDate(0, 0, nr_zile_adunare)
	
	// cerinta 4
	scadere_zile := t2.AddDate(0, 0, nr_zile_scadere*-1)

	// cerinta 5
	zi_din_saptamana := t2.Weekday()

	// afisare rezultate
	fmt.Println("\t  de la data:\t", t1.Format(layout))
	fmt.Println("\tpana la data:\t", t2.Format(layout))
	fmt.Printf("Diferenta dintre cele doua date calendaristice: %.f zile\n", zile)
	fmt.Printf("Diferenta dintre cele doua date calendaristice: %v ani, %v luni, %v saptamani, %v zile. \n", ani, luni, saptamani, zi)
	fmt.Printf("%v + %v\tzile >>> %v \n", t2.Format(layout), nr_zile_adunare, adunare_zile.Format(layout))
	fmt.Printf("%v - %v\tzile >>> %v \n", t2.Format(layout), nr_zile_scadere, scadere_zile.Format(layout))
	fmt.Printf("In ce zi a saptamanii se gaseste data calendaristica %v? %v", t2.Format(layout), zi_din_saptamana)
	fmt.Println()
}

func Date(year, month, day int) time.Time {  // functia converteste datele de intrare in format time.Date
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// verificare calcule cu
// https://www.timeanddate.com/date/durationresult.html?d1=16&m1=4&y1=2014&d2=16&m2=4&y2=2019