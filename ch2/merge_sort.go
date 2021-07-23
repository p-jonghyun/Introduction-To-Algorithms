package ch2

// O(NlogN)
func MergeSort(A []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		Merge(A, p, q , r)
	}
}

// O(N)
func Merge(A []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1)
	R := make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = A[p + i]
	}
	for j := 0; j < n2; j++ {
		R[j] = A[q + j + 1]
	}
	i, j := 0, 0
	for k := p; k <= r; k ++ {
		if i >= n1 {
			A[k] = R[j]
			j++
		} else if j >= n2 {
			A[k] = L[i]
			i++
		} else if L[i] > R[j] {
			A[k] = R[j]
			j++
		} else {
			A[k] = L[i]
			i++
		}
	}
}