package select_e

import "fmt"

func fibSelect(out chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case out <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func StartFibSelect(num int) {
	out := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < num; i++ {
			fmt.Println(<-out)
		}
		quit <- true
	}()
	fibSelect(out, quit)
}
