package channel

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func Run2() {
	batch := make(chan int)
	wg2.Add(1)
	go Runner(batch)

	batch <- 1
	wg2.Wait()
}

func Runner(baton chan int) {
	runner := <-baton
	fmt.Println("runner ", runner)

	if runner != 100 {
		go Runner(baton)
		runner += 1
		baton <- runner
	} else {
		fmt.Println("runner finish", runner)
		wg2.Done()
	}
}
