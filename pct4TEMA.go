package main

import "fmt"

type date struct{
	zi int
	luna int
	an int
}



func ebisect(d2 date) bool{

	if d2.an % 4 == 0 {
		if d2.an % 100 == 0 {
			if d2.an % 400 == 0{
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	}
	return false
}

func plusUnu(d2 date) date{

	var zileMaxLuna[13] int
	zileMaxLuna[0] = 0
	zileMaxLuna[1] = 31
	zileMaxLuna[2] = 28
	zileMaxLuna[3] = 31
	zileMaxLuna[4] = 30
	zileMaxLuna[5] = 31
	zileMaxLuna[6] = 30
	zileMaxLuna[7] = 31
	zileMaxLuna[8] = 31
	zileMaxLuna[9] = 30
	zileMaxLuna[10] = 31
	zileMaxLuna[11] = 30
	zileMaxLuna[12] = 31

	if d2.zi < zileMaxLuna[d2.luna]{
		d2.zi++
	}

	if d2.zi >= zileMaxLuna[d2.luna]{
		if d2.luna == 2 && ebisect(d2) && d2.zi < 29{
			d2.zi ++
		} else if d2.luna == 12  {
				d2.zi = 1
				d2.luna = 1
				d2.an ++
		} else if d2.luna < 12{
					d2.zi = 1
					d2.luna ++
		}
	}
	return d2
	}

func compara(primaData date, aDouaData date) int{
	if primaData.an < aDouaData.an {
		return -1
	}
	if primaData.an > aDouaData.an{
		return 1
	}
	if primaData.luna < aDouaData.luna{
		return -1
	}
	if primaData.luna > aDouaData.luna{
		return 1
	}
	if primaData.zi < aDouaData.zi{
		return -1
	}
	if primaData.zi > aDouaData.zi{
		return 1
	} else {return 0}
}

func numaraZile (primaData date, aDouaData date) int{
	var nrZile int = 0

	if compara(primaData, aDouaData) == -1 {
		for compara(primaData, aDouaData) == - 1 {
			primaData = plusUnu(primaData)
			nrZile ++
		}
	} else if compara(primaData, aDouaData) == 1{
		for compara(aDouaData, primaData) == - 1 {
			aDouaData = plusUnu(aDouaData)
			nrZile ++
		}
	} else {nrZile = 0}

	return nrZile
}


func numaraALSZ (primaData date, aDouaData date) (nrAni,nrLuni, nrSaptamani, nrZile int){

	for compara(primaData, aDouaData) == -1{
			primaData.an++
			nrAni++
		}
	primaData.an--
	nrAni--

	for compara(primaData, aDouaData) == -1 {
		primaData.luna++
		if primaData.luna>12{
			primaData.luna=1
			primaData.an++
		}

		nrLuni++
	}
	primaData.luna--
	nrLuni--

 	nrSaptamani = numaraZile(primaData, aDouaData)/7
	nrZile = numaraZile(primaData, aDouaData)%7
	return nrAni,nrLuni, nrSaptamani, nrZile
}

func main(){

	d1 := date{zi: 02, luna: 5, an: 1985}
	d2 := date{zi: 25, luna: 4, an: 2018}

 c := ebisect(d1)

 fmt.Println(c)
 /*fmt.Println("data initiala:", d1.zi, d1.luna, d1.an)

 df := plusUnu(d1)
 fmt.Println("data plus unu:", df.zi, df.luna, df.an)*/
	//zile := numaraZile(d1,d2)
	//dif := numaraALSZ(d1, d2)
	fmt.Println("Intre cele doua date sunt:", numaraZile(d1,d2),"zile")
	fmt.Print("Intre cele doua date sunt:")
	fmt.Println( numaraALSZ(d1,d2))
	fmt.Println("                       ani / luni / saptamani / zile")

}