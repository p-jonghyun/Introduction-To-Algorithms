package ch2

// O(N^2)
// It's ok to pass slice by a value
// since slice has a pointer to array
func InsertionSort(A []int) {
	for j := 1; j < len(A); j++ {
		key := A[j]
		// Insert A[j] into the sorted sequence A[0..j-1]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}