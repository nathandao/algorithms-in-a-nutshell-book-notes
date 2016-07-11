package main

func selectMax(arr []int) int {
	maxPos := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[maxPos] {
			maxPos = i
		}
	}
	return maxPos
}

func SelectionSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		// Find largest value in each sub-slice
		maxPos := selectMax(arr[0 : i+1])

		// If maxPos is not already at the end of slice,
		// swich its position to the end of slice
		if maxPos != i {
			temp := arr[maxPos]
			arr[maxPos] = arr[i]
			arr[i] = temp
		}
	}
}
