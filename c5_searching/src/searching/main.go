package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Generate a slice comprises of  n-size arrays, in which the later is 2 times
// larger than the former.
func generateBenchmarkSlices() ([][]int, []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	slices := [][]int{}
	searchVals := []int{}

	for i := 1; i <= 18; i++ {
		size := int(math.Pow(2, float64(i)))
		slices = append(slices, make([]int, size))

		// Add random values to slice.
		for j := 0; j < size; j++ {
			slices[i-1][j] = r.Int()
		}

		// Get a random key and use value at that the search value.
		key := r.Intn(size - 1)
		searchVals = append(searchVals, slices[i-1][key])
	}

	return slices, searchVals
}

func main() {
	testSlice := []int{19, 23, 12, 32, 44, 12, 10, 93, 41, 234, 5, 32, 55, 23}
	valFound := 234
	valNotFound := 9999

	// Slices & search values for benchmarking.
	slices, searchVals := generateBenchmarkSlices()

	methods := map[string]func([]int, int) bool{
		"SEQUENTIAL SEARCH": SequentialSearch,
	}

	for searchMethod, searchHandler := range methods {
		fmt.Println("\n=====", searchMethod, "=====")

		// Visual test to see if seach works correctly.
		fmt.Println(testSlice)
		fmt.Println("Search", valFound, ":", searchHandler(testSlice, valFound),
			", Search", valNotFound, ":", searchHandler(testSlice, valNotFound))
		fmt.Println("-------------------------------")

		// Benchmark info.
		for i, s := range slices {
			start := time.Now().UnixNano()
			searchHandler(s, searchVals[i])
			execTime := time.Now().UnixNano() - start
			fmt.Println("Size:", len(s), ",", "Execution time:", execTime)
		}

	}
}
