package ch2

func SelectionSort(A []int) {
	for i := 0; i < len(A) - 1; i++ {
		tmp := i
		for j := i + 1; j < len(A); j++ {
			if A[tmp] > A[j] {
				tmp = j
			}
		}
		A[tmp], A[i] = A[i], A[tmp]
	}
}