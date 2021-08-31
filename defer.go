package main

import "fmt"

func deferReturn() string {
	defer fmt.Println(1)
	return "return"
}

func deferNoReturn() {
	defer fmt.Println(2)
}

func main() {
	fmt.Println(deferReturn())
	deferNoReturn()
}
