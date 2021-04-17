package main

import (
	"fmt"
	"sync"
	"time"
)

func Printer(ch chan int, wg *sync.WaitGroup) {
	select {
	case num, ok := <-ch:
		if ok {
			fmt.Printf("Printer%d-%d\n", num%3+1, num+1)
			wg.Done()
			num++
			ch <- num
		}
	}
}

func main() {
	var (
		ch = make(chan int)
		wg = sync.WaitGroup{}
	)
	defer close(ch)
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go Printer(ch, &wg)
	}

	ch <- 0
	wg.Wait()
	time.Sleep(time.Second * 2)
}
