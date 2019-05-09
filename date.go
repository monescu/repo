package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var numar int
	optiune := 0

	for true {
		fmt.Println("1. Diferenta dintre doua date in zile.")
		fmt.Println("2. Diferenta dintre doua date calendaristice in ani, luni, saptamani, zile.")
		fmt.Println("3. Data + nr zile.")
		fmt.Println("4. Data - nr zile.")
		fmt.Println("5. In ce zi a saptamanii se gaseste o data calendaristica.")
		fmt.Println("6. Exit.")
		fmt.Println("Alege o optiune.")
		fmt.Scanln(&optiune)
		switch optiune {
		case 1:
			firstDate :=citireData()
			secondDate := citireData()
			diferentaInZile(firstDate,secondDate)
			break
		case 2:
			firstDate := citireData()
			secondDate := citireData()
			diferentaInAni(firstDate, secondDate)
			break
		case 3:
			data := citireData()
			fmt.Println("Cate zile vrei sa adaugi la aceasta data?")
			_,err := fmt.Scanln(&numar)
			if(err != nil){
				fmt.Println("Numar invalid!")
				break
			}
			dateaPlusZile(data,numar)
			break
		case 4:
			data := citireData()
			fmt.Println("Cate zile vrei sa scazi de la aceasta data?")
			_,err := fmt.Scanln(&numar)
			if(err != nil){
				fmt.Println("Numar invalid!")
				break
			}
			dateaMinusZile(data,numar)
			break
		case 5:
			data := citireData()
			ziuaDinSaptamanaADatei(data)
			break
		case 6:
			os.Exit(1)
		default:
			fmt.Println("Optiune invalida!")
		}
	}
}
func citireData() (data time.Time){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Introdu o data (aaaa-ll-zz). \n")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)
	data, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Printf("Data nu este valida!")
		os.Exit(0)
	}
	return data
}

func diferentaInZile(firstDate , secondDate time.Time){
	diferenta := firstDate.Sub(secondDate)
	zile := int16 (diferenta.Hours()/24)
	if(zile<0){
		zile *= -1
	}
	fmt.Println("Diferenta dintre cele doua date in zile: ", zile)
	return
}

func diferentaInAni(firstDate , secondDate time.Time){

	an1, _, _ := firstDate.Date()
	an2, _, _ := secondDate.Date()
	ani := an1-an2
	if(ani<0){
		ani *= -1
	}

	fmt.Println("Diferenta dintre cele doua date in ani: ", ani, " ani.")
	return
}

func dateaPlusZile(date time.Time, zile int){
	fmt.Println("Data minus numar de zile: ", date.AddDate(0, 0,(- zile)))
	return
}

func dateaMinusZile(date time.Time, zile int){
	fmt.Println("Data minus numar de zile: ", date.AddDate(0, 0,(- zile)))
	return
}

func ziuaDinSaptamanaADatei(date time.Time){
	fmt.Println("Data se gaseste in ziua: " ,date.Weekday())
	return
}


