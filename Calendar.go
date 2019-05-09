package main

import (
"fmt"
"bufio"
"os"
)

type Date struct {
   day int
   month int
   year int
   dayOfWeek int
}
func main() {
   var date1 Date   /* Declare day1 */
   var date2 Date   /* Declare day2 */
   
   // Declare the array of how many days per month
   daysPerMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
 
   /* day 1 specification */
   date1.day = 20
   date1.month = 3
   date1.year = 1993
   date1.dayOfWeek = 1

   /* day 2 specification */
   date2.day = 21
   date2.month = 4
   date2.year = 2012
   date2.dayOfWeek = 1
 
   /* print day1 info */
   printDate(&date1)

   /* print day2 info */
   printDate(&date2)
   
   //compare dates
   var res int = comp(&date1, &date2)
   fmt.Printf("Comp date %d\n", res)
   
   //diferenta zile
   var dif int = difZile(&date1, &date2, daysPerMonth)
   fmt.Printf(" Dif zile: %d\n", dif)
   
   //Find out the day of week for the following:
   var date3 Date
   date3.day = 14
   date3.month = 4
   date3.year = 2012
   date3.dayOfWeek = 0
   
   CheckTheDay(&date1)
}
func printDate( date *Date ) {
   fmt.Printf( "date : %d\n", date.day);
   fmt.Printf( "month : %d\n", date.month);
   fmt.Printf( "year : %d\n", date.year);
   fmt.Printf( "dayOfWeek : %d\n", date.dayOfWeek);
   
}
func comp (date1 *Date, date2 *Date) int{

	//var result int = 0
	result := 0
	if(date1.year < date2.year){
		result = -1
		return result
	}
	if(date1.year > date2.year){
		result = 1
		return result
		}
	if(date1.month < date2.month){
		result = -1
		return result
		}
	if(date1.month > date2.month){
		result = 1
		return result
		}
	if(date1.day < date2.day){
		result = -1
		return result
		}
	if(date1.day > date2.day){
		result = 1
		return result
		}
	return result
}
func bisect (d *Date) int{
	if (d.year % 4 == 0){
		if (d.year % 100 == 0){
			return 0
		}
		if (d.year % 400 == 0){
			return 1
		}
	}
	return 0
}


func difZile(d1 *Date, d2 *Date, daysPerM [12] *int) int{
	t := 0
	res := 0
	res = comp(d1, d2)
	fmt.Printf("Comp date %d\n", res)
	for comp(d1,d2) != 0{
	if (res == -1){
		if (&d1.day < &daysPerM[&d1.month]){
			d1.day++
			t++
		}
		if (&d1.day >= daysPerM[d1.month]){
			if (d1.month == 1 && bisect(d1) == 1){
				d1.day ++
				t++
			}
			if (d1.month == 12){
				d1.month = 1
				d1.year ++
				d1.day = 1
				t ++
			}
			if (d1.month < 12){
				d1.day = 1
				d1.month ++
				t ++
			}
		}
	}
	}
	return t
}

func CheckTheDay(d2 *Date, d3 *Date, daysPerM [12] *int) int{
	var dif int = difZile(&date1, &date2, daysPerM)
	
	if ((7-dif%7)%7 == 0)
		&d3.dayOfWeek = &d2.dayOfWeek
	if ((7-dif%7)%7 == 1)
		&d3.dayOfWeek = &d2.dayOfWeek + 1
	if ((7-dif%7)%7 == 2)
		&d3.dayOfWeek = &d2.dayOfWeek + 2
	if ((7-dif%7)%7 == 3)
		&d3.dayOfWeek = &d2.dayOfWeek + 3
	if ((7-dif%7)%7 == 4)
		&d3.dayOfWeek = &d2.dayOfWeek + 4
	if ((7-dif%7)%7 == 5)
		&d3.dayOfWeek = &d2.dayOfWeek + 5
	if ((7-dif%7)%7 == 6)
		&d3.dayOfWeek = &d2.dayOfWeek + 6
	
}