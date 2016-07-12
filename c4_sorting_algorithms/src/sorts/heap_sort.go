package main

func HeapSort(s []int) []int {
	// Build heap
	buildHeap(s)

	// Iterate heap sort (n - 1) times, each time move the largest element to the
	// end and find largest element of the remaining elements and move the largest
	// element of the end of the remaining.
	for i := len(s) - 1; i >= 1; i-- {
		temp := s[0]
		s[0] = s[i]
		s[i] = temp

		heapify(s, 0, i)
	}

	return s
}

func buildHeap(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		heapify(s, i, len(s))
	}
}

func heapify(s []int, idx int, max int) {
	largest := idx
	left := 2*idx + 1
	right := 2*idx + 2

	if left < max && s[left] > s[idx] {
		largest = left
	}

	if right < max && s[right] > s[largest] {
		largest = right
	}

	if largest != idx {
		temp := s[idx]
		s[idx] = s[largest]
		s[largest] = temp

		heapify(s, largest, max)
	}
}
