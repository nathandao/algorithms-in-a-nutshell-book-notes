package main

import (
	"container/heap"
	"math"
	"strconv"
)

type Bucket []int

// Implementing IntHeap interface for Bucket.
// https://golang.org/pkg/container/heap/#example__intHeap
func (b Bucket) Len() int           { return len(b) }
func (h Bucket) Less(i, j int) bool { return h[i] < h[j] }
func (h Bucket) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Bucket) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *Bucket) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BucketSort(s []int) []int {
	// Set bucket size to be 10, 100, 1000, ... depending on the list size.
	bl := int(math.Pow(10, float64(len(strconv.Itoa(len(s)))-2)))
	if bl > 1000 {
		bl = 1000
	}
	buckets := make([]Bucket, bl)

	// Get max & min value
	max := s[0]
	min := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	// Calculate the range width for each bucket
	width := (max-min)/bl + 1

	// Hashing & placing values into buckets. Heap is used so that sorting inside
	// buckets happens on the fly.
	for _, v := range s {
		idx := (v - min) / width
		heap.Push(&buckets[idx], v)
	}

	// Join buckets together.
	idx := 0
	for i := 0; i < bl; i++ {
		for _, _ = range buckets[i] {
			s[idx] = heap.Pop(&buckets[i]).(int)
			idx++
		}
	}

	return s
}
