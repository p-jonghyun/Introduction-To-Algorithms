package ch2

func LinearSearch(A []int, v int) int {
	for i := 0; i < len(A); i++ {
		if A[i] == v {
			return i
		}
	}
	return -1
}