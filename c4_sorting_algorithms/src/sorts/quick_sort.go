package main

import (
	"math/rand"
)

func QuickSort(s []int) []int {
	doQuickSort(s, 0, len(s)-1)
	return s
}

func doQuickSort(s []int, left, right int) {
	if left < right {
		pivot := partition(s, left, right)
		doQuickSort(s, left, pivot-1)
		doQuickSort(s, pivot+1, right)
	}
}

// Partition choose a random pivot point in the slice, move all smaller values
// to the left of the pivot, and larger values to the right of the pivot.
func partition(s []int, left, right int) int {

	pivot := rand.Intn(right-left) + left
	// Swap pivot to the outter left.
	if pivot != left {
		temp := s[left]
		s[left] = s[pivot]
		s[pivot] = temp
		pivot = left
	}

	for i := left + 1; i <= right; i++ {
		if s[i] < s[pivot] {
			// Shift smaller value to left of pivot.
			temp := s[pivot]
			s[pivot] = s[i]
			s[i] = s[pivot+1]
			s[pivot+1] = temp

			// Move pivot up
			pivot = pivot + 1
		}
	}

	return pivot
}
