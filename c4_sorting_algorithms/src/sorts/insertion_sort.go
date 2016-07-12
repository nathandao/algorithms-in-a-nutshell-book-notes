package main

func InsertionSort(arr []int) []int {
	values := arr
	for i := 1; i < len(values); i++ {
		// Initial values
		j := i - 1
		val := values[i]
		pos := i

		// Interate from current position down, shift every number larger than val
		// up 1 position until reaches one that is smaller or equals val.
		for j >= 0 && values[j] > val {
			values[j+1] = values[j]
			pos = j
			j = j - 1
		}

		values[pos] = val
	}

	return values
}
