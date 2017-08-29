package main

import (
	"sync"
	"runtime"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("x: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("y: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
