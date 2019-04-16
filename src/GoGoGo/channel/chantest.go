package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Run() {

	court := make(chan int)
	wg.Add(2)

	go player("lecon", court)
	go player("spawn", court)

	court <- 1
	wg.Wait()

}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		val := rand.Intn(100)
		if val%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道，表示我们输了
			close(court)
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}
