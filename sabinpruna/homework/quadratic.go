package homework

import "math/cmplx"

////fmt.Println(homework.SolveQuadratic(1, 5, 6)) //1 ,-2, 1

//SolveQuadratic returns the 2 values of a quadratic equation
func SolveQuadratic(a, b, c complex128) (positiveResult, negativeResult complex128) {
	discriminant := b*b - 4*a*c
	squareRoot := cmplx.Sqrt(complex128(discriminant))

	positiveResult = (-b + squareRoot) / (2 * a)
	negativeResult = (-b - squareRoot) / (2 * a)

	return positiveResult, negativeResult
}
