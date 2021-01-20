package main

import (
	"fmt"
	"sync"
	"time"
)

func Printer(num int, count *int, wg *sync.WaitGroup, inch, outch chan int) {
forlabel:
	for {
		select {
		case <-inch:
			fmt.Printf("Printer%d-%d\n", num, *count)
			*count++
			if *count > 100 {
				wg.Done()
				break forlabel
			}
			outch <- 1
		}
	}
}

func main() {
	var (
		ch1   = make(chan int)
		ch2   = make(chan int)
		ch3   = make(chan int)
		wg    sync.WaitGroup
		count = 1
	)

	wg.Add(1)

	go Printer(1, &count, &wg, ch3, ch1)
	go Printer(2, &count, &wg, ch1, ch2)
	go Printer(3, &count, &wg, ch2, ch3)

	ch3 <- 1

	wg.Wait()
	time.Sleep(time.Second * 2)
}
