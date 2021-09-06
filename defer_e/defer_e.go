package defer_e

import "fmt"

func DeferReturn() string {
	defer fmt.Println(1)
	return "return"
}

func DeferNoReturn() {
	defer fmt.Println(2)
}
