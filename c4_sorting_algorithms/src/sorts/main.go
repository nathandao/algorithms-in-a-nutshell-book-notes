package main

import "fmt"

func main() {
	arr := []int{17, 3, 23, 52, 12, 32, 45}

	fmt.Println("Initial array: ", arr)

	fmt.Println("Insertion Sort:")
	InsertionSort(arr)
	fmt.Println(arr)

}
