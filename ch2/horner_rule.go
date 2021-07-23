package ch2

func NaiveHornerRule(A[] int, x int) int {
	y := 0
	for k := 0; k < len(A); k++ {
		tmp := 1
		for i := 1; i <= k; i++ {
			tmp *= x
		}
		y += A[k] * tmp
	}
	return y
}
func HornerRule(A[] int, x int) int {
	y := 0
	for i := len(A) - 1; i >= 0; i-- {
		y = A[i] + x * y
	}
	return y
}
