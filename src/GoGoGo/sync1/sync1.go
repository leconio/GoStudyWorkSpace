package sync1

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg      sync.WaitGroup
	counter int
	mutex   sync.Mutex
)

func Run() {
	wg.Add(2)

	go increment(1)
	go increment(2)

	wg.Wait()
	fmt.Println("finish", counter)
}

func increment(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			// 捕获 counter 的值
			value := counter
			// 当前 goroutine 从线程退出，并放回到队列
			runtime.Gosched()
			// 增加本地 value 变量的值
			value++
			// 将该值保存回 counter
			counter = value
		}
		mutex.Unlock()
	}
}
