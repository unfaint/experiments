package slice_e

import "fmt"

func Extend(slice []int, item int) []int {
	n := len(slice)
	if n == cap(slice) {
		fmt.Println("recreating...")
		newSlice := make([]int, len(slice), len(slice)*3/2+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = item
	return slice
}

func Append(slice []int, items ...int) []int {
	n := len(slice)
	total := len(slice) + len(items)
	if total > cap(slice) {
		newSlice := make([]int, total, total*3/2+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], items)
	return slice
}
