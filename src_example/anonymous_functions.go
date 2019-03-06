package main

func add(x, y int) int{
	return x + y
}

func substract(x, y int) int{
	return x - y
}

func multiply(x, y int) int{
	return x * y
}

func divide(x, y int) int{
	return x / y
}

func operate(x , y int, operatie func(int, int) int)int{
	return operatie(x, y)

}

func swap(x *int, y *int){
	*x, *y = *y, *x
}


func main() {
	a:=operate(2,3,add)
	println(a)
	a=operate(2,3,substract)
	println(a)
	a=operate(2,3,multiply)
	println(a)
	a=operate(2,3,divide)
	println(a)


	//swap
	x := 2
	y := 7
	println(x, y)
	swap(&x, &y)
	println(x, y)
}
