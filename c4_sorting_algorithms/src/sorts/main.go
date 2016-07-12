package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate a slice comprises of  n-size arrays, in which the later is 2 times
// larger than the former.
func generateBenchmarkSlices() [][]int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	slices := [][]int{
		make([]int, 10),
		make([]int, 20),
		make([]int, 40),
		make([]int, 80),
		make([]int, 160),
		make([]int, 320),
		make([]int, 640),
		make([]int, 1280),
		make([]int, 2560),
		make([]int, 5120),
		make([]int, 10240),
		make([]int, 20480),
	}

	for i, arr := range slices {
		for j := 0; j < len(arr); j++ {
			slices[i][j] = r.Int()
		}
	}

	return slices
}

func main() {
	benchmark_slices := generateBenchmarkSlices()

	methods := map[string]func([]int) []int{
		"INSERTION": InsertionSort,
		"SELECTION": SelectionSort,
		"HEAP":      HeapSort,
		"QUICK":     QuickSort,
	}

	for sort_name, sort_handler := range methods {
		fmt.Println("\n===", sort_name, "SORT ===")

		// First, sort the test array to check if the algorithm works correctly.
		test_slice := []int{5, 4, 1, 8, 34, 198, 89, 1, 0, 12, 53}
		fmt.Println("Slice:", test_slice)
		fmt.Println("Sorted slice:", sort_handler(test_slice))

		prev_exec_t := int64(0)
		t_diff := 0.0

		// Start benchmark with different slice sizes.
		fmt.Println("Benchmark:")
		for _, s := range benchmark_slices {
			start := time.Now().UnixNano()
			sort_handler(s)
			exec_t := time.Now().UnixNano() - start

			// Calculate t(n) / t(n - 1)
			if prev_exec_t != 0 {
				t_diff = float64(exec_t) / float64(prev_exec_t)
			}
			prev_exec_t = exec_t

			fmt.Println("Size:", len(s), ", Time:", time.Now().UnixNano()-start,
				", Time Diff:", t_diff)
		}
	}

}
