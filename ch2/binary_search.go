package ch2

// O(logN)
func IterativeBinarySearch(A []int, v, low, high int) int {
	for low <= high {
		mid := (low + high) / 2
		if A[mid] == v {
			return mid
		}
		if A[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// O(logN)
func RecursiveBinarySearch(A []int, v, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if A[mid] == v {
		return mid
	}
	if A[mid] > v {
		return RecursiveBinarySearch(A, v, low, mid - 1)
	} else {
		return RecursiveBinarySearch(A, v, mid + 1, high)
	}
}