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

	// go forselect(ch)
	close(ch)

	for j := 0; j < 10; j++ {
		if IsOpening(ch) {
			log.Println("ch IsOpening j  ", j, "\n")
		} else {
			log.Println("ch IsClosed j  ", j, "\n")
			break
		}
	}

	// close(ch)

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
				// close(ch)
				break LOOP
			}
			retry++
			ticker.Reset(time.Second * retryInterval)
		}
	}

}
