package sync_e

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type message struct {
	text string
	cond *sync.Cond
	read bool
}

func worker(i int, msg *message) {
	msg.cond.L.Lock()
	msg.cond.Wait()
	if !msg.read {
		msg.read = true
		fmt.Printf("%d: %s\n", i, msg.text)
	}
	msg.cond.L.Unlock()
}

func spawnWorkers(number int, msg *message) {
	for i := 0; i < number; i++ {
		go worker(i, msg)
	}
}

func RunLoop() {
	msg := message{
		cond: sync.NewCond(&sync.Mutex{}),
	}
	scanner := bufio.NewScanner(os.Stdin)
	spawnWorkers(3, &msg)
	fmt.Print("Enter text: ")
	for scanner.Scan() {
		msg.cond.L.Lock()
		msg.text = scanner.Text()
		msg.read = false
		msg.cond.L.Unlock()
		msg.cond.Broadcast()
	}
}
