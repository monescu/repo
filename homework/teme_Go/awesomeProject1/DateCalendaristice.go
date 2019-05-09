package main
// Lemnaru Rebecca - anul 3 ID
// Tema 04. Operatii cu date calendaristice, aplicatie  Google GO
//			1- diferenta dintre doua date in zile;
//			2- diferenta dintre doua date calendaristice in ani, luni, saptamani, zile
//			3- data + nr zile
//			4- data - nr zile
//			5- in ce zi a saptamanii se gaseste o data calendaristica
import (
	"fmt"
	"math"
	"time"
	"os"
)

func main() {
	fmt.Println()

	var starting_year, starting_month, starting_day, end_year, end_month, end_day, nr_days_to_add, nr_days_to_substract int

	fmt.Print("Starting date [YEAR]:\t\t")
	fmt.Scanln(&starting_year)
	if (starting_year < 0 ) {
		fmt.Println("\t\t\t\terror - only years over 0")
		os.Exit(2)
	}

	fmt.Print("Starting date [MONTH]:\t\t")
	fmt.Scanln(&starting_month)
	if (starting_month <= 0 || starting_month > 12) {
		fmt.Println("\t\t\t\terror - please enter a MONTH[1..12]")
		os.Exit(2)
	}

	fmt.Print("Starting date [DAY]:\t\t")
	fmt.Scanln(&starting_day)
	if (starting_day <= 0 || starting_day > 31) {
		fmt.Println("\t\t\t\terror - please enter a DAY[1..31]")
		os.Exit(2)
	}

	fmt.Println()

	fmt.Print("End date [YEAR]:\t\t")
	fmt.Scanln(&end_year)
	if (end_year < 0 ) {
		fmt.Println("\t\t\t\terror - only years over 0")
		os.Exit(2)
	}

	fmt.Print("End date [MONTH]:\t\t")
	fmt.Scanln(&end_month)
	if (end_month <= 0 || end_month > 12) {
		fmt.Println("\t\t\t\terror - please enter a MONTH[1..12]")
		os.Exit(2)
	}


	fmt.Print("End date [DAY]:\t\t\t")
	fmt.Scanln(&end_day)
	if (end_day <= 0 || end_day > 31) {
		fmt.Println("\t\t\t\terror - please enter a DAY[1..31]")
		os.Exit(2)
	}

	fmt.Println()
	fmt.Print("Please enter a number of days to be added to the end date:\t\t")
	fmt.Scanln(&nr_days_to_add)

	fmt.Print("Please enter a number of days to be substracted from the end date\t\t")
	fmt.Scanln(&nr_days_to_substract)

	fmt.Println()

	const layout_ro = "02-01-2006"
	layout := layout_ro 	// formatul ales de afisare a datei

	t1 := Date(starting_year, starting_month, starting_day)
	t2 := Date(end_year, end_month, end_day)

	// cerinta 1 - diferenta dintre doua date in zile;
	days := t2.Sub(t1).Hours() / 24  // variabila days contine diferenta in zile dintre cele doua date calendaristice definite la inceput

	// cerinta 2 - diferenta dintre doua date calendaristice in ani, luni, saptamani, zile
	years := math.Floor(days / 365)  // aproximeaza in jos numarul de ani
	months := math.Floor((days - years*365) / 30)  // aproximeaza in jos numarul de luni
	weeks := math.Floor((days - years*365 - months*30) / 7)  // aproximeaza in jos numarul de saptamani
	day := days - years*365 - months*30 - weeks*7 // calculeaza numarul de zile ramase din calcul

	// cerinta 3 - data + nr zile
	add_days := t2.AddDate(0, 0, nr_days_to_add)

	// cerinta 4 - data - nr zile
	substract_days := t2.AddDate(0, 0, nr_days_to_substract*-1)

	// cerinta 5 - in ce zi a saptamanii se gaseste o data calendaristica
	day_of_week := t2.Weekday()

	// afisare rezultate
	fmt.Println("\tFrom date:\t", t1.Format(layout))
	fmt.Println("\t  To date:\t", t2.Format(layout))
	fmt.Printf("Difference between the 2 dates (in days): %.f days\n", days)
	fmt.Printf("Difference between the 2 dates: %v years, %v months, %v weeks, %v days. \n", years, months, weeks, day)
	fmt.Printf("%v + %v\tdays -> %v \n", t2.Format(layout), nr_days_to_add, add_days.Format(layout))
	fmt.Printf("%v - %v\tdays -> %v \n", t2.Format(layout), nr_days_to_substract, substract_days.Format(layout))
	fmt.Printf("What day of the week is the end date %v? %v", t2.Format(layout), day_of_week)
	fmt.Println()
}

func Date(year, month, day int) time.Time {  // functia converteste datele de intrare in format time.Date
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}