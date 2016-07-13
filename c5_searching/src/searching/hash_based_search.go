package main

import (
	"container/heap"
)

type HashList []int

// Implementing IntHeap interface for HashList
// https://golang.org/pkg/container/heap/#example__intHeap
func (b HashList) Len() int           { return len(b) }
func (h HashList) Less(i, j int) bool { return h[i] < h[j] }
func (h HashList) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HashList) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *HashList) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Hash function
func hash(v int, tableSize int) int {
	// Only get positive values for index
	if v < 0 {
		v = -v
	}
	return v % tableSize
}

func loadTable(s []int) []HashList {
	hTable := make([]HashList, len(s))
	for _, v := range s {
		idx := hash(v, len(s))
		heap.Push(&hTable[idx], v)
	}
	return hTable
}

func HashBasedSearch(s []int, v int) bool {
	hTable := loadTable(s)
	hIdx := hash(v, len(s))

	if len(hTable[hIdx]) == 0 {
		return false
	} else {
		return BinarySearch(hTable[hIdx], v)
	}
}
