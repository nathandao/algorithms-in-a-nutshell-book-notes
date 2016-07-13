package main

// Binary search is only applicable for sorted arrays.
func BinarySearch(s []int, searchVal int) bool {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := (low + high) / 2
		if searchVal < s[mid] {
			high = mid - 1
		} else if searchVal > s[mid] {
			low = mid + 1
		} else {
			return true
		}
	}

	return false
}
