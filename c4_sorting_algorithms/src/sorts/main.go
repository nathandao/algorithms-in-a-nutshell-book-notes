package main

import "fmt"

var initArr = []int{13, 5, 6, 1, 45, 123, 824, 62, 940, 173, 891, 401, 441}

func main() {
	arr := initArr
	fmt.Println("Initial array: ", arr)

	fmt.Println("Insertion Sort:")
	InsertionSort(arr)
	fmt.Println(arr)

	// Reset arr & start selection sort
	arr = initArr
	fmt.Println("Selection Sort:")
	SelectionSort(arr)
	fmt.Println(arr)

}
