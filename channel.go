package main

import (
	"log"
	"time"
)

func main() {
	initLog()

	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	if IsOpening(ch) {
		log.Println("ch IsOpening 1")
	}

	go forselect(ch)

	if IsOpening(ch) {
		log.Println("ch IsOpening 1")
	}

	close(ch)
	// time.Sleep(4 * time.Second)
	for j := 0; j < 10; j++ {
		if IsOpening(ch) {
			log.Println("ch IsOpening 2  ", j, "\n")
		} else {
			break
		}
	}

	// for j := 0; j < 10; j++ {
	// 	i, isOpening := <-ch
	// 	log.Println("j,i", j, i)
	// 	if isOpening {
	// 		// log.Println("main IsOpening\n")
	// 	} else {
	// 		log.Println("main Isclosed\n")
	// 		break
	// 	}
	// }

	IsclosedTicker(ch)
}

func IsOpening(cha <-chan int) bool {
	_, isOpening := <-cha

	if isOpening {
		log.Println("IsOpening return true")
	} else {
		log.Println("IsOpening return false")
	}
	return isOpening
}

func IsClosedSelect(ch <-chan int) bool {
	select {
	case <-ch:
		return true
		// default:
	}

	return false
}

func checkChannelStatus(ch <-chan int) bool {
	select {
	case val, ok := <-ch:
		if !ok {
			return true
		} else {
			log.Println("channel is opening,val=", val)
			return false
		}
	}
}

func forselect(ch <-chan int) {
	var retry = 0
	for {
		// isClosed := IsClosedSelect(ch)
		isClosed := checkChannelStatus(ch)
		if isClosed {
			log.Println("isClosed return true")
		} else {
			log.Println("isClosed return false")
		}
		if retry == 4 {
			log.Println("isClosed retry= ", retry, " end")
			break
		}
		retry++
	}
}

func initLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func IsclosedTicker(ch chan int) {
	log.Println("IsclosedTicker")
	ackChan := make(chan struct{})

	// initialize timer and retry count for sending message
	var (
		retry                       = 0
		retryInterval time.Duration = 5
	)
	ticker := time.NewTimer(retryInterval * time.Second)

LOOP:
	for {
		select {
		case <-ackChan:
			log.Println("IsclosedTicker closed")
			break LOOP
		case <-ticker.C:
			log.Println("IsclosedTicker ticker loop")
			if retry == 4 {
				close(ch)
				break LOOP
			}
			retry++
			ticker.Reset(time.Second * retryInterval)
		}
	}

}

func define() {
	make(chan int)
	make(chan int, 10)
	make(chan<- int)
	make(<-chan int)
	make(chan string)
	make(chan bool, 2)
	make(<-chan bool)
	make(chan<- bool)
	make(<-chan float32)
	make(chan<- float64, 64)

}
