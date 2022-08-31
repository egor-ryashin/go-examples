package main

import (
	"fmt"
	"sync"
)

func parallel() {
	var wg sync.WaitGroup
	out := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			k := 4
			for j := 0; j < 1000000000; j++ {
				k += j % 7
			}
			out <- k
			wg.Done()
		}()
		wg.Add(1)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	for k := range out {
		fmt.Println(k)
	}
}
