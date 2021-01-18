package main

import (
	"fmt"
	"sync"
	"time"
)

func Printer(num int, chin, chout chan int, wg *sync.WaitGroup, tatolconunt *int) {
	// count := 0
	for {
		select {
		case <-chin:
			fmt.Printf("Printer%d-%d\n", num, *tatolconunt)
			*tatolconunt++
			wg.Done()

			if *tatolconunt > 100 {
				break
			}
			chout <- 1
		}
	}
}

func main() {
	var tatolconunt = 1
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)
	var wg sync.WaitGroup
	wg.Add(100)

	go Printer(1, ch3, ch1, &wg, &tatolconunt)
	go Printer(2, ch1, ch2, &wg, &tatolconunt)
	go Printer(3, ch2, ch3, &wg, &tatolconunt)
	ch3 <- 1
	wg.Wait()
	time.Sleep(time.Second * 5)
}
