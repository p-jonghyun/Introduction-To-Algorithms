package ch2

// 5 2 3 8 1
func BubbleSort(A []int) {
	for i := 0; i < len(A) - 1; i++ {
		for j := len(A) - 1; j >= i + 1; j-- {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
		}
	}
}