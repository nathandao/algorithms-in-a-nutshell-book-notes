package main

func SequentialSearch(s []int, val int) bool {
	for _, v := range s {
		if val == v {
			return true
		}
	}
	return false
}
