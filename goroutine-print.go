package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		ch = make(chan int, 100)
		wg = sync.WaitGroup{}
	)
	defer close(ch)

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			select {
			case num, ok := <-ch:
				if ok {
					fmt.Printf("Printer%d-%d\n", num%3+1, num+1)
					num++
					ch <- num
					wg.Done()
				}
			}
		}()
	}

	ch <- 0
	wg.Wait()
	time.Sleep(time.Second * 2)
}
