package homework

//FindNumber return number from tth element right to left
func FindNumber(x int, p uint, t uint) int {
	x = x >> t
	x = x & (1<<p - 1)

	return x
}
