package main

import (
	"fmt"
	"sync"
)

func Printer(num int, chin, chout chan int, wg *sync.WaitGroup) {
	count := 0
	for {
		select {
		case <-chin:
			fmt.Printf("Printer%d-%d\n", num, num+count*3)
			count++
			chout <- 1
			wg.Done()
		}
	}
}

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)
	var wg sync.WaitGroup
	wg.Add(100)

	go Printer(1, ch3, ch1, &wg)
	go Printer(2, ch1, ch2, &wg)
	go Printer(3, ch2, ch3, &wg)
	ch3 <- 1
	wg.Wait()
}
