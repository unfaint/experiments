package main

import (
	"fmt"

	"github.com/unfaint/experiments/defer_e"
)

func main() {
	fmt.Println(defer_e.DeferReturn())
	defer_e.DeferNoReturn()
}
