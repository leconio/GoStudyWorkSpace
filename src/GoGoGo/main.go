package main

import (
	"./channel"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//sync1.Run()
	//channel.Run()
	channel.Run2()
}

func hello() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}

		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)

			}
		}
	}()

	fmt.Println("waiting to print")
	wg.Wait()

	fmt.Println("\nTerminating Program")

}
