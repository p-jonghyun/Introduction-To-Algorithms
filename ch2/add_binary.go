package ch2

func AddBinary(A, B []int) []int {
	C := make([]int, len(A) + 1)
	carry := 0
	for i := len(A) - 1; i >= 0; i-- {
		C[i + 1] = (A[i] + B[i] + carry) % 2
		if A[i] + B[i] + carry >= 2 {
			carry = 1
		} else {
			carry = 0
		}
	}
	C[0] = carry
	return C
}