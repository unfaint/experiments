package main

import (
	"fmt"

	"github.com/unfaint/experiments/slice_e"
)

func main() {
	slice := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		fmt.Printf("before %v %v: %v \n", len(slice), cap(slice), slice)
		slice = slice_e.Extend(slice, i)
		fmt.Printf("after %v %v: %v \n", len(slice), cap(slice), slice)
	}
	slice = slice_e.Append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Printf("after append %v %v: %v \n", len(slice), cap(slice), slice)
}
