package main

import (
	"fmt"
	"math"
)

type Fractie struct {
	Numarator int
	Numitor  int

}


func cmmdc( x, y int) (int){
	for x%y!=0{
		r:=x%y
		x=y
		y=r
	}
	return y
}

func(f *Fractie) simplificare(){
	c:=cmmdc(f.Numarator,f.Numitor)
	f.Numarator/= c
	f.Numitor /= c

}

func (f Fractie) suma(f1 Fractie) Fractie{
	s:=Fractie{f1.Numitor*f.Numarator+f.Numitor*f1.Numarator,f.Numitor*f1.Numitor}
	s.simplificare()
	return s
}

func (f Fractie) dif(f1 Fractie) Fractie{
	s:=Fractie{f1.Numitor*f.Numarator-f.Numitor*f1.Numarator,f.Numitor*f1.Numitor}
	s.simplificare()
	return s
}

func (f Fractie) prod(f1 Fractie) Fractie {
	s := Fractie{f.Numarator * f1.Numarator, f.Numitor * f1.Numitor}
	s.simplificare()
	return s
}

func (f Fractie) cat(f1 Fractie) Fractie {
	s := Fractie{f.Numarator * f1.Numitor, f.Numitor * f1.Numarator}
	s.simplificare()
	return s
}

func (f Fractie) getValue()float64{

	return float64(f.Numarator)/float64(f.Numitor)
}

func fact(n int) int{
	p:=1
	for i:=1;i<=n;{
		p*=i
		i++
	}
	return p
}

func e(eps float64) float64{
	t:=1.0
	n:=1
	for true {
		t1:=t
		t+=Fractie{1,fact(n)}.getValue()
		if math.Abs(t-t1)<eps{
			return t
		}
		n++
	}
	return -1
}

func main() {
	f:=Fractie{5,10}
	f.simplificare()
	fmt.Println(f)
	g:=Fractie{4,6}
	fmt.Println(f.suma(g))
	fmt.Println(f.prod(g))
	fmt.Println(f.dif(g))
	fmt.Println(f.cat(g).getValue())
	fmt.Println(e(1e-9))
}
