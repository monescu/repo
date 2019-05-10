package main

import "fmt"

type Data struct{
	zi, luna, an int
}


var zile_luni = [13]int{0,31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
func citireData()(Data){
var dataCitita Data
	var zi,luna,an int

	for  {
		fmt.Print("Anul: ")
		fmt.Scanf("%d",&an)
		if (an<1 ){
			fmt.Println("Anul trebuie sa fie mai mare decat 0! ")
		}else{
			break
		}
	}

	for  {
		fmt.Print("Luna: ")
		fmt.Scanf("%d",&luna)
		if (luna<1 || luna>12){
			fmt.Println("Luna are valori intre 1 si 12! ")
		}else{
			break
		}
	}

	for  {
		fmt.Print("Ziua: ")
		fmt.Scanf("%d", &zi)
		if luna==2{
			if anbisect(an){
				if (zi<1 || zi>29){
					fmt.Println("Ziua are valori intre 1 si 29! ")
				}else{
					break
				}
			}else{
				if (zi<1 || zi>28){
					fmt.Println("Ziua are valori intre 1 si 28! ")
				}else{
					break
				}
			}

		}else if (zi<1 || zi>31){
			fmt.Println("Ziua are valori intre 1 si 31! ")
		}else{
			break
		}
	}


	dataCitita.zi=zi
	dataCitita.luna=luna
	dataCitita.an=an

	return dataCitita
}

func anbisect(an int) bool{
	// verific daca anul este bisect
	// dacă (an nu este divizibil cu 4) atunci (an obișnuit)
	// altfel dacă (an nu este divizibil cu 100) atunci (an bisect)
	// altfel dacă (an nu este divizibil cu 400) atunci (an obișnuit)
	// altfel (an bisect)

	if an % 100 == 0 && an % 400 == 0{
		return true
	}
	if an % 100 != 0 && an % 4 == 0{
		return true
	}
	return false
}
//functia plus o zi
func plus_1zi(d Data) Data{
	nrZileLuni := [13]int{0,31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if d.zi<nrZileLuni[d.luna]{
		d.zi++
	}else{
		if d.zi==nrZileLuni[d.luna] {
			if d.luna == 2 && anbisect(d.an)  {
				d.zi++
			}else{
				if d.luna==12 {
					d.zi=1
					d.luna=1
					d.an++

				}else{
					d.zi=1
					d.luna++
				}
			}

		}else{
			if d.luna == 2 && anbisect(d.an)  {
				d.zi=1
				d.luna++
			}
		}

	}

	return d
}
// functia minus o zi
func minusOZi(d Data) Data {
	nrZileLuni := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if d.zi > 1 {
		d.zi--
	} else {
		//daca este 1 martie atunci va fi 28 feb sau 29 feb
		if d.luna == 3 {
			if anbisect(d.an) {
				d.luna--
				d.zi = nrZileLuni[d.luna] + 1
			} else {
				d.luna--
				d.zi = nrZileLuni[d.luna]
			}
		}
		// daca e 1 ian va fi 31 dec si un an mai putin
		if d.luna == 1 {
			d.luna = 12
			d.zi = nrZileLuni[d.luna]
			d.an--
		}else{
			//cand nu e ian si nici feb
			d.luna--
			d.zi = nrZileLuni[d.luna]
		}

	}

	return d
}
func minus1zi(d Data, zile int) Data{
	for i:=0;i<zile;i++{
		d=minus1zi(d, 1)
	}

	return d
}
//functia pentru a compara 2 date
func comparDate(d1, d2 Data) int{
	if d1.an<d2.an{
		return -1
	}
	if d1.an>d2.an{
		return 1
	}
	if d1.luna<d2.luna{
		return -1
	}
	if d1.luna>d2.luna{
		return 1
	}
	if d1.zi<d2.zi{
		return -1
	}else{
		return 1
	}
	return 0
}
//functia pentru a numara zilele
func numarZile(d1,d2 Data) int{
	var totalZile int
	if comparDate(d1,d2)== -1{
		for (comparDate(d1,d2)== -1){
			d1=plus_1zi(d1)
			totalZile++
		}
	}else if comparDate(d1,d2)== 1{
		for (comparDate(d2,d1)== -1){
			d2=plus_1zi(d2)
			totalZile--
		}

	}else
	{
		totalZile=0
	}
	return totalZile
}
//functia pentru a stabili ce zi a saptamanii este
func ziuasaptamana(d Data) string{
	saptamana:=[7] string {"Luni","Marti","Miercuri","Joi","Vineri","Sambata", "Duminica"}
	dataStart:= Data{3,1,2011} //data de start pentru  a compara 3.1.2011 este luni
	var nrzile int
	if comparDate(dataStart,d)==-1{
		nrzile=numarZile(dataStart,d)
		return saptamana[nrzile%7]
	}else if comparDate(dataStart,d)==1{
		nrzile=numarZile(d,dataStart)
		return saptamana[(7-(nrzile%7))%7]
	}else{
		return saptamana[0]
	}

}

func anLuniSaptamaniZile(zileTotal int) (int,int,int,int){
	var an,luni,saptamani,zile int
	an=zileTotal/365
	zileTotal-=an*365
	if an>4{
		zileTotal-=an/4

	}
	if zileTotal>0 {

		luni = int( float64(zileTotal) /30.5)
		zileTotal-=int(float64(luni)*30.5)

		if zileTotal>0 {
			saptamani = zileTotal  / 7
			zileTotal-=saptamani*7
			zile = zileTotal
		}
	}
	return an,luni,saptamani,zile
}
func main() {
	var data1,data2,data3 Data
	fmt.Println(" Student Sraier Alina-Gabriela, anul III ID, aplicatie cu date calendaristice\n ")
	fmt.Println(" Introdu prima data calendaristica:\n ")
	data1 = citireData()
	fmt.Printf(" Data de %v este  %s \n",data1,ziuasaptamana(data1))

	fmt.Print("Introdu a doua data calendaristica: \n")
	data2 = citireData()
	fmt.Printf("Data de %v este %s \n",data2,ziuasaptamana(data2))

	nrZile:=numarZile(data1,data2)
	fmt.Printf("\nIntre datele %v - %v aven %d zile \n",data1,data2,nrZile)

	ani, luni, saptamani, zile := anLuniSaptamaniZile(nrZile)
	fmt.Printf(" \n Numarul de zile = %d insemna %d ani,%d luni %d saptamani, %d zile \n",nrZile,ani,luni,saptamani,zile)


	data3=plus_1zi(data2)
	fmt.Printf("Data %v + 1 zile = %v",data2,data3)
	fmt.Printf(" ziua din saptamana= %s \n",ziuasaptamana(data3))

	data3=minus1zi(data2,1)
	fmt.Printf("Data %v - 1 zile = %v",data2,data3)
	fmt.Printf(" ziua din saptamana= %s \n",ziuasaptamana(data3))
}